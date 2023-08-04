package types

import (
	"time"
)

type Hellos struct {
	Msg string
}

// User user info
type User struct {
	UUID      string    `db:"uuid" json:"uuid"`
	UserName  string    `db:"user_name" json:"user_name"`
	PassHash  string    `db:"pass_hash" json:"pass_hash"`
	Address   string    `db:"address" json:"address"`
	Public    string    `db:"public" json:"public"`
	Token     string    `db:"token" json:"token"`
	Role      int32     `db:"role" json:"role"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type DescribePriceResponse struct {
	Currency      string
	OriginalPrice float32
	TradePrice    float32
}

// todo
type CreateInstanceReq struct {
	Id                      string `db:"id"`
	RegionId                string `db:"region_id"`
	InstanceType            string `db:"instance_type"`
	DryRun                  bool   `db:"dry_run"`
	ImageId                 string `db:"image_id"`
	SecurityGroupId         string `db:"security_group_id"`
	InstanceChargeType      string `db:"instanceCharge_type"`
	PeriodUnit              string `db:"period_unit"`
	Period                  int32  `db:"period"`
	InternetMaxBandwidthOut int32  `db:"bandwidth_out"`
	InternetMaxBandwidthIn  int32  `db:"bandwidth_in"`
}

type CreateInstanceResponse struct {
	InstanceId      string
	OrderId         string
	RequestId       string
	TradePrice      float32
	PublicIpAddress string
	PrivateKey      string
}

type CreateKeyPairResponse struct {
	KeyPairId      string
	KeyPairName    string
	PrivateKeyBody string
}
type AttachKeyPairResponse struct {
	Code       string
	InstanceId string
	Message    string
	Success    string
}

// OrderRecord represents information about an order record
type OrderRecord struct {
	OrderID       string    `db:"order_id"`
	From          string    `db:"from_addr"`
	To            string    `db:"to_addr"`
	Value         int64     `db:"value"`
	State         int64     `db:"state"`
	DoneState     int64     `db:"done_state"`
	CreatedHeight int64     `db:"created_height"`
	CreatedTime   time.Time `db:"created_time"`
	DoneTime      time.Time `db:"done_time"`
	DoneHeight    int64     `db:"done_height"`
	VpsID         int64     `db:"vps_id"`
}

type CreateOrderReq struct {
	Vps  CreateInstanceReq
	User string
}

type PaymentCompletedReq struct {
	OrderID       string
	TransactionID string
}
type UserReq struct {
	UserId    string
	Signature string
	Address   string
	PublicKey string
	Token     string
}
type UserResponse struct {
	UserId   string
	SignCode string
	Token    string
}

type UserInfoTmp struct {
	UserLogin UserResponse
	OrderInfo OrderRecord
}
type Token struct {
	TokenString string
	UserId      string
	Expiration  time.Time
}

// EventTopics represents topics for pub/sub events
type EventTopics string

const (
	// EventTransfer node online event
	EventTransfer EventTopics = "transfer"
)

func (t EventTopics) String() string {
	return string(t)
}

type FvmTransfer struct {
	From  string
	To    string
	Value int64
}
