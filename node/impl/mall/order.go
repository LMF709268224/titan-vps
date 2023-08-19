package mall

import (
	"context"
	"strconv"
	"strings"

	"github.com/LMF709268224/titan-vps/node/utils"
	"github.com/google/uuid"

	"github.com/LMF709268224/titan-vps/api/types"
	"github.com/LMF709268224/titan-vps/node/handler"
)

func (m *Mall) CreateOrder(ctx context.Context, req types.CreateOrderReq) (string, error) {
	userID := handler.GetID(ctx)
	priceReq := &types.DescribePriceReq{
		RegionId:                req.RegionId,
		InstanceType:            req.InstanceType,
		PriceUnit:               req.PeriodUnit,
		Period:                  req.Period,
		Amount:                  req.Amount,
		InternetChargeType:      req.InternetChargeType,
		ImageID:                 req.ImageId,
		InternetMaxBandwidthOut: req.InternetMaxBandwidthOut,
		SystemDiskCategory:      req.SystemDiskCategory,
		SystemDiskSize:          req.SystemDiskSize,
	}
	priceInfo, err := m.DescribePrice(ctx, priceReq)
	if err != nil {
		log.Errorf("DescribePrice:%v", err)
		return "", err
	}

	hash := uuid.NewString()
	orderID := strings.Replace(hash, "-", "", -1)
	req.OrderID = orderID
	req.TradePrice = priceInfo.USDPrice / float32(req.Amount)
	//for i := int32(0); i < req.Amount; i++ {
	//	id, err = m.SaveVpsInstance(&req)
	//	if err != nil {
	//		log.Errorf("SaveVpsInstance:%v", err)
	//	}
	//}
	id, err := m.SaveVpsInstance(&req)
	if err != nil {
		log.Errorf("SaveVpsInstance:%v", err)
	}
	TradePriceString := strconv.FormatFloat(float64(req.TradePrice), 'f', -1, 64)
	info := &types.OrderRecord{
		VpsID:      id,
		OrderID:    orderID,
		UserID:     userID,
		Value:      "10000000000",
		TradePrice: TradePriceString,
	}
	oldBalance, err := m.LoadUserBalance(userID)
	if err != nil {
		log.Errorf("LoadUserBalance:%v", err)
	}
	newBalanceString := strconv.FormatFloat(float64(priceInfo.USDPrice)*1000000000000000000, 'f', -1, 64)
	newBalanceString, ok := utils.BigIntReduce(oldBalance, newBalanceString)
	if ok {
		err = m.UpdateUserBalance(userID, newBalanceString, oldBalance)
		if err != nil {
			log.Errorf("UpdateUserBalance:%v", err)
			return "", err
		}
	}
	err = m.OrderMgr.CreatedOrder(info)
	if err != nil {
		return "", err
	}

	return info.To, nil
}

func (m *Mall) GetOrderWaitingPayment(ctx context.Context, limit, offset int64) (*types.OrderRecordResponse, error) {
	userID := handler.GetID(ctx)

	info, err := m.LoadOrderRecordByUserUndone(userID, limit, offset)
	if err != nil {
		return nil, err
	}

	return info, nil
}

func (m *Mall) GetOrderInfo(ctx context.Context, limit, offset int64) (*types.OrderRecordResponse, error) {
	userID := handler.GetID(ctx)

	info, err := m.LoadOrderRecordByUserAll(userID, limit, offset)
	if err != nil {
		return nil, err
	}

	return info, nil
}

func (m *Mall) CancelOrder(ctx context.Context, orderID string) error {
	return m.OrderMgr.CancelOrder(orderID)
}

func (m *Mall) PaymentCompleted(ctx context.Context, orderID string) error {
	return m.OrderMgr.PaymentCompleted(orderID)
}
