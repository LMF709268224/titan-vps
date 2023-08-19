package mall

import (
	"context"

	"github.com/LMF709268224/titan-vps/api"
	"github.com/LMF709268224/titan-vps/api/terrors"
	"github.com/LMF709268224/titan-vps/api/types"
	"github.com/LMF709268224/titan-vps/node/handler"
	"github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/gbrlsnchs/jwt/v3"
)

func (m *Mall) GetAdminSignCode(ctx context.Context, userID string) (string, error) {
	exist, err := m.AdminExists(userID)
	if err != nil {
		return "", &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}

	if !exist {
		return "", &api.ErrWeb{Code: terrors.NotAdministrator.Int(), Message: terrors.NotAdministrator.String()}
	}

	return m.UserMgr.GenerateSignCode(userID), nil
}

func (m *Mall) LoginAdmin(ctx context.Context, user *types.UserReq) (*types.LoginResponse, error) {
	userID := user.UserId

	exist, err := m.AdminExists(userID)
	if err != nil {
		return nil, &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}

	if !exist {
		return nil, &api.ErrWeb{Code: terrors.NotAdministrator.Int(), Message: terrors.NotAdministrator.String()}
	}

	code, err := m.UserMgr.GetSignCode(userID)
	if err != nil {
		return nil, &api.ErrWeb{Code: terrors.NotFoundSignCode.Int(), Message: terrors.NotFoundSignCode.String()}
	}
	signature := user.Signature
	address, err := verifyEthMessage(code, signature)
	if err != nil {
		return nil, &api.ErrWeb{Code: terrors.SignError.Int(), Message: err.Error()}
	}

	p := types.JWTPayload{
		ID:        address,
		LoginType: int64(user.Type),
		Allow:     []auth.Permission{api.RoleAdmin},
	}
	rsp := &types.LoginResponse{}
	tk, err := jwt.Sign(&p, m.APISecret)
	if err != nil {
		return rsp, &api.ErrWeb{Code: terrors.SignError.Int(), Message: err.Error()}
	}
	rsp.UserId = address
	rsp.Token = string(tk)

	return rsp, nil
}

func (m *Mall) GetRechargeAddresses(ctx context.Context, limit, offset int64) (*types.GetRechargeAddressResponse, error) {
	info, err := m.LoadRechargeAddresses(limit, offset)
	if err != nil {
		return nil, &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}

	return info, nil
}

func (m *Mall) GetWithdrawalRecords(ctx context.Context, limit, offset int64) (*types.WithdrawResponse, error) {
	info, err := m.LoadWithdrawRecords(limit, offset)
	if err != nil {
		return nil, &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}

	return info, nil
}

func (m *Mall) UpdateWithdrawalRecord(ctx context.Context, orderID, withdrawHash string) error {
	userID := handler.GetID(ctx)

	info, err := m.LoadWithdrawRecord(orderID)
	if err != nil {
		return &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}

	info.WithdrawHash = withdrawHash
	info.Executor = userID

	err = m.UpdateWithdrawRecord(info, types.WithdrawDone)
	if err != nil {
		return &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}

	return nil
}

func (m *Mall) AddAdminUser(ctx context.Context, userID, nickName string) error {
	err := m.SaveAdminInfo(userID, nickName)
	if err != nil {
		return &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}

	return nil
}
