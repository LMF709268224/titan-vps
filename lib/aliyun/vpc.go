package aliyun

import (
	"sync"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	vpc20160428 "github.com/alibabacloud-go/vpc-20160428/v5/client"
)

var VpsClient sync.Map

// ClientInit /**
func newVpsClient(regionID, keyID, keySecret string) (*vpc20160428.Client, *tea.SDKError) {
	if v, ok := VpsClient.Load(regionID); ok {
		c := v.(*vpc20160428.Client)
		return c, nil
	}
	configClient := &openapi.Config{
		AccessKeyId:     tea.String(keyID),
		AccessKeySecret: tea.String(keySecret),
	}

	configClient.RegionId = tea.String(regionID)
	client, err := vpc20160428.NewClient(configClient)
	if err != nil {
		errors := &tea.SDKError{}
		if _t, ok := err.(*tea.SDKError); ok {
			errors = _t
		} else {
			errors.Message = tea.String(err.Error())
		}
		return nil, errors
	}

	VpsClient.Store(regionID, client)
	return client, nil
}

// CreateVpc Create Vpc
func CreateVpc(keyID, keySecret, regionID string) (string, *tea.SDKError) {
	var out string

	client, err := newVpsClient(regionID, keyID, keySecret)
	if err != nil {
		return out, err
	}
	request := &vpc20160428.CreateVpcRequest{
		RegionId: tea.String(regionID),
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()

		result, err := client.CreateVpcWithOptions(request, getRunTime())
		if err != nil {
			return err
		}

		out = *result.Body.VpcId

		return nil
	}()

	if tryErr != nil {
		errors := &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			errors = _t
		} else {
			errors.Message = tea.String(tryErr.Error())
		}
		return out, errors
	}
	return out, nil
}

// DescribeVpcs Describe Vpcs
func DescribeVpcs(keyID, keySecret, regionID string) (string, *tea.SDKError) {
	var out string

	client, err := newVpsClient(regionID, keyID, keySecret)
	if err != nil {
		return out, err
	}
	request := &vpc20160428.DescribeVpcsRequest{
		RegionId: tea.String(regionID),
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()

		result, err := client.DescribeVpcsWithOptions(request, getRunTime())
		if err != nil {
			return err
		}
		list := result.Body.Vpcs.Vpc
		if len(list) > 0 {
			out = *list[0].VpcId
		}
		return nil
	}()

	if tryErr != nil {
		errors := &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			errors = _t
		} else {
			errors.Message = tea.String(tryErr.Error())
		}
		return out, errors
	}
	return out, nil
}

// DeleteVpc delete Vpcs
func DeleteVpc(keyID, keySecret, regionID, vpcID string) *tea.SDKError {
	client, err := newVpsClient(regionID, keyID, keySecret)
	if err != nil {
		return err
	}
	request := &vpc20160428.DeleteVpcRequest{
		RegionId: tea.String(regionID),
		VpcId:    tea.String(vpcID),
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()

		_, err := client.DeleteVpcWithOptions(request, getRunTime())
		if err != nil {
			return err
		}
		return nil
	}()

	if tryErr != nil {
		errors := &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			errors = _t
		} else {
			errors.Message = tea.String(tryErr.Error())
		}
		return errors
	}

	return nil
}
