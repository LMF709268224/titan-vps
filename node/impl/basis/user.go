package basis

import (
	"context"

	"github.com/LMF709268224/titan-vps/api/types"
	"github.com/LMF709268224/titan-vps/node/handler"
	"golang.org/x/xerrors"
)

func (m *Basis) GetBalance(ctx context.Context) (string, error) {
	userID := handler.GetID(ctx)

	return m.LoadUserToken(userID)
}

func (m *Basis) GetRechargeAddress(ctx context.Context) (string, error) {
	userID := handler.GetID(ctx)

	addr, err := m.GetRechargeAddressOfUser(userID)
	if err != nil {
		return "", err
	}

	if addr == "" {
		list, err := m.GetRechargeAddresses()
		if err != nil {
			return "", err
		}

		if len(list) == 0 {
			return "", xerrors.New("not found address")
		}

		addr := list[0]
		err = m.UpdateRechargeAddressOfUser(addr, userID)
		if err != nil {
			return "", err
		}
	}

	return addr, nil
}

func (m *Basis) Withdraw(ctx context.Context, withdrawAddr, value string) error {
	userID := handler.GetID(ctx)

	return m.WithdrawManager.CreateWithdrawOrder(userID, withdrawAddr, value)
}

func (m *Basis) GetRechargeRecord(ctx context.Context, limit, offset int64) (*types.RechargeResponse, error) {
	userID := handler.GetID(ctx)

	return m.LoadRechargeRecordsByUser(userID, limit, offset)
}

func (m *Basis) GetWithdrawRecord(ctx context.Context, limit, offset int64) (*types.WithdrawResponse, error) {
	userID := handler.GetID(ctx)

	return m.LoadWithdrawRecordsByUser(userID, limit, offset)
}
