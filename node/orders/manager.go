package orders

import (
	"context"
	"sync"
	"time"

	"github.com/LMF709268224/titan-vps/api"
	"github.com/LMF709268224/titan-vps/api/terrors"
	"github.com/LMF709268224/titan-vps/api/types"
	"github.com/LMF709268224/titan-vps/lib/filecoinbridge"
	"github.com/LMF709268224/titan-vps/node/config"
	"github.com/LMF709268224/titan-vps/node/db"
	"github.com/LMF709268224/titan-vps/node/modules/dtypes"
	"github.com/LMF709268224/titan-vps/node/transaction"
	"github.com/LMF709268224/titan-vps/node/vps"
	"github.com/filecoin-project/go-statemachine"
	"github.com/filecoin-project/pubsub"
	"github.com/ipfs/go-datastore"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("orders")

const (
	checkOrderInterval = 10 * time.Second
	orderTimeoutMinute = 10
	orderTimeoutTime   = orderTimeoutMinute * time.Minute
)

// Manager manager order
type Manager struct {
	stateMachineWait   sync.WaitGroup
	orderStateMachines *statemachine.StateGroup
	*db.SQLDB

	notification *pubsub.PubSub

	ongoingOrders map[string]*types.OrderRecord
	orderLock     *sync.Mutex

	cfg  config.MallCfg
	tMgr *transaction.Manager
	vMgr *vps.Manager
}

// NewManager returns a new manager instance
func NewManager(ds datastore.Batching, sdb *db.SQLDB, pb *pubsub.PubSub, getCfg dtypes.GetMallConfigFunc, fm *transaction.Manager, vm *vps.Manager) (*Manager, error) {
	cfg, err := getCfg()
	if err != nil {
		return nil, err
	}

	m := &Manager{
		SQLDB:         sdb,
		notification:  pb,
		ongoingOrders: make(map[string]*types.OrderRecord),
		orderLock:     &sync.Mutex{},
		cfg:           cfg,
		tMgr:          fm,
		vMgr:          vm,
	}

	// state machine initialization
	m.stateMachineWait.Add(1)
	m.orderStateMachines = statemachine.New(ds, m, OrderInfo{})

	return m, nil
}

// Start initializes and starts the order state machine and associated tickers
func (m *Manager) Start(ctx context.Context) {
	if err := m.initStateMachines(ctx); err != nil {
		log.Errorf("restartStateMachines err: %s", err.Error())
	}

	// go m.subscribeEvents()
	go m.checkOrdersTimeout()
}

func (m *Manager) checkOrdersTimeout() {
	ticker := time.NewTicker(checkOrderInterval)
	defer ticker.Stop()

	for {
		<-ticker.C

		for _, orderRecord := range m.ongoingOrders {
			orderID := orderRecord.OrderID
			addr := orderRecord.To

			info, err := m.LoadOrderRecord(orderID, orderTimeoutMinute)
			if err != nil {
				log.Errorf("checkOrderTimeout LoadOrderRecord %s , %s err:%s", addr, orderID, err.Error())
				continue
			}

			log.Debugf("checkout %s , %s ", addr, orderID)

			if info.State.Int() != Done.Int() && info.CreatedTime.Add(orderTimeoutTime).Before(time.Now()) {

				height := m.getHeight()

				err = m.orderStateMachines.Send(OrderHash(orderID), OrderTimeOut{Height: height})
				if err != nil {
					log.Errorf("checkOrderTimeout Send %s , %s err:%s", addr, orderID, err.Error())
					continue
				}
			}
		}
	}
}

func (m *Manager) getOrderIDByToAddress(to string) (string, bool) {
	for _, orderRecord := range m.ongoingOrders {
		if orderRecord.To == to {
			return orderRecord.OrderID, true
		}
	}

	return "", false
}

func (m *Manager) subscribeEvents() {
	subTransfer := m.notification.Sub(types.EventFvmTransferWatch.String())
	defer m.notification.Unsub(subTransfer)

	for {
		select {
		case u := <-subTransfer:
			tr := u.(*types.FvmTransferWatch)

			if orderID, exist := m.getOrderIDByToAddress(tr.To); exist {
				err := m.orderStateMachines.Send(OrderHash(orderID), PaymentResult{
					&PaymentInfo{
						TxHash: tr.TxHash,
						From:   tr.From,
						To:     tr.To,
						Value:  tr.Value,
					},
				})
				if err != nil {
					log.Errorf("subscribeNodeEvents Send %s err:%s", orderID, err.Error())
					continue
				}
			}
		}
	}
}

// Terminate stops the order state machine
func (m *Manager) Terminate(ctx context.Context) error {
	return m.orderStateMachines.Stop(ctx)
}

// CancelOrder cancel vps order
func (m *Manager) CancelOrder(orderID, userID string) error {
	order, err := m.LoadOrderRecord(orderID, orderTimeoutMinute)
	if err != nil {
		return &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}

	if order.UserID != userID {
		return &api.ErrWeb{Code: terrors.UserMismatch.Int(), Message: terrors.UserMismatch.String()}
	}

	height := m.getHeight()

	err = m.orderStateMachines.Send(OrderHash(orderID), OrderCancel{Height: height})
	if err != nil {
		return &api.ErrWeb{Code: terrors.StateMachinesError.Int(), Message: err.Error()}
	}

	return nil
}

// PaymentCompleted cancel vps order
func (m *Manager) PaymentCompleted(orderID, userID string) error {
	order, err := m.LoadOrderRecord(orderID, orderTimeoutMinute)
	if err != nil {
		return &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}

	if order.UserID != userID {
		return &api.ErrWeb{Code: terrors.UserMismatch.Int(), Message: terrors.UserMismatch.String()}
	}

	err = m.orderStateMachines.Send(OrderHash(orderID), PaymentResult{})
	if err != nil {
		return &api.ErrWeb{Code: terrors.StateMachinesError.Int(), Message: err.Error()}
	}

	return nil
}

// CreatedOrder create vps order
func (m *Manager) CreatedOrder(req *types.OrderRecord) error {
	m.stateMachineWait.Wait()
	req.CreatedHeight = m.getHeight()

	m.addOrder(req)

	// create order task
	err := m.orderStateMachines.Send(OrderHash(req.OrderID), CreateOrder{orderInfoFrom(req)})
	if err != nil {
		return &api.ErrWeb{Code: terrors.StateMachinesError.Int(), Message: err.Error()}
	}

	return nil
}

func (m *Manager) addOrder(req *types.OrderRecord) {
	m.orderLock.Lock()
	defer m.orderLock.Unlock()

	// if _, exist := m.ongoingOrders[req.OrderID]; exist {
	// 	return xerrors.New("user have order")
	// }

	m.ongoingOrders[req.OrderID] = req

	return
}

func (m *Manager) removeOrder(orderID string) {
	m.orderLock.Lock()
	defer m.orderLock.Unlock()

	delete(m.ongoingOrders, orderID)
}

func (m *Manager) getHeight() int64 {
	var msg filecoinbridge.TipSet
	err := filecoinbridge.ChainHead(&msg, m.cfg.LotusHTTPSAddr)
	if err != nil {
		log.Errorf("ChainHead err:%s", err.Error())
		return 0
	}

	return msg.Height
}

func (m *Manager) GetOrderTimeoutMinute() int {
	return orderTimeoutMinute
}
