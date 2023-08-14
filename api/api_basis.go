package api

import (
	"context"

	"github.com/LMF709268224/titan-vps/api/types"
)

type Basis interface {
	Common
	OrderAPI
	UserAPI

	MintToken(ctx context.Context, address string) (string, error) //perm:admin

	DescribeRegions(ctx context.Context) ([]string, error)                                                                                 //perm:default
	DescribeInstanceType(ctx context.Context, instanceTypeReq *types.DescribeInstanceTypeReq) (*types.DescribeInstanceTypeResponse, error) //perm:default
	DescribeImages(ctx context.Context, regionID, instanceType string) ([]*types.DescribeImageResponse, error)                             //perm:default
	DescribePrice(ctx context.Context, describePriceReq *types.DescribePriceReq) (*types.DescribePriceResponse, error)                     //perm:default
	CreateInstance(ctx context.Context, vpsInfo *types.CreateInstanceReq) (*types.CreateInstanceResponse, error)                           //perm:default
	CreateKeyPair(ctx context.Context, regionID, KeyPairName string) (*types.CreateKeyPairResponse, error)                                 //perm:default
	AttachKeyPair(ctx context.Context, regionID, KeyPairName string, instanceIds []string) ([]*types.AttachKeyPairResponse, error)         //perm:default
	RebootInstance(ctx context.Context, regionID, instanceID string) (string, error)                                                       //perm:default
}

// OrderAPI is an interface for order
type OrderAPI interface {
	// order
	CreateOrder(ctx context.Context, req types.CreateInstanceReq) (string, error)        //perm:user
	PaymentCompleted(ctx context.Context, req types.PaymentCompletedReq) (string, error) //perm:user
	CancelOrder(ctx context.Context, orderID string) error                               //perm:user
}

// UserAPI is an interface for user
type UserAPI interface {
	// user
	GetBalance(ctx context.Context) (string, error)                                              //perm:user
	RebootInstance(ctx context.Context, regionID, instanceID string) (string, error)             //perm:user
	SignCode(ctx context.Context, userID string) (string, error)                                 //perm:default
	Login(ctx context.Context, user *types.UserReq) (*types.UserResponse, error)                 //perm:default
	Logout(ctx context.Context, user *types.UserReq) error                                       //perm:user
	Recharge(ctx context.Context) (string, error)                                                //perm:user
	Withdraw(ctx context.Context, withdrawAddr string) (string, error)                           //perm:user
	GetRechargeRecord(ctx context.Context, limit, offset int64) (*types.RechargeResponse, error) //perm:user
	GetWithdrawRecord(ctx context.Context, limit, offset int64) (*types.WithdrawResponse, error) //perm:user
}
