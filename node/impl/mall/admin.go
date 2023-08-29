package mall

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/LMF709268224/titan-vps/api"
	"github.com/LMF709268224/titan-vps/api/terrors"
	"github.com/LMF709268224/titan-vps/api/types"
	"github.com/LMF709268224/titan-vps/node/handler"
	"github.com/LMF709268224/titan-vps/node/utils"
	"github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/gbrlsnchs/jwt/v3"
)

// GetAdminSignCode generates a sign code for an admin user.
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

// LoginAdmin authenticates an admin user and generates a JWT token.
func (m *Mall) LoginAdmin(ctx context.Context, user *types.UserReq) (*types.LoginResponse, error) {
	userID := user.UserId

	code, err := m.UserMgr.GetSignCode(userID)
	if err != nil {
		return nil, &api.ErrWeb{Code: terrors.NotFoundSignCode.Int(), Message: terrors.NotFoundSignCode.String()}
	}

	signature := user.Signature
	address, err := verifyEthMessage(code, signature)
	if err != nil {
		return nil, &api.ErrWeb{Code: terrors.SignError.Int(), Message: err.Error()}
	}

	if strings.ToLower(userID) != strings.ToLower(address) {
		return nil, &api.ErrWeb{Code: terrors.UserMismatch.Int(), Message: fmt.Sprintf("%s,%s", userID, address)}
	}

	exist, err := m.AdminExists(address)
	if err != nil {
		return nil, &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}

	if !exist {
		return nil, &api.ErrWeb{Code: terrors.NotAdministrator.Int(), Message: terrors.NotAdministrator.String()}
	}

	p := types.JWTPayload{
		ID:        address,
		LoginType: int64(user.Type),
		Allow:     []auth.Permission{api.RoleAdmin},
	}

	tk, err := jwt.Sign(&p, m.APISecret)
	if err != nil {
		return nil, &api.ErrWeb{Code: terrors.SignError.Int(), Message: err.Error()}
	}

	rsp := &types.LoginResponse{}
	rsp.UserId = address
	rsp.Token = string(tk)

	return rsp, nil
}

// GetRechargeAddresses retrieves recharge addresses with pagination.
func (m *Mall) GetRechargeAddresses(ctx context.Context, limit, page int64) (*types.GetRechargeAddressResponse, error) {
	info, err := m.LoadRechargeAddresses(limit, page)
	if err != nil {
		return nil, &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}

	return info, nil
}

// GetWithdrawalRecords retrieves withdrawal records with optional filtering.
func (m *Mall) GetWithdrawalRecords(ctx context.Context, req *types.GetWithdrawRequest) (*types.GetWithdrawResponse, error) {
	statuses := make([]types.WithdrawState, 0)
	if req.State == "" {
		statuses = []types.WithdrawState{types.WithdrawCreate, types.WithdrawDone, types.WithdrawRefund}
	} else {
		s2, err := strconv.Atoi(req.State)
		if err != nil {
			return nil, &api.ErrWeb{Code: terrors.ParametersWrong.Int(), Message: fmt.Sprintf("state is %s , err:%s", req.State, err.Error())}
		}

		statuses = []types.WithdrawState{types.WithdrawState(s2)}
	}

	info, err := m.LoadWithdrawRecords(req.Limit, req.Offset, statuses, req.UserID, req.StartDate, req.EndDate)
	if err != nil {
		return nil, &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}

	return info, nil
}

// ApproveUserWithdrawal approves a user withdrawal request.
func (m *Mall) ApproveUserWithdrawal(ctx context.Context, orderID, withdrawHash string) error {
	userID := handler.GetID(ctx)

	info, err := m.LoadWithdrawRecord(orderID)
	if err != nil {
		return &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}

	if info.State != types.WithdrawCreate {
		return &api.ErrWeb{Code: terrors.StatusNotEditable.Int(), Message: string(info.State)}
	}

	info.WithdrawHash = withdrawHash
	info.Executor = userID

	err = m.UpdateWithdrawRecord(info, types.WithdrawDone)
	if err != nil {
		return &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}

	return nil
}

// RejectUserWithdrawal rejects a user withdrawal request.
func (m *Mall) RejectUserWithdrawal(ctx context.Context, orderID string) error {
	userID := handler.GetID(ctx)

	info, err := m.LoadWithdrawRecord(orderID)
	if err != nil {
		return &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}

	if info.State != types.WithdrawCreate {
		return &api.ErrWeb{Code: terrors.StatusNotEditable.Int(), Message: string(info.State)}
	}

	info.Executor = userID

	err = m.UpdateWithdrawRecord(info, types.WithdrawRefund)
	if err != nil {
		return &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}

	original, err := m.LoadUserBalance(info.UserID)
	if err != nil {
		return &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}

	newValue, err := utils.AddBigInt(original, info.Value)
	if err != nil {
		return err
	}

	err = m.UpdateUserBalance(info.UserID, newValue, original)
	if err != nil {
		return &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}

	return nil
}

// AddAdminUser adds an admin user with a userID and nickname.
func (m *Mall) AddAdminUser(ctx context.Context, userID, nickName string) error {
	err := m.SaveAdminInfo(userID, nickName)
	if err != nil {
		return &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}

	return nil
}

// SupplementRechargeOrder supplements a recharge order.
func (m *Mall) SupplementRechargeOrder(ctx context.Context, hash string) error {
	return m.TransactionMgr.SupplementOrder(hash)
}
