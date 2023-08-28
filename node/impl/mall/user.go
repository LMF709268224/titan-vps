package mall

import (
	"context"
	"encoding/json"

	"github.com/LMF709268224/titan-vps/api"
	"github.com/LMF709268224/titan-vps/lib/aliyun"
	"github.com/LMF709268224/titan-vps/lib/trxbridge"
	"github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/gbrlsnchs/jwt/v3"

	"github.com/LMF709268224/titan-vps/api/terrors"
	"github.com/LMF709268224/titan-vps/api/types"
	"github.com/LMF709268224/titan-vps/node/handler"
	"github.com/LMF709268224/titan-vps/node/utils"
)

// GetBalance Get user balance
func (m *Mall) GetBalance(ctx context.Context) (*types.UserInfo, error) {
	userID := handler.GetID(ctx)

	uInfo := &types.UserInfo{UserID: userID}

	balance, err := m.LoadUserBalance(userID)
	if err != nil {
		return uInfo, &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}

	uInfo.Balance = balance

	list, err := m.LoadWithdrawRecordsByUserAndState(userID, types.WithdrawCreate)
	if err != nil {
		return uInfo, &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}

	lockBalance := "0"
	for _, info := range list {
		b, err := utils.BigIntAdd(info.Value, lockBalance)
		if err == nil {
			lockBalance = b
		}
	}

	uInfo.LockedBalance = lockBalance

	return uInfo, nil
}

// GetRechargeAddress  Get user recharge address
func (m *Mall) GetRechargeAddress(ctx context.Context) (string, error) {
	userID := handler.GetID(ctx)

	address, err := m.LoadRechargeAddressOfUser(userID)
	if err != nil {
		return address, &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}

	if address == "" {
		_, err = m.TransactionMgr.AllocateTronAddress(userID)
		if err != nil {
			return "", err
		}
	}

	return address, nil
}

// Withdraw user Withdraw
func (m *Mall) Withdraw(ctx context.Context, withdrawAddr, value string) error {
	userID := handler.GetID(ctx)

	if withdrawAddr == "" || value == "" {
		return &api.ErrWeb{Code: terrors.ParametersWrong.Int(), Message: terrors.ParametersWrong.String()}
	}

	_, err := utils.BigIntReduce(value, "0")
	if err != nil {
		return err
	}

	cfg, err := m.GetMallConfigFunc()
	if err != nil {
		log.Errorf("get config err:%s", err.Error())
		return &api.ErrWeb{Code: terrors.ConfigError.Int(), Message: err.Error()}
	}

	node := trxbridge.NewGrpcClient(cfg.TrxHTTPSAddr)
	err = node.Start()
	if err != nil {
		return &api.ErrWeb{Code: terrors.ConfigError.Int(), Message: err.Error()}
	}

	_, err = node.GetAccount(withdrawAddr)
	if err != nil {
		return &api.ErrWeb{Code: terrors.WithdrawAddrError.Int(), Message: err.Error()}
	}

	return m.WithdrawManager.CreateWithdrawOrder(userID, withdrawAddr, value)
}

func (m *Mall) GetUserRechargeRecords(ctx context.Context, limit, page int64) (*types.RechargeResponse, error) {
	userID := handler.GetID(ctx)

	info, err := m.LoadRechargeRecordsByUser(userID, limit, page)
	if err != nil {
		return nil, &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}

	return info, nil
}

func (m *Mall) GetUserWithdrawalRecords(ctx context.Context, limit, page int64) (*types.GetWithdrawResponse, error) {
	userID := handler.GetID(ctx)

	info, err := m.LoadWithdrawRecordsByUser(userID, limit, page)
	if err != nil {
		return nil, &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}

	return info, nil
}

