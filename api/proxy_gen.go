// Code generated by titan/gen/api. DO NOT EDIT.

package api

import (
	"context"
	"github.com/LMF709268224/titan-vps/api/types"
	"github.com/LMF709268224/titan-vps/journal/alerting"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

var ErrNotSupported = xerrors.New("method not supported")

type AdminAPIStruct struct {
	Internal struct {
		AddAdminUser func(p0 context.Context, p1 string, p2 string) error `perm:"admin"`

		ApproveUserWithdrawal func(p0 context.Context, p1 string, p2 string) error `perm:"admin"`

		GetAdminSignCode func(p0 context.Context, p1 string) (string, error) `perm:"default"`

		GetRechargeAddresses func(p0 context.Context, p1 int64, p2 int64) (*types.GetRechargeAddressResponse, error) `perm:"admin"`

		GetWithdrawalRecords func(p0 context.Context, p1 *types.GetWithdrawRequest) (*types.GetWithdrawResponse, error) `perm:"default"`

		LoginAdmin func(p0 context.Context, p1 *types.UserReq) (*types.LoginResponse, error) `perm:"default"`

		MintToken func(p0 context.Context, p1 string) (string, error) `perm:"admin"`

		RejectUserWithdrawal func(p0 context.Context, p1 string) error `perm:"admin"`
	}
}

type AdminAPIStub struct {
}

type CommonStruct struct {
	Internal struct {
		AuthNew func(p0 context.Context, p1 *types.JWTPayload) (string, error) `perm:"admin"`

		AuthVerify func(p0 context.Context, p1 string) (*types.JWTPayload, error) `perm:"default"`

		Closing func(p0 context.Context) (<-chan struct{}, error) `perm:"admin"`

		Discover func(p0 context.Context) (types.OpenRPCDocument, error) `perm:"admin"`

		LogAlerts func(p0 context.Context) ([]alerting.Alert, error) `perm:"admin"`

		LogList func(p0 context.Context) ([]string, error) `perm:"admin"`

		LogSetLevel func(p0 context.Context, p1 string, p2 string) error `perm:"admin"`

		Session func(p0 context.Context) (uuid.UUID, error) `perm:"admin"`

		Shutdown func(p0 context.Context) error `perm:"admin"`

		Version func(p0 context.Context) (APIVersion, error) `perm:"default"`
	}
}

type CommonStub struct {
}

type MallStruct struct {
	CommonStruct

	OrderAPIStruct

	UserAPIStruct

	AdminAPIStruct

	Internal struct {
		AttachKeyPair func(p0 context.Context, p1 string, p2 string, p3 []string) ([]*types.AttachKeyPairResponse, error) `perm:"default"`

		CreateInstance func(p0 context.Context, p1 *types.CreateInstanceReq) (*types.CreateInstanceResponse, error) `perm:"default"`

		CreateKeyPair func(p0 context.Context, p1 string, p2 string) (*types.CreateKeyPairResponse, error) `perm:"default"`

		DescribeAvailableResourceForDesk func(p0 context.Context, p1 *types.AvailableResourceReq) ([]*types.AvailableResourceResponse, error) `perm:"default"`

		DescribeImages func(p0 context.Context, p1 string, p2 string) ([]*types.DescribeImageResponse, error) `perm:"default"`

		DescribeInstanceType func(p0 context.Context, p1 *types.DescribeInstanceTypeReq) (*types.DescribeInstanceTypeResponse, error) `perm:"default"`

		DescribeInstances func(p0 context.Context, p1 string, p2 string) error `perm:"default"`

		DescribePrice func(p0 context.Context, p1 *types.DescribePriceReq) (*types.DescribePriceResponse, error) `perm:"default"`

		DescribeRecommendInstanceType func(p0 context.Context, p1 *types.DescribeRecommendInstanceTypeReq) ([]*types.DescribeRecommendInstanceResponse, error) `perm:"default"`

		DescribeRegions func(p0 context.Context) (map[string]string, error) `perm:"default"`

		GetInstanceCpuInfo func(p0 context.Context, p1 *types.InstanceTypeFromBaseReq) ([]*int32, error) `perm:"default"`

		GetInstanceDefaultInfo func(p0 context.Context, p1 *types.InstanceTypeFromBaseReq) (*types.InstanceTypeResponse, error) `perm:"default"`

		GetInstanceMemoryInfo func(p0 context.Context, p1 *types.InstanceTypeFromBaseReq) ([]*float32, error) `perm:"default"`

		RebootInstance func(p0 context.Context, p1 string, p2 string) error `perm:"default"`

		UpdateInstanceDefaultInfo func(p0 context.Context) error `perm:"default"`
	}
}

type MallStub struct {
	CommonStub

	OrderAPIStub

	UserAPIStub

	AdminAPIStub
}

type OrderAPIStruct struct {
	Internal struct {
		CancelUserOrder func(p0 context.Context, p1 string) error `perm:"user"`

		CreateOrder func(p0 context.Context, p1 types.CreateOrderReq) (string, error) `perm:"user"`

		GetUseWaitingPaymentOrders func(p0 context.Context, p1 int64, p2 int64) (*types.OrderRecordResponse, error) `perm:"user"`

		GetUserOrderRecords func(p0 context.Context, p1 int64, p2 int64) (*types.OrderRecordResponse, error) `perm:"user"`

		PaymentUserOrder func(p0 context.Context, p1 string) error `perm:"user"`

		RenewOrder func(p0 context.Context, p1 types.RenewOrderReq) (string, error) `perm:"user"`
	}
}

type OrderAPIStub struct {
}

type TransactionStruct struct {
	CommonStruct

	Internal struct {
		Hello func(p0 context.Context) error `perm:"read"`
	}
}

type TransactionStub struct {
	CommonStub
}

type UserAPIStruct struct {
	Internal struct {
		GetBalance func(p0 context.Context) (*types.UserInfo, error) `perm:"user"`

		GetInstanceDetailsInfo func(p0 context.Context, p1 string) (*types.InstanceDetails, error) `perm:"user"`

		GetRechargeAddress func(p0 context.Context) (string, error) `perm:"user"`

		GetSignCode func(p0 context.Context, p1 string) (string, error) `perm:"default"`

		GetUserInstanceRecords func(p0 context.Context, p1 int64, p2 int64) (*types.MyInstanceResponse, error) `perm:"user"`

		GetUserRechargeRecords func(p0 context.Context, p1 int64, p2 int64) (*types.RechargeResponse, error) `perm:"user"`

		GetUserWithdrawalRecords func(p0 context.Context, p1 int64, p2 int64) (*types.GetWithdrawResponse, error) `perm:"user"`

		Login func(p0 context.Context, p1 *types.UserReq) (*types.LoginResponse, error) `perm:"default"`

		Logout func(p0 context.Context, p1 *types.UserReq) error `perm:"user"`

		RebootInstance func(p0 context.Context, p1 string, p2 string) error `perm:"user"`

		UpdateInstanceName func(p0 context.Context, p1 string, p2 string) error `perm:"user"`

		Withdraw func(p0 context.Context, p1 string, p2 string) error `perm:"user"`
	}
}

type UserAPIStub struct {
}

func (s *AdminAPIStruct) AddAdminUser(p0 context.Context, p1 string, p2 string) error {
	if s.Internal.AddAdminUser == nil {
		return ErrNotSupported
	}
	return s.Internal.AddAdminUser(p0, p1, p2)
}

func (s *AdminAPIStub) AddAdminUser(p0 context.Context, p1 string, p2 string) error {
	return ErrNotSupported
}

func (s *AdminAPIStruct) ApproveUserWithdrawal(p0 context.Context, p1 string, p2 string) error {
	if s.Internal.ApproveUserWithdrawal == nil {
		return ErrNotSupported
	}
	return s.Internal.ApproveUserWithdrawal(p0, p1, p2)
}

func (s *AdminAPIStub) ApproveUserWithdrawal(p0 context.Context, p1 string, p2 string) error {
	return ErrNotSupported
}

func (s *AdminAPIStruct) GetAdminSignCode(p0 context.Context, p1 string) (string, error) {
	if s.Internal.GetAdminSignCode == nil {
		return "", ErrNotSupported
	}
	return s.Internal.GetAdminSignCode(p0, p1)
}

func (s *AdminAPIStub) GetAdminSignCode(p0 context.Context, p1 string) (string, error) {
	return "", ErrNotSupported
}

func (s *AdminAPIStruct) GetRechargeAddresses(p0 context.Context, p1 int64, p2 int64) (*types.GetRechargeAddressResponse, error) {
	if s.Internal.GetRechargeAddresses == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.GetRechargeAddresses(p0, p1, p2)
}

func (s *AdminAPIStub) GetRechargeAddresses(p0 context.Context, p1 int64, p2 int64) (*types.GetRechargeAddressResponse, error) {
	return nil, ErrNotSupported
}

func (s *AdminAPIStruct) GetWithdrawalRecords(p0 context.Context, p1 *types.GetWithdrawRequest) (*types.GetWithdrawResponse, error) {
	if s.Internal.GetWithdrawalRecords == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.GetWithdrawalRecords(p0, p1)
}

func (s *AdminAPIStub) GetWithdrawalRecords(p0 context.Context, p1 *types.GetWithdrawRequest) (*types.GetWithdrawResponse, error) {
	return nil, ErrNotSupported
}

func (s *AdminAPIStruct) LoginAdmin(p0 context.Context, p1 *types.UserReq) (*types.LoginResponse, error) {
	if s.Internal.LoginAdmin == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.LoginAdmin(p0, p1)
}

func (s *AdminAPIStub) LoginAdmin(p0 context.Context, p1 *types.UserReq) (*types.LoginResponse, error) {
	return nil, ErrNotSupported
}

func (s *AdminAPIStruct) MintToken(p0 context.Context, p1 string) (string, error) {
	if s.Internal.MintToken == nil {
		return "", ErrNotSupported
	}
	return s.Internal.MintToken(p0, p1)
}

func (s *AdminAPIStub) MintToken(p0 context.Context, p1 string) (string, error) {
	return "", ErrNotSupported
}

func (s *AdminAPIStruct) RejectUserWithdrawal(p0 context.Context, p1 string) error {
	if s.Internal.RejectUserWithdrawal == nil {
		return ErrNotSupported
	}
	return s.Internal.RejectUserWithdrawal(p0, p1)
}

func (s *AdminAPIStub) RejectUserWithdrawal(p0 context.Context, p1 string) error {
	return ErrNotSupported
}

func (s *CommonStruct) AuthNew(p0 context.Context, p1 *types.JWTPayload) (string, error) {
	if s.Internal.AuthNew == nil {
		return "", ErrNotSupported
	}
	return s.Internal.AuthNew(p0, p1)
}

func (s *CommonStub) AuthNew(p0 context.Context, p1 *types.JWTPayload) (string, error) {
	return "", ErrNotSupported
}

func (s *CommonStruct) AuthVerify(p0 context.Context, p1 string) (*types.JWTPayload, error) {
	if s.Internal.AuthVerify == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.AuthVerify(p0, p1)
}

func (s *CommonStub) AuthVerify(p0 context.Context, p1 string) (*types.JWTPayload, error) {
	return nil, ErrNotSupported
}

func (s *CommonStruct) Closing(p0 context.Context) (<-chan struct{}, error) {
	if s.Internal.Closing == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.Closing(p0)
}

func (s *CommonStub) Closing(p0 context.Context) (<-chan struct{}, error) {
	return nil, ErrNotSupported
}

func (s *CommonStruct) Discover(p0 context.Context) (types.OpenRPCDocument, error) {
	if s.Internal.Discover == nil {
		return *new(types.OpenRPCDocument), ErrNotSupported
	}
	return s.Internal.Discover(p0)
}

func (s *CommonStub) Discover(p0 context.Context) (types.OpenRPCDocument, error) {
	return *new(types.OpenRPCDocument), ErrNotSupported
}

func (s *CommonStruct) LogAlerts(p0 context.Context) ([]alerting.Alert, error) {
	if s.Internal.LogAlerts == nil {
		return *new([]alerting.Alert), ErrNotSupported
	}
	return s.Internal.LogAlerts(p0)
}

func (s *CommonStub) LogAlerts(p0 context.Context) ([]alerting.Alert, error) {
	return *new([]alerting.Alert), ErrNotSupported
}

func (s *CommonStruct) LogList(p0 context.Context) ([]string, error) {
	if s.Internal.LogList == nil {
		return *new([]string), ErrNotSupported
	}
	return s.Internal.LogList(p0)
}

func (s *CommonStub) LogList(p0 context.Context) ([]string, error) {
	return *new([]string), ErrNotSupported
}

func (s *CommonStruct) LogSetLevel(p0 context.Context, p1 string, p2 string) error {
	if s.Internal.LogSetLevel == nil {
		return ErrNotSupported
	}
	return s.Internal.LogSetLevel(p0, p1, p2)
}

func (s *CommonStub) LogSetLevel(p0 context.Context, p1 string, p2 string) error {
	return ErrNotSupported
}

func (s *CommonStruct) Session(p0 context.Context) (uuid.UUID, error) {
	if s.Internal.Session == nil {
		return *new(uuid.UUID), ErrNotSupported
	}
	return s.Internal.Session(p0)
}

func (s *CommonStub) Session(p0 context.Context) (uuid.UUID, error) {
	return *new(uuid.UUID), ErrNotSupported
}

func (s *CommonStruct) Shutdown(p0 context.Context) error {
	if s.Internal.Shutdown == nil {
		return ErrNotSupported
	}
	return s.Internal.Shutdown(p0)
}

func (s *CommonStub) Shutdown(p0 context.Context) error {
	return ErrNotSupported
}

func (s *CommonStruct) Version(p0 context.Context) (APIVersion, error) {
	if s.Internal.Version == nil {
		return *new(APIVersion), ErrNotSupported
	}
	return s.Internal.Version(p0)
}

func (s *CommonStub) Version(p0 context.Context) (APIVersion, error) {
	return *new(APIVersion), ErrNotSupported
}

func (s *MallStruct) AttachKeyPair(p0 context.Context, p1 string, p2 string, p3 []string) ([]*types.AttachKeyPairResponse, error) {
	if s.Internal.AttachKeyPair == nil {
		return *new([]*types.AttachKeyPairResponse), ErrNotSupported
	}
	return s.Internal.AttachKeyPair(p0, p1, p2, p3)
}

func (s *MallStub) AttachKeyPair(p0 context.Context, p1 string, p2 string, p3 []string) ([]*types.AttachKeyPairResponse, error) {
	return *new([]*types.AttachKeyPairResponse), ErrNotSupported
}

func (s *MallStruct) CreateInstance(p0 context.Context, p1 *types.CreateInstanceReq) (*types.CreateInstanceResponse, error) {
	if s.Internal.CreateInstance == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.CreateInstance(p0, p1)
}

func (s *MallStub) CreateInstance(p0 context.Context, p1 *types.CreateInstanceReq) (*types.CreateInstanceResponse, error) {
	return nil, ErrNotSupported
}

func (s *MallStruct) CreateKeyPair(p0 context.Context, p1 string, p2 string) (*types.CreateKeyPairResponse, error) {
	if s.Internal.CreateKeyPair == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.CreateKeyPair(p0, p1, p2)
}

func (s *MallStub) CreateKeyPair(p0 context.Context, p1 string, p2 string) (*types.CreateKeyPairResponse, error) {
	return nil, ErrNotSupported
}

func (s *MallStruct) DescribeAvailableResourceForDesk(p0 context.Context, p1 *types.AvailableResourceReq) ([]*types.AvailableResourceResponse, error) {
	if s.Internal.DescribeAvailableResourceForDesk == nil {
		return *new([]*types.AvailableResourceResponse), ErrNotSupported
	}
	return s.Internal.DescribeAvailableResourceForDesk(p0, p1)
}

func (s *MallStub) DescribeAvailableResourceForDesk(p0 context.Context, p1 *types.AvailableResourceReq) ([]*types.AvailableResourceResponse, error) {
	return *new([]*types.AvailableResourceResponse), ErrNotSupported
}

func (s *MallStruct) DescribeImages(p0 context.Context, p1 string, p2 string) ([]*types.DescribeImageResponse, error) {
	if s.Internal.DescribeImages == nil {
		return *new([]*types.DescribeImageResponse), ErrNotSupported
	}
	return s.Internal.DescribeImages(p0, p1, p2)
}

func (s *MallStub) DescribeImages(p0 context.Context, p1 string, p2 string) ([]*types.DescribeImageResponse, error) {
	return *new([]*types.DescribeImageResponse), ErrNotSupported
}

func (s *MallStruct) DescribeInstanceType(p0 context.Context, p1 *types.DescribeInstanceTypeReq) (*types.DescribeInstanceTypeResponse, error) {
	if s.Internal.DescribeInstanceType == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.DescribeInstanceType(p0, p1)
}

func (s *MallStub) DescribeInstanceType(p0 context.Context, p1 *types.DescribeInstanceTypeReq) (*types.DescribeInstanceTypeResponse, error) {
	return nil, ErrNotSupported
}

func (s *MallStruct) DescribeInstances(p0 context.Context, p1 string, p2 string) error {
	if s.Internal.DescribeInstances == nil {
		return ErrNotSupported
	}
	return s.Internal.DescribeInstances(p0, p1, p2)
}

func (s *MallStub) DescribeInstances(p0 context.Context, p1 string, p2 string) error {
	return ErrNotSupported
}

func (s *MallStruct) DescribePrice(p0 context.Context, p1 *types.DescribePriceReq) (*types.DescribePriceResponse, error) {
	if s.Internal.DescribePrice == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.DescribePrice(p0, p1)
}

func (s *MallStub) DescribePrice(p0 context.Context, p1 *types.DescribePriceReq) (*types.DescribePriceResponse, error) {
	return nil, ErrNotSupported
}

func (s *MallStruct) DescribeRecommendInstanceType(p0 context.Context, p1 *types.DescribeRecommendInstanceTypeReq) ([]*types.DescribeRecommendInstanceResponse, error) {
	if s.Internal.DescribeRecommendInstanceType == nil {
		return *new([]*types.DescribeRecommendInstanceResponse), ErrNotSupported
	}
	return s.Internal.DescribeRecommendInstanceType(p0, p1)
}

func (s *MallStub) DescribeRecommendInstanceType(p0 context.Context, p1 *types.DescribeRecommendInstanceTypeReq) ([]*types.DescribeRecommendInstanceResponse, error) {
	return *new([]*types.DescribeRecommendInstanceResponse), ErrNotSupported
}

func (s *MallStruct) DescribeRegions(p0 context.Context) (map[string]string, error) {
	if s.Internal.DescribeRegions == nil {
		return *new(map[string]string), ErrNotSupported
	}
	return s.Internal.DescribeRegions(p0)
}

func (s *MallStub) DescribeRegions(p0 context.Context) (map[string]string, error) {
	return *new(map[string]string), ErrNotSupported
}

func (s *MallStruct) GetInstanceCpuInfo(p0 context.Context, p1 *types.InstanceTypeFromBaseReq) ([]*int32, error) {
	if s.Internal.GetInstanceCpuInfo == nil {
		return *new([]*int32), ErrNotSupported
	}
	return s.Internal.GetInstanceCpuInfo(p0, p1)
}

func (s *MallStub) GetInstanceCpuInfo(p0 context.Context, p1 *types.InstanceTypeFromBaseReq) ([]*int32, error) {
	return *new([]*int32), ErrNotSupported
}

func (s *MallStruct) GetInstanceDefaultInfo(p0 context.Context, p1 *types.InstanceTypeFromBaseReq) (*types.InstanceTypeResponse, error) {
	if s.Internal.GetInstanceDefaultInfo == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.GetInstanceDefaultInfo(p0, p1)
}

func (s *MallStub) GetInstanceDefaultInfo(p0 context.Context, p1 *types.InstanceTypeFromBaseReq) (*types.InstanceTypeResponse, error) {
	return nil, ErrNotSupported
}

func (s *MallStruct) GetInstanceMemoryInfo(p0 context.Context, p1 *types.InstanceTypeFromBaseReq) ([]*float32, error) {
	if s.Internal.GetInstanceMemoryInfo == nil {
		return *new([]*float32), ErrNotSupported
	}
	return s.Internal.GetInstanceMemoryInfo(p0, p1)
}

func (s *MallStub) GetInstanceMemoryInfo(p0 context.Context, p1 *types.InstanceTypeFromBaseReq) ([]*float32, error) {
	return *new([]*float32), ErrNotSupported
}

func (s *MallStruct) RebootInstance(p0 context.Context, p1 string, p2 string) error {
	if s.Internal.RebootInstance == nil {
		return ErrNotSupported
	}
	return s.Internal.RebootInstance(p0, p1, p2)
}

func (s *MallStub) RebootInstance(p0 context.Context, p1 string, p2 string) error {
	return ErrNotSupported
}

func (s *MallStruct) UpdateInstanceDefaultInfo(p0 context.Context) error {
	if s.Internal.UpdateInstanceDefaultInfo == nil {
		return ErrNotSupported
	}
	return s.Internal.UpdateInstanceDefaultInfo(p0)
}

func (s *MallStub) UpdateInstanceDefaultInfo(p0 context.Context) error {
	return ErrNotSupported
}

func (s *OrderAPIStruct) CancelUserOrder(p0 context.Context, p1 string) error {
	if s.Internal.CancelUserOrder == nil {
		return ErrNotSupported
	}
	return s.Internal.CancelUserOrder(p0, p1)
}

func (s *OrderAPIStub) CancelUserOrder(p0 context.Context, p1 string) error {
	return ErrNotSupported
}

func (s *OrderAPIStruct) CreateOrder(p0 context.Context, p1 types.CreateOrderReq) (string, error) {
	if s.Internal.CreateOrder == nil {
		return "", ErrNotSupported
	}
	return s.Internal.CreateOrder(p0, p1)
}

func (s *OrderAPIStub) CreateOrder(p0 context.Context, p1 types.CreateOrderReq) (string, error) {
	return "", ErrNotSupported
}

func (s *OrderAPIStruct) GetUseWaitingPaymentOrders(p0 context.Context, p1 int64, p2 int64) (*types.OrderRecordResponse, error) {
	if s.Internal.GetUseWaitingPaymentOrders == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.GetUseWaitingPaymentOrders(p0, p1, p2)
}

func (s *OrderAPIStub) GetUseWaitingPaymentOrders(p0 context.Context, p1 int64, p2 int64) (*types.OrderRecordResponse, error) {
	return nil, ErrNotSupported
}

func (s *OrderAPIStruct) GetUserOrderRecords(p0 context.Context, p1 int64, p2 int64) (*types.OrderRecordResponse, error) {
	if s.Internal.GetUserOrderRecords == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.GetUserOrderRecords(p0, p1, p2)
}

func (s *OrderAPIStub) GetUserOrderRecords(p0 context.Context, p1 int64, p2 int64) (*types.OrderRecordResponse, error) {
	return nil, ErrNotSupported
}

func (s *OrderAPIStruct) PaymentUserOrder(p0 context.Context, p1 string) error {
	if s.Internal.PaymentUserOrder == nil {
		return ErrNotSupported
	}
	return s.Internal.PaymentUserOrder(p0, p1)
}

func (s *OrderAPIStub) PaymentUserOrder(p0 context.Context, p1 string) error {
	return ErrNotSupported
}

func (s *OrderAPIStruct) RenewOrder(p0 context.Context, p1 types.RenewOrderReq) (string, error) {
	if s.Internal.RenewOrder == nil {
		return "", ErrNotSupported
	}
	return s.Internal.RenewOrder(p0, p1)
}

func (s *OrderAPIStub) RenewOrder(p0 context.Context, p1 types.RenewOrderReq) (string, error) {
	return "", ErrNotSupported
}

func (s *TransactionStruct) Hello(p0 context.Context) error {
	if s.Internal.Hello == nil {
		return ErrNotSupported
	}
	return s.Internal.Hello(p0)
}

func (s *TransactionStub) Hello(p0 context.Context) error {
	return ErrNotSupported
}

func (s *UserAPIStruct) GetBalance(p0 context.Context) (*types.UserInfo, error) {
	if s.Internal.GetBalance == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.GetBalance(p0)
}

func (s *UserAPIStub) GetBalance(p0 context.Context) (*types.UserInfo, error) {
	return nil, ErrNotSupported
}

func (s *UserAPIStruct) GetInstanceDetailsInfo(p0 context.Context, p1 string) (*types.InstanceDetails, error) {
	if s.Internal.GetInstanceDetailsInfo == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.GetInstanceDetailsInfo(p0, p1)
}

func (s *UserAPIStub) GetInstanceDetailsInfo(p0 context.Context, p1 string) (*types.InstanceDetails, error) {
	return nil, ErrNotSupported
}

func (s *UserAPIStruct) GetRechargeAddress(p0 context.Context) (string, error) {
	if s.Internal.GetRechargeAddress == nil {
		return "", ErrNotSupported
	}
	return s.Internal.GetRechargeAddress(p0)
}

func (s *UserAPIStub) GetRechargeAddress(p0 context.Context) (string, error) {
	return "", ErrNotSupported
}

func (s *UserAPIStruct) GetSignCode(p0 context.Context, p1 string) (string, error) {
	if s.Internal.GetSignCode == nil {
		return "", ErrNotSupported
	}
	return s.Internal.GetSignCode(p0, p1)
}

func (s *UserAPIStub) GetSignCode(p0 context.Context, p1 string) (string, error) {
	return "", ErrNotSupported
}

func (s *UserAPIStruct) GetUserInstanceRecords(p0 context.Context, p1 int64, p2 int64) (*types.MyInstanceResponse, error) {
	if s.Internal.GetUserInstanceRecords == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.GetUserInstanceRecords(p0, p1, p2)
}

func (s *UserAPIStub) GetUserInstanceRecords(p0 context.Context, p1 int64, p2 int64) (*types.MyInstanceResponse, error) {
	return nil, ErrNotSupported
}

func (s *UserAPIStruct) GetUserRechargeRecords(p0 context.Context, p1 int64, p2 int64) (*types.RechargeResponse, error) {
	if s.Internal.GetUserRechargeRecords == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.GetUserRechargeRecords(p0, p1, p2)
}

func (s *UserAPIStub) GetUserRechargeRecords(p0 context.Context, p1 int64, p2 int64) (*types.RechargeResponse, error) {
	return nil, ErrNotSupported
}

func (s *UserAPIStruct) GetUserWithdrawalRecords(p0 context.Context, p1 int64, p2 int64) (*types.GetWithdrawResponse, error) {
	if s.Internal.GetUserWithdrawalRecords == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.GetUserWithdrawalRecords(p0, p1, p2)
}

func (s *UserAPIStub) GetUserWithdrawalRecords(p0 context.Context, p1 int64, p2 int64) (*types.GetWithdrawResponse, error) {
	return nil, ErrNotSupported
}

func (s *UserAPIStruct) Login(p0 context.Context, p1 *types.UserReq) (*types.LoginResponse, error) {
	if s.Internal.Login == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.Login(p0, p1)
}

func (s *UserAPIStub) Login(p0 context.Context, p1 *types.UserReq) (*types.LoginResponse, error) {
	return nil, ErrNotSupported
}

func (s *UserAPIStruct) Logout(p0 context.Context, p1 *types.UserReq) error {
	if s.Internal.Logout == nil {
		return ErrNotSupported
	}
	return s.Internal.Logout(p0, p1)
}

func (s *UserAPIStub) Logout(p0 context.Context, p1 *types.UserReq) error {
	return ErrNotSupported
}

func (s *UserAPIStruct) RebootInstance(p0 context.Context, p1 string, p2 string) error {
	if s.Internal.RebootInstance == nil {
		return ErrNotSupported
	}
	return s.Internal.RebootInstance(p0, p1, p2)
}

func (s *UserAPIStub) RebootInstance(p0 context.Context, p1 string, p2 string) error {
	return ErrNotSupported
}

func (s *UserAPIStruct) UpdateInstanceName(p0 context.Context, p1 string, p2 string) error {
	if s.Internal.UpdateInstanceName == nil {
		return ErrNotSupported
	}
	return s.Internal.UpdateInstanceName(p0, p1, p2)
}

func (s *UserAPIStub) UpdateInstanceName(p0 context.Context, p1 string, p2 string) error {
	return ErrNotSupported
}

func (s *UserAPIStruct) Withdraw(p0 context.Context, p1 string, p2 string) error {
	if s.Internal.Withdraw == nil {
		return ErrNotSupported
	}
	return s.Internal.Withdraw(p0, p1, p2)
}

func (s *UserAPIStub) Withdraw(p0 context.Context, p1 string, p2 string) error {
	return ErrNotSupported
}

var _ AdminAPI = new(AdminAPIStruct)
var _ Common = new(CommonStruct)
var _ Mall = new(MallStruct)
var _ OrderAPI = new(OrderAPIStruct)
var _ Transaction = new(TransactionStruct)
var _ UserAPI = new(UserAPIStruct)
