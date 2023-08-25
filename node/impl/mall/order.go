package mall

import (
	"context"
	"encoding/json"
	"math"
	"strconv"
	"strings"

	"github.com/google/uuid"

	"github.com/LMF709268224/titan-vps/api"
	"github.com/LMF709268224/titan-vps/api/terrors"
	"github.com/LMF709268224/titan-vps/api/types"
	"github.com/LMF709268224/titan-vps/node/handler"
)

func (m *Mall) CreateOrder(ctx context.Context, req types.CreateOrderReq) (string, error) {
	userID := handler.GetID(ctx)
	if len(req.DataDisk) > 0 {
		dataDisk, err := json.Marshal(req.DataDisk)
		if err != nil {
			log.Errorf("Marshal DataDisk:%v", err)
			return "", &api.ErrWeb{Code: terrors.ParametersWrong.Int(), Message: err.Error()}
		}
		req.DataDiskString = string(dataDisk)
	}
	priceReq := &types.DescribePriceReq{
		RegionId:                     req.RegionId,
		InstanceType:                 req.InstanceType,
		PriceUnit:                    req.PeriodUnit,
		Period:                       req.Period,
		Amount:                       req.Amount,
		InternetChargeType:           req.InternetChargeType,
		ImageID:                      req.ImageId,
		InternetMaxBandwidthOut:      req.InternetMaxBandwidthOut,
		SystemDiskCategory:           req.SystemDiskCategory,
		SystemDiskSize:               req.SystemDiskSize,
		DescribePriceRequestDataDisk: req.DataDisk,
	}
	priceInfo, err := m.DescribePrice(ctx, priceReq)
	if err != nil {
		log.Errorf("DescribePrice:%v", err)
		return "", &api.ErrWeb{Code: terrors.DescribePriceError.Int(), Message: err.Error()}
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
		return "", err
	}
	newBalanceString := strconv.FormatFloat(math.Ceil(float64(priceInfo.USDPrice)*1000000), 'f', 0, 64)

	info := &types.OrderRecord{
		VpsID:     id,
		OrderID:   orderID,
		UserID:    userID,
		Value:     newBalanceString,
		OrderType: types.BuyVPS,
	}

	err = m.OrderMgr.CreatedOrder(info)
	if err != nil {
		return "", err
	}

	return orderID, nil
}

func (m *Mall) RenewOrder(ctx context.Context, renewReq types.RenewOrderReq) (string, error) {
	userID := handler.GetID(ctx)
	req, err := m.LoadVpsInfoByInstanceId(renewReq.InstanceId)
	if len(req.DataDisk) > 0 {
		dataDisk, err := json.Marshal(req.DataDisk)
		if err != nil {
			log.Errorf("Marshal DataDisk:%v", err)
			return "", &api.ErrWeb{Code: terrors.ParametersWrong.Int(), Message: err.Error()}
		}
		req.DataDiskString = string(dataDisk)
	}
	priceReq := &types.DescribePriceReq{
		RegionId:                     req.RegionId,
		InstanceType:                 req.InstanceType,
		PriceUnit:                    renewReq.PeriodUnit,
		Period:                       renewReq.Period,
		Amount:                       1,
		InternetChargeType:           req.InternetChargeType,
		ImageID:                      req.ImageId,
		InternetMaxBandwidthOut:      req.InternetMaxBandwidthOut,
		SystemDiskCategory:           req.SystemDiskCategory,
		SystemDiskSize:               req.SystemDiskSize,
		DescribePriceRequestDataDisk: req.DataDisk,
	}
	priceInfo, err := m.DescribePrice(ctx, priceReq)
	if err != nil {
		log.Errorf("DescribePrice:%v", err)
		return "", &api.ErrWeb{Code: terrors.DescribePriceError.Int(), Message: err.Error()}
	}

	hash := uuid.NewString()
	orderID := strings.Replace(hash, "-", "", -1)
	req.OrderID = orderID
	req.TradePrice = priceInfo.USDPrice
	req.PeriodUnit = renewReq.PeriodUnit
	req.Period = renewReq.Period
	req.Renew = renewReq.Renew
	err = m.RenewVpsInstance(req)
	if err != nil {
		log.Errorf("SaveVpsInstance:%v", err)
		return "", err
	}
	newBalanceString := strconv.FormatFloat(math.Ceil(float64(priceInfo.USDPrice)*1000000), 'f', 0, 64)

	info := &types.OrderRecord{
		VpsID:     req.Id,
		OrderID:   orderID,
		UserID:    userID,
		Value:     newBalanceString,
		OrderType: types.RenewVPS,
	}

	err = m.OrderMgr.CreatedOrder(info)
	if err != nil {
		return "", err
	}

	return orderID, nil
}

func (m *Mall) GetUseWaitingPaymentOrders(ctx context.Context, limit, page int64) (*types.OrderRecordResponse, error) {
	userID := handler.GetID(ctx)

	info, err := m.LoadOrderRecordByUserUndone(userID, limit, page, m.OrderMgr.GetOrderTimeoutMinute())
	if err != nil {
		return nil, &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}

	return info, nil
}

func (m *Mall) GetUserOrderRecords(ctx context.Context, limit, page int64) (*types.OrderRecordResponse, error) {
	userID := handler.GetID(ctx)

	info, err := m.LoadOrderRecordsByUser(userID, limit, page, m.OrderMgr.GetOrderTimeoutMinute())
	if err != nil {
		return nil, &api.ErrWeb{Code: terrors.DatabaseError.Int(), Message: err.Error()}
	}

	return info, nil
}

func (m *Mall) CancelUserOrder(ctx context.Context, orderID string) error {
	userID := handler.GetID(ctx)
	return m.OrderMgr.CancelOrder(orderID, userID)
}

func (m *Mall) PaymentUserOrder(ctx context.Context, orderID string) error {
	userID := handler.GetID(ctx)
	return m.OrderMgr.PaymentCompleted(orderID, userID)
}