func (m *Mall) GetUserInstanceRecords(ctx context.Context, limit, offset int64) (*types.MyInstanceResponse, error) {
	userID := handler.GetID(ctx)
	k, s := m.getAccessKeys()
	instanceInfos, err := m.LoadMyInstancesInfo(userID, limit, offset)
	if err != nil {
		log.Errorf("LoadMyInstancesInfo err: %s", err.Error())
		return nil, &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}
	for _, instanceInfo := range instanceInfos.List {
		var instanceIds []string
		instanceIds = append(instanceIds, instanceInfo.InstanceId)
		rsp, err := aliyun.DescribeInstanceStatus(instanceInfo.Location, k, s, instanceIds)
		if err != nil {
			log.Errorf("DescribeInstanceStatus err: %s", *err.Message)
			return nil, &api.ErrWeb{Code: terrors.ParametersWrong.Int(), Message: *err.Message}
		}
		renewInfo := types.SetRenewOrderReq{
			RegionID:   instanceInfo.Location,
			InstanceId: instanceInfo.InstanceId,
		}
		instanceInfo.State = *rsp.Body.InstanceStatuses.InstanceStatus[0].Status
		instanceInfo.Renew = ""
		if instanceInfo.State == "Stopped" {
			continue
		}
		status, errEk := m.GetRenewInstance(ctx, renewInfo)
		if errEk != nil {
			log.Errorf("GetRenewInstance err: %s", errEk.Error())
			continue
		}
		instanceInfo.Renew = status
		instanceExpiredTime, err := aliyun.DescribeInstances(instanceInfo.Location, k, s, instanceIds)
		if err != nil {
			log.Errorf("DescribeInstances err: %s", *err.Message)
			continue
		}
		if len(instanceExpiredTime.Body.Instances.Instance) > 0 {
			instanceInfo.ExpiredTime = *instanceExpiredTime.Body.Instances.Instance[0].ExpiredTime
		}
	}

	return instanceInfos, nil
}

func (m *Mall) GetInstanceDetailsInfo(ctx context.Context, instanceID string) (*types.InstanceDetails, error) {
	userID := handler.GetID(ctx)

	info, err := m.LoadInstanceDetailsInfo(userID, instanceID)
	if err != nil {
		return nil, &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}
	if info.DataDiskString != "" {
		if err := json.Unmarshal([]byte(info.DataDiskString), &info.DataDisk); err != nil {
			return info, nil
		}
	}
	return info, nil
}

func (m *Mall) UpdateInstanceName(ctx context.Context, instanceID, instanceName string) error {
	userID := handler.GetID(ctx)
	err := m.UpdateVpsInstanceName(instanceID, instanceName, userID)
	if err != nil {
		return &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}
	return nil
}

func (m *Mall) GetInstanceDefaultInfo(ctx context.Context, req *types.InstanceTypeFromBaseReq) (*types.InstanceTypeResponse, error) {
	req.Offset = req.Limit * (req.Page - 1)
	instanceInfo, err := m.LoadInstanceDefaultInfo(req)
	if err != nil {
		return nil, &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}

	usdRate := utils.GetUSDRate()
	for _, info := range instanceInfo.List {
		info.OriginalPrice = info.OriginalPrice / usdRate
		info.Price = info.Price / usdRate
	}
	return instanceInfo, nil
}

func (m *Mall) GetInstanceCpuInfo(ctx context.Context, req *types.InstanceTypeFromBaseReq) ([]*int32, error) {
	return m.LoadInstanceCpuInfo(req)
}

func (m *Mall) GetInstanceMemoryInfo(ctx context.Context, req *types.InstanceTypeFromBaseReq) ([]*float32, error) {
	return m.LoadInstanceMemoryInfo(req)
}

func (m *Mall) GetSignCode(ctx context.Context, userID string) (string, error) {
	return m.UserMgr.GenerateSignCode(userID), nil
}

func (m *Mall) Login(ctx context.Context, user *types.UserReq) (*types.LoginResponse, error) {
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
	p := types.JWTPayload{
		ID:        address,
		LoginType: int64(user.Type),
		Allow:     []auth.Permission{api.RoleUser},
	}
	rsp := &types.LoginResponse{}
	tk, err := jwt.Sign(&p, m.APISecret)
	if err != nil {
		return rsp, &api.ErrWeb{Code: terrors.SignError.Int(), Message: err.Error()}
	}
	rsp.UserId = address
	rsp.Token = string(tk)
	err = m.initUser(address)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func (m *Mall) initUser(userID string) error {
	exist, err := m.UserExists(userID)
	if err != nil {
		return &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}
	if !exist {
		err = m.SaveUserInfo(&types.UserInfo{UserID: userID, Balance: "0"})
		if err != nil {
			return &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
		}
	}
	// init recharge address
	addr, err := m.LoadRechargeAddressOfUser(userID)
	if err != nil {
		return &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}
	if addr == "" {
		_, err = m.TransactionMgr.AllocateTronAddress(userID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Mall) Logout(ctx context.Context, user *types.UserReq) error {
	userID := handler.GetID(ctx)
	log.Warnf("user id : %s", userID)
	// delete(m.UserMgr.User, user.UserId)
	return nil
}
