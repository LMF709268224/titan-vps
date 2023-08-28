package exchange

import (
	"strings"

	"github.com/LMF709268224/titan-vps/api"
	"github.com/LMF709268224/titan-vps/api/terrors"
	"github.com/LMF709268224/titan-vps/api/types"
	"github.com/LMF709268224/titan-vps/node/config"
	"github.com/LMF709268224/titan-vps/node/db"
	"github.com/LMF709268224/titan-vps/node/modules/dtypes"
	"github.com/LMF709268224/titan-vps/node/transaction"
	"github.com/LMF709268224/titan-vps/node/utils"
	"github.com/filecoin-project/pubsub"
	"github.com/google/uuid"
)

// WithdrawManager manager withdraw order
type WithdrawManager struct {
	*db.SQLDB
	cfg          config.MallCfg
	notification *pubsub.PubSub

	tMgr *transaction.Manager
}

// NewWithdrawManager returns a new manager instance
func NewWithdrawManager(sdb *db.SQLDB, pb *pubsub.PubSub, getCfg dtypes.GetMallConfigFunc, fm *transaction.Manager) (*WithdrawManager, error) {
	cfg, err := getCfg()
	if err != nil {
		return nil, err
	}

	m := &WithdrawManager{
		SQLDB:        sdb,
		notification: pb,
		cfg:          cfg,

		tMgr: fm,
	}

	return m, nil
}

// CreateWithdrawOrder create a withdraw order
func (m *WithdrawManager) CreateWithdrawOrder(userID, withdrawAddr, value string) (err error) {
	original, err := m.LoadUserBalance(userID)
	if err != nil {
		return &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}

	newValue, err := utils.BigIntReduce(original, value)
	if err != nil {
		return &api.ErrWeb{Code: terrors.InsufficientBalance.Int(), Message: terrors.InsufficientBalance.String()}
	}

	hash := uuid.NewString()
	orderID := strings.Replace(hash, "-", "", -1)

	info := &types.WithdrawRecord{
		OrderID:      orderID,
		UserID:       userID,
		WithdrawAddr: withdrawAddr,
		Value:        value,
		State:        types.WithdrawCreate,
	}

	err = m.SaveWithdrawInfoAndUserBalance(info, newValue, original)
	if err != nil {
		return &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}

	return nil
}
