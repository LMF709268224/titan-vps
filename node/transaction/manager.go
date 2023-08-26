package transaction

import (
	"sync"

	"github.com/LMF709268224/titan-vps/api"
	"github.com/LMF709268224/titan-vps/api/terrors"
	"github.com/LMF709268224/titan-vps/node/config"
	"github.com/LMF709268224/titan-vps/node/db"
	"github.com/LMF709268224/titan-vps/node/modules/dtypes"
	"github.com/filecoin-project/pubsub"
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("transaction")

// Manager is the node manager responsible for managing the online nodes
type Manager struct {
	notification *pubsub.PubSub
	*db.SQLDB

	cfg config.MallCfg

	tronAddrs sync.Map
}

// NewManager creates a new instance of the node manager
func NewManager(pb *pubsub.PubSub, getCfg dtypes.GetMallConfigFunc, db *db.SQLDB) (*Manager, error) {
	cfg, err := getCfg()
	if err != nil {
		return nil, err
	}

	manager := &Manager{
		notification: pb,
		cfg:          cfg,
		SQLDB:        db,
	}

	manager.initTronAddress(cfg.RechargeAddresses)

	go manager.watchTronTransactions()

	return manager, nil
}

// AllocateTronAddress get a fvm address
func (m *Manager) AllocateTronAddress(userID string) (string, error) {
	addr, err := m.LoadUnusedRechargeAddress()
	if err != nil {
		return "", &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}
	if addr == "" {
		return "", &api.ErrWeb{Code: terrors.NotFoundAddress.Int(), Message: terrors.NotFoundAddress.String()}
	}

	err = m.UpdateRechargeAddressOfUser(addr, userID)
	if err != nil {
		return "", &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}
	m.addTronAddr(addr, userID)

	return addr, nil
}
