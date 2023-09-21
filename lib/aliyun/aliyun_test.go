package aliyun

import (
	"fmt"
	"testing"
	"time"

	"github.com/LMF709268224/titan-vps/api/types"
)

var (
	// AliyunAccessKeyID     = "LTAI5tLczyBS9h4APo9u9c1X"
	// AliyunAccessKeySecret = "xHF9hl7qesCYqv8Wp1TLjhDU6EPPjQ"
	AliyunAccessKeyID     = "LTAI5tQht2pcTNAaZicFjEQm"
	AliyunAccessKeySecret = "L7ZetXCPzlIWmVwljazD0yE4xnUeR4"
)

func TestDelete(t *testing.T) {
	// err := DeleteVpc(AliyunAccessKeyID, AliyunAccessKeySecret, "cn-qingdao", "vpc-m5e24v3cd7ercgta3n2dt")
	// fmt.Println("delete vpc err:", err)

	// err := DeleteLaunchTemplate(AliyunAccessKeyID, AliyunAccessKeySecret, "cn-qingdao", "lt-m5e0vtm7m9x70ig9qgb8")
	// fmt.Println("delete err:", err)
}

func TestDescribeInstanceBill(t *testing.T) {
	DescribeInstanceBill(AliyunAccessKeyID, AliyunAccessKeySecret)
}

func TestF(t *testing.T) {
	_, err := RefundInstance(AliyunAccessKeyID, AliyunAccessKeySecret, "i-m5e3qxkg5kep3eor3kji")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}

func TestDescribeInstances(t *testing.T) {
	regionID := "cn-zhangjiakou"
	list := []string{"i-8vbb4g761gonesej1bcg"}

	result, err := DescribeInstances(regionID, AliyunAccessKeyID, AliyunAccessKeySecret, list)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println("result:", result)
}

func TestInquiryPriceRefundInstance(t *testing.T) {
	_, err := InquiryPriceRefundInstance(AliyunAccessKeyID, AliyunAccessKeySecret, "i-m5e9d2l4yilyum2nrvs5")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}

func TestQueryProductList(t *testing.T) {
	// QueryProductList(AliyunAccessKeyID, AliyunAccessKeySecret)
	result, err := DescribePrice(AliyunAccessKeyID, AliyunAccessKeySecret, &types.DescribePriceReq{RegionId: "cn-qingdao"})
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println(result)
}

func TestTime(t *testing.T) {
	txx, err := time.Parse("2006-01-02T15:04Z", "2023-09-06T16:00Z")
	fmt.Println(err)
	fmt.Println(txx.String())
}
