// Code generated by titan/gen/api. DO NOT EDIT.

package api

import (
	"context"
	"github.com/LMF709268224/titan-vps/api/types"
	"github.com/LMF709268224/titan-vps/journal/alerting"
	"github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

var ErrNotSupported = xerrors.New("method not supported")

type BasisStruct struct {
	CommonStruct

	Internal struct {
		AttachKeyPair func(p0 context.Context, p1 string, p2 string, p3 []string) ([]*types.AttachKeyPairResponse, error) `perm:"read"`

		CancelOrder func(p0 context.Context, p1 string) error `perm:"read"`

		CreateInstance func(p0 context.Context, p1 string, p2 string, p3 string, p4 string, p5 int32) (*types.CreateInstanceResponse, error) `perm:"read"`

		CreateKeyPair func(p0 context.Context, p1 string, p2 string) (*types.CreateKeyPairResponse, error) `perm:"read"`

		CreateOrder func(p0 context.Context, p1 types.CreateOrderReq) (string, error) `perm:"read"`

		DescribeImages func(p0 context.Context, p1 string, p2 string) ([]string, error) `perm:"read"`

		DescribeInstanceType func(p0 context.Context, p1 string, p2 int32, p3 float32) ([]string, error) `perm:"read"`

		DescribePrice func(p0 context.Context, p1 string, p2 string, p3 string, p4 string, p5 int32) (*types.DescribePriceResponse, error) `perm:"read"`

		DescribeRegions func(p0 context.Context) ([]string, error) `perm:"read"`

		Login func(p0 context.Context, p1 *types.UserReq) (*types.UserResponse, error) `perm:"read"`

		Logout func(p0 context.Context, p1 *types.UserReq) error `perm:"read"`

		PaymentCompleted func(p0 context.Context, p1 types.PaymentCompletedReq) (string, error) `perm:"read"`

		RebootInstance func(p0 context.Context, p1 string, p2 string) (string, error) `perm:"read"`

		SignCode func(p0 context.Context, p1 string) (string, error) `perm:"read"`
	}
}

type BasisStub struct {
	CommonStub
}

type CommonStruct struct {
	Internal struct {
		AuthNew func(p0 context.Context, p1 []auth.Permission) ([]byte, error) `perm:"admin"`

		AuthVerify func(p0 context.Context, p1 string) ([]auth.Permission, error) `perm:"read"`

		Closing func(p0 context.Context) (<-chan struct{}, error) `perm:"admin"`

		Discover func(p0 context.Context) (types.OpenRPCDocument, error) `perm:"admin"`

		LogAlerts func(p0 context.Context) ([]alerting.Alert, error) `perm:"admin"`

		LogList func(p0 context.Context) ([]string, error) `perm:"admin"`

		LogSetLevel func(p0 context.Context, p1 string, p2 string) error `perm:"admin"`

		Session func(p0 context.Context) (uuid.UUID, error) `perm:"admin"`

		Shutdown func(p0 context.Context) error `perm:"admin"`

		Version func(p0 context.Context) (APIVersion, error) `perm:"read"`
	}
}

type CommonStub struct {
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

func (s *BasisStruct) AttachKeyPair(p0 context.Context, p1 string, p2 string, p3 []string) ([]*types.AttachKeyPairResponse, error) {
	if s.Internal.AttachKeyPair == nil {
		return *new([]*types.AttachKeyPairResponse), ErrNotSupported
	}
	return s.Internal.AttachKeyPair(p0, p1, p2, p3)
}

func (s *BasisStub) AttachKeyPair(p0 context.Context, p1 string, p2 string, p3 []string) ([]*types.AttachKeyPairResponse, error) {
	return *new([]*types.AttachKeyPairResponse), ErrNotSupported
}

func (s *BasisStruct) CancelOrder(p0 context.Context, p1 string) error {
	if s.Internal.CancelOrder == nil {
		return ErrNotSupported
	}
	return s.Internal.CancelOrder(p0, p1)
}

func (s *BasisStub) CancelOrder(p0 context.Context, p1 string) error {
	return ErrNotSupported
}

func (s *BasisStruct) CreateInstance(p0 context.Context, p1 string, p2 string, p3 string, p4 string, p5 int32) (*types.CreateInstanceResponse, error) {
	if s.Internal.CreateInstance == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.CreateInstance(p0, p1, p2, p3, p4, p5)
}

func (s *BasisStub) CreateInstance(p0 context.Context, p1 string, p2 string, p3 string, p4 string, p5 int32) (*types.CreateInstanceResponse, error) {
	return nil, ErrNotSupported
}

func (s *BasisStruct) CreateKeyPair(p0 context.Context, p1 string, p2 string) (*types.CreateKeyPairResponse, error) {
	if s.Internal.CreateKeyPair == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.CreateKeyPair(p0, p1, p2)
}

func (s *BasisStub) CreateKeyPair(p0 context.Context, p1 string, p2 string) (*types.CreateKeyPairResponse, error) {
	return nil, ErrNotSupported
}

func (s *BasisStruct) CreateOrder(p0 context.Context, p1 types.CreateOrderReq) (string, error) {
	if s.Internal.CreateOrder == nil {
		return "", ErrNotSupported
	}
	return s.Internal.CreateOrder(p0, p1)
}

func (s *BasisStub) CreateOrder(p0 context.Context, p1 types.CreateOrderReq) (string, error) {
	return "", ErrNotSupported
}

func (s *BasisStruct) DescribeImages(p0 context.Context, p1 string, p2 string) ([]string, error) {
	if s.Internal.DescribeImages == nil {
		return *new([]string), ErrNotSupported
	}
	return s.Internal.DescribeImages(p0, p1, p2)
}

func (s *BasisStub) DescribeImages(p0 context.Context, p1 string, p2 string) ([]string, error) {
	return *new([]string), ErrNotSupported
}

func (s *BasisStruct) DescribeInstanceType(p0 context.Context, p1 string, p2 int32, p3 float32) ([]string, error) {
	if s.Internal.DescribeInstanceType == nil {
		return *new([]string), ErrNotSupported
	}
	return s.Internal.DescribeInstanceType(p0, p1, p2, p3)
}

func (s *BasisStub) DescribeInstanceType(p0 context.Context, p1 string, p2 int32, p3 float32) ([]string, error) {
	return *new([]string), ErrNotSupported
}

func (s *BasisStruct) DescribePrice(p0 context.Context, p1 string, p2 string, p3 string, p4 string, p5 int32) (*types.DescribePriceResponse, error) {
	if s.Internal.DescribePrice == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.DescribePrice(p0, p1, p2, p3, p4, p5)
}

func (s *BasisStub) DescribePrice(p0 context.Context, p1 string, p2 string, p3 string, p4 string, p5 int32) (*types.DescribePriceResponse, error) {
	return nil, ErrNotSupported
}

func (s *BasisStruct) DescribeRegions(p0 context.Context) ([]string, error) {
	if s.Internal.DescribeRegions == nil {
		return *new([]string), ErrNotSupported
	}
	return s.Internal.DescribeRegions(p0)
}

func (s *BasisStub) DescribeRegions(p0 context.Context) ([]string, error) {
	return *new([]string), ErrNotSupported
}

func (s *BasisStruct) Login(p0 context.Context, p1 *types.UserReq) (*types.UserResponse, error) {
	if s.Internal.Login == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.Login(p0, p1)
}

func (s *BasisStub) Login(p0 context.Context, p1 *types.UserReq) (*types.UserResponse, error) {
	return nil, ErrNotSupported
}

func (s *BasisStruct) Logout(p0 context.Context, p1 *types.UserReq) error {
	if s.Internal.Logout == nil {
		return ErrNotSupported
	}
	return s.Internal.Logout(p0, p1)
}

func (s *BasisStub) Logout(p0 context.Context, p1 *types.UserReq) error {
	return ErrNotSupported
}

func (s *BasisStruct) PaymentCompleted(p0 context.Context, p1 types.PaymentCompletedReq) (string, error) {
	if s.Internal.PaymentCompleted == nil {
		return "", ErrNotSupported
	}
	return s.Internal.PaymentCompleted(p0, p1)
}

func (s *BasisStub) PaymentCompleted(p0 context.Context, p1 types.PaymentCompletedReq) (string, error) {
	return "", ErrNotSupported
}

func (s *BasisStruct) RebootInstance(p0 context.Context, p1 string, p2 string) (string, error) {
	if s.Internal.RebootInstance == nil {
		return "", ErrNotSupported
	}
	return s.Internal.RebootInstance(p0, p1, p2)
}

func (s *BasisStub) RebootInstance(p0 context.Context, p1 string, p2 string) (string, error) {
	return "", ErrNotSupported
}

func (s *BasisStruct) SignCode(p0 context.Context, p1 string) (string, error) {
	if s.Internal.SignCode == nil {
		return "", ErrNotSupported
	}
	return s.Internal.SignCode(p0, p1)
}

func (s *BasisStub) SignCode(p0 context.Context, p1 string) (string, error) {
	return "", ErrNotSupported
}

func (s *CommonStruct) AuthNew(p0 context.Context, p1 []auth.Permission) ([]byte, error) {
	if s.Internal.AuthNew == nil {
		return *new([]byte), ErrNotSupported
	}
	return s.Internal.AuthNew(p0, p1)
}

func (s *CommonStub) AuthNew(p0 context.Context, p1 []auth.Permission) ([]byte, error) {
	return *new([]byte), ErrNotSupported
}

func (s *CommonStruct) AuthVerify(p0 context.Context, p1 string) ([]auth.Permission, error) {
	if s.Internal.AuthVerify == nil {
		return *new([]auth.Permission), ErrNotSupported
	}
	return s.Internal.AuthVerify(p0, p1)
}

func (s *CommonStub) AuthVerify(p0 context.Context, p1 string) ([]auth.Permission, error) {
	return *new([]auth.Permission), ErrNotSupported
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

func (s *TransactionStruct) Hello(p0 context.Context) error {
	if s.Internal.Hello == nil {
		return ErrNotSupported
	}
	return s.Internal.Hello(p0)
}

func (s *TransactionStub) Hello(p0 context.Context) error {
	return ErrNotSupported
}

var _ Basis = new(BasisStruct)
var _ Common = new(CommonStruct)
var _ Transaction = new(TransactionStruct)
