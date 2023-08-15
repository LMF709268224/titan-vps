package exchange

import (
	"time"

	"github.com/LMF709268224/titan-vps/api/types"
	"github.com/LMF709268224/titan-vps/lib/trxbridge/core"
	"github.com/LMF709268224/titan-vps/node/config"
	"github.com/LMF709268224/titan-vps/node/db"
	"github.com/LMF709268224/titan-vps/node/modules/dtypes"
	"github.com/LMF709268224/titan-vps/node/transaction"
	"github.com/LMF709268224/titan-vps/node/utils"
	"github.com/filecoin-project/pubsub"
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("exchange")

const (
	checkOrderInterval = 10 * time.Second
	orderTimeoutTime   = 10 * time.Minute
)

// RechargeManager manager recharge order
type RechargeManager struct {
	*db.SQLDB
	cfg    config.BasisCfg
	notify *pubsub.PubSub

	tMgr *transaction.Manager
}

// NewRechargeManager returns a new manager instance
func NewRechargeManager(sdb *db.SQLDB, pb *pubsub.PubSub, getCfg dtypes.GetBasisConfigFunc, fm *transaction.Manager) (*RechargeManager, error) {
	cfg, err := getCfg()
	if err != nil {
		return nil, err
	}

	m := &RechargeManager{
		SQLDB:  sdb,
		notify: pb,
		cfg:    cfg,

		tMgr: fm,
	}

	go m.subscribeEvents()

	return m, nil
}

func (m *RechargeManager) subscribeEvents() {
	subTransfer := m.notify.Sub(types.EventTronTransferWatch.String())
	defer m.notify.Unsub(subTransfer)

	for {
		select {
		case u := <-subTransfer:
			tr := u.(*types.TronTransferWatch)

			m.handleTronTransfer(tr)
		}
	}
}

func (m *RechargeManager) handleTronTransfer(tr *types.TronTransferWatch) {
	if tr.State != core.Transaction_Result_SUCCESS {
		return
	}

	userAddr := tr.UserAddr
	height := getTronHeight(m.cfg.TrxHTTPSAddr)

	info := &types.RechargeRecord{
		OrderID:       tr.TxHash,
		User:          userAddr,
		CreatedHeight: height,
		DoneHeight:    height,
		Value:         tr.Value,
		From:          tr.From,
		State:         types.RechargeCreate,
	}

	info.Msg = tr.State.String()
	err := m.SaveRechargeInfo(info)
	if err != nil {
		return
	}

	state := types.RechargeRefund

	oldValue, err := m.LoadUserToken(userAddr)
	if err != nil {
		log.Errorf("%s LoadUserToken hash:%s,state:%d, err:%s", info.OrderID, info.RechargeHash, state, err.Error())
		return
	}

	value := utils.BigIntAdd(oldValue, tr.Value)
	err = m.UpdateUserToken(userAddr, value, oldValue)
	if err != nil {
		log.Errorf("%s UpdateUserToken hash:%s,state:%d, err:%s", info.OrderID, info.RechargeHash, state, err.Error())
		return
	}

	state = types.RechargeDone

	err = m.UpdateRechargeRecord(info, state)
	if err != nil {
		log.Errorf("%s UpdateRechargeRecord hash:%s,state:%d, err:%s", info.OrderID, info.RechargeHash, state, err.Error())
	}
}
