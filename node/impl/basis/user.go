package basis

import (
	"context"
	"math/big"

	"github.com/LMF709268224/titan-vps/lib/filecoinbridge"
	"github.com/LMF709268224/titan-vps/node/handler"
)

func (m *Basis) GetBalance(ctx context.Context) (*big.Int, error) {
	userID := handler.GetID(ctx)

	cfg, err := m.GetBasisConfigFunc()
	if err != nil {
		log.Errorf("get config err:%s", err.Error())
		return big.NewInt(0), err
	}

	client := filecoinbridge.NewGrpcClient(cfg.LotusHTTPSAddr, cfg.TitanContractorAddr)

	return client.GetBalance(userID)
}

func (m *Basis) Recharge(ctx context.Context, address, rechargeAddr string) (string, error) {
	return m.RechargeManager.CreateRechargeOrder(address, rechargeAddr)
}

func (m *Basis) CancelRecharge(ctx context.Context, orderID string) error {
	return m.RechargeManager.CancelRechargeOrder(orderID)
}

func (m *Basis) Withdraw(ctx context.Context, address, withdrawAddr string) (string, error) {
	return m.WithdrawManager.CreateWithdrawOrder(address, withdrawAddr)
}

func (m *Basis) CancelWithdraw(ctx context.Context, orderID string) error {
	return m.WithdrawManager.CancelWithdrawOrder(orderID)
}
