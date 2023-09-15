package aliyun

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/LMF709268224/titan-vps/api/types"
	"github.com/opentracing/opentracing-go/log"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

const (
	defaultRegionID = "cn-hangzhou"
)

var EcsClient sync.Map

func getRunTime() *util.RuntimeOptions {
	return &util.RuntimeOptions{ConnectTimeout: tea.Int(int(5 * time.Second))}
}

// ClientInit /**
func newEcsClient(regionID, keyID, keySecret string) (*ecs20140526.Client, *tea.SDKError) {
	if v, ok := EcsClient.Load(regionID); ok {
		c := v.(*ecs20140526.Client)
		return c, nil
	}
	configClient := &openapi.Config{
		AccessKeyId:     tea.String(keyID),
		AccessKeySecret: tea.String(keySecret),
	}

	configClient.RegionId = tea.String(regionID)
	client, err := ecs20140526.NewClient(configClient)
	if err != nil {
		errors := &tea.SDKError{}
		if _t, ok := err.(*tea.SDKError); ok {
			errors = _t
		} else {
			errors.Message = tea.String(err.Error())
		}
		return nil, errors
	}

	EcsClient.Store(regionID, client)
	return client, nil
}

// CreateInstance crate an instance
func CreateInstance(keyID, keySecret string, instanceReq *types.CreateInstanceReq, dryRun bool) (*types.CreateInstanceResponse, *tea.SDKError) {
	var out *types.CreateInstanceResponse

	client, err := newEcsClient(instanceReq.RegionId, keyID, keySecret)
	if err != nil {
		return out, err
	}
	fmt.Println("instanceReq.SecurityGroupID : ", instanceReq.SecurityGroupID)
	createInstanceRequest := &ecs20140526.CreateInstanceRequest{
		RegionId:     tea.String(instanceReq.RegionId),
		InstanceType: tea.String(instanceReq.InstanceType),
		DryRun:       tea.Bool(dryRun),
		ImageId:      tea.String(instanceReq.ImageID),
		// SecurityGroupId:    tea.String(instanceReq.SecurityGroupID),
		InstanceChargeType: tea.String("PrePaid"),
		PeriodUnit:         tea.String(instanceReq.PeriodUnit),
		InternetChargeType: tea.String(instanceReq.InternetChargeType),
		Period:             tea.Int32(instanceReq.Period),
		// Password:                tea.String(password),
		InternetMaxBandwidthOut: tea.Int32(1),
		InternetMaxBandwidthIn:  tea.Int32(1),
		SystemDisk: &ecs20140526.CreateInstanceRequestSystemDisk{
			Size:     tea.Int32(instanceReq.SystemDiskSize),
			Category: tea.String(instanceReq.SystemDiskCategory),
		},
		DataDisk: []*ecs20140526.CreateInstanceRequestDataDisk{},
	}
	if len(instanceReq.DataDisk) > 0 {
		for _, v := range instanceReq.DataDisk {
			size := v.Size
			size32 := int32(size)
			DataDiskInfo := &ecs20140526.CreateInstanceRequestDataDisk{
				Category: tea.String(v.Category),
				Size:     tea.Int32(size32),
			}
			createInstanceRequest.DataDisk = append(createInstanceRequest.DataDisk, DataDiskInfo)
		}
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()

		result, err := client.CreateInstanceWithOptions(createInstanceRequest, getRunTime())
		if err != nil {
			return err
		}

		out = &types.CreateInstanceResponse{
			InstanceID: *result.Body.InstanceId,
			OrderId:    *result.Body.OrderId,
			RequestId:  *result.Body.RequestId,
			TradePrice: *result.Body.TradePrice,
			AccessKey:  keyID,
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

// RunInstances run aliyun instances
func RunInstances(keyID, keySecret, launchTemplateID string, instanceReq *types.CreateInstanceReq, dryRun bool) (*ecs20140526.RunInstancesResponse, *tea.SDKError) {
	var out *ecs20140526.RunInstancesResponse

	client, err := newEcsClient(instanceReq.RegionId, keyID, keySecret)
	if err != nil {
		return out, err
	}

	createInstanceRequest := &ecs20140526.RunInstancesRequest{
		RegionId:           tea.String(instanceReq.RegionId),
		InstanceType:       tea.String(instanceReq.InstanceType),
		DryRun:             tea.Bool(dryRun),
		ImageId:            tea.String(instanceReq.ImageID),
		SecurityGroupId:    tea.String(instanceReq.SecurityGroupID),
		InstanceChargeType: tea.String("PrePaid"),
		PeriodUnit:         tea.String(instanceReq.PeriodUnit),
		InternetChargeType: tea.String(instanceReq.InternetChargeType),
		Period:             tea.Int32(instanceReq.Period),
		LaunchTemplateId:   tea.String(launchTemplateID),
		// Password:                tea.String(password),
		InternetMaxBandwidthOut: tea.Int32(10),
		InternetMaxBandwidthIn:  tea.Int32(10),
		SystemDisk: &ecs20140526.RunInstancesRequestSystemDisk{
			Size:     tea.String(strconv.FormatInt(int64(instanceReq.SystemDiskSize), 10)),
			Category: tea.String(instanceReq.SystemDiskCategory),
		},
		DataDisk: []*ecs20140526.RunInstancesRequestDataDisk{},
	}
	if len(instanceReq.DataDisk) > 0 {
		for _, v := range instanceReq.DataDisk {
			size := v.Size
			size32 := int32(size)
			DataDiskInfo := &ecs20140526.RunInstancesRequestDataDisk{
				Category: tea.String(v.Category),
				Size:     tea.Int32(size32),
			}
			createInstanceRequest.DataDisk = append(createInstanceRequest.DataDisk, DataDiskInfo)
		}
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()

		result, err := client.RunInstancesWithOptions(createInstanceRequest, getRunTime())
		if err != nil {
			return err
		}

		out = result

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

// StartInstance start an instance
func StartInstance(regionID, keyID, keySecret, instanceID string) *tea.SDKError {
	client, err := newEcsClient(regionID, keyID, keySecret)
	if err != nil {
		return err
	}

	startInstancesRequest := &ecs20140526.StartInstancesRequest{
		RegionId:   tea.String(regionID),
		InstanceId: tea.StringSlice([]string{instanceID}),
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()

		_, err := client.StartInstancesWithOptions(startInstancesRequest, getRunTime())
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

// DescribeSecurityGroups describe user security groups
func DescribeSecurityGroups(regionID, keyID, keySecret string) ([]string, *tea.SDKError) {
	var out []string

	client, err := newEcsClient(regionID, keyID, keySecret)
	if err != nil {
		return out, err
	}

	describeSecurityGroupsRequest := &ecs20140526.DescribeSecurityGroupsRequest{
		RegionId: tea.String(regionID),
		// NetworkType: tea.String("classic"),
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()

		response, err := client.DescribeSecurityGroupsWithOptions(describeSecurityGroupsRequest, getRunTime())
		if err != nil {
			return err
		}

		grop := response.Body.SecurityGroups.SecurityGroup
		for _, g := range grop {
			out = append(out, *g.SecurityGroupId)
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

// DescribeInstanceAttribute describe attribute of instance
func DescribeInstanceAttribute(regionID, keyID, keySecret, instanceID string) (*ecs20140526.DescribeInstanceAttributeResponse, *tea.SDKError) {
	var out *ecs20140526.DescribeInstanceAttributeResponse

	client, err := newEcsClient(regionID, keyID, keySecret)
	if err != nil {
		return out, err
	}

	describeInstanceAttributeRequest := &ecs20140526.DescribeInstanceAttributeRequest{
		InstanceId: tea.String(instanceID),
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()

		result, err := client.DescribeInstanceAttributeWithOptions(describeInstanceAttributeRequest, getRunTime())
		if err != nil {
			return err
		}

		out = result

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

// AllocatePublicIPAddress Allocate IP address
func AllocatePublicIPAddress(regionID, keyID, keySecret, instanceID string) (string, *tea.SDKError) {
	var out string

	client, err := newEcsClient(regionID, keyID, keySecret)
	if err != nil {
		return out, err
	}

	allocatePublicIPAddressRequest := &ecs20140526.AllocatePublicIpAddressRequest{
		InstanceId: tea.String(instanceID),
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()

		result, err := client.AllocatePublicIpAddressWithOptions(allocatePublicIPAddressRequest, getRunTime())
		if err != nil {
			return err
		}

		out = *result.Body.IpAddress

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

// DescribePrice describe instance price
func DescribePrice(keyID, keySecret string, priceReq *types.DescribePriceReq) (*types.DescribePriceResponse, *tea.SDKError) {
	var out *types.DescribePriceResponse

	client, err := newEcsClient(priceReq.RegionId, keyID, keySecret)
	if err != nil {
		return out, err
	}
	describePriceRequest := &ecs20140526.DescribePriceRequest{
		RegionId:                tea.String(priceReq.RegionId),
		InstanceType:            tea.String(priceReq.InstanceType),
		ResourceType:            tea.String("instance"),
		PriceUnit:               tea.String(priceReq.PriceUnit),
		Period:                  tea.Int32(priceReq.Period),
		ImageId:                 tea.String(priceReq.ImageID),
		InternetChargeType:      tea.String(priceReq.InternetChargeType),
		Amount:                  tea.Int32(priceReq.Amount),
		InternetMaxBandwidthOut: tea.Int32(priceReq.InternetMaxBandwidthOut),
		// PayByBandwidth
		SystemDisk: &ecs20140526.DescribePriceRequestSystemDisk{
			Category: tea.String(priceReq.SystemDiskCategory),
			Size:     tea.Int32(priceReq.SystemDiskSize),
		},
		DataDisk: []*ecs20140526.DescribePriceRequestDataDisk{},
	}
	if len(priceReq.DescribePriceRequestDataDisk) > 0 {
		for _, v := range priceReq.DescribePriceRequestDataDisk {
			DataDiskInfo := &ecs20140526.DescribePriceRequestDataDisk{
				Category: tea.String(v.Category),
				// PerformanceLevel: tea.String("PL0"),
				Size: tea.Int64(v.Size),
			}
			describePriceRequest.DataDisk = append(describePriceRequest.DataDisk, DataDiskInfo)
		}
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		result, err := client.DescribePriceWithOptions(describePriceRequest, getRunTime())
		if err != nil {
			return err
		}
		price := result.Body.PriceInfo.Price
		out = &types.DescribePriceResponse{
			Currency:      *price.Currency,
			OriginalPrice: *price.OriginalPrice,
			TradePrice:    *price.TradePrice,
			USDPrice:      *price.TradePrice,
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

// AuthorizeSecurityGroup authorize security group
func AuthorizeSecurityGroup(regionID, keyID, keySecret, securityGroupID string) *tea.SDKError {
	client, err := newEcsClient(regionID, keyID, keySecret)
	if err != nil {
		return err
	}

	authorizeSecurityGroupRequest := &ecs20140526.AuthorizeSecurityGroupRequest{
		RegionId:        tea.String(regionID),
		SecurityGroupId: tea.String(securityGroupID),
		Permissions: []*ecs20140526.AuthorizeSecurityGroupRequestPermissions{
			{
				// TODO
				IpProtocol:   tea.String("ALL"),
				SourceCidrIp: tea.String("0.0.0.0/0"),
				PortRange:    tea.String("-1/-1"),
			},
		},
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		_, err := client.AuthorizeSecurityGroupWithOptions(authorizeSecurityGroupRequest, getRunTime())
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

// DescribeRegions describe regions
func DescribeRegions(keyID, keySecret string) (*ecs20140526.DescribeRegionsResponse, *tea.SDKError) {
	client, err := newEcsClient(defaultRegionID, keyID, keySecret)
	if err != nil {
		return nil, err
	}

	var result *ecs20140526.DescribeRegionsResponse
	describeRegionsRequest := &ecs20140526.DescribeRegionsRequest{}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()

		result, _e = client.DescribeRegionsWithOptions(describeRegionsRequest, getRunTime())
		if _e != nil {
			return _e
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
		return result, errors
	}
	return result, nil
}

// DescribeRecommendInstanceType Describe Instance Type
func DescribeRecommendInstanceType(keyID, keySecret string, instanceTypeReq *types.DescribeRecommendInstanceTypeReq) (*ecs20140526.DescribeRecommendInstanceTypeResponse, *tea.SDKError) {
	var result *ecs20140526.DescribeRecommendInstanceTypeResponse
	client, err := newEcsClient(instanceTypeReq.RegionId, keyID, keySecret)
	if err != nil {
		return result, err
	}

	describeRecommendInstanceTypeRequest := &ecs20140526.DescribeRecommendInstanceTypeRequest{
		NetworkType:        tea.String("vpc"),
		RegionId:           tea.String(instanceTypeReq.RegionId),
		InstanceChargeType: tea.String(instanceTypeReq.InstanceChargeType),
	}
	if instanceTypeReq.Cores > 0 {
		describeRecommendInstanceTypeRequest.Cores = tea.Int32(instanceTypeReq.Cores)
	}
	if instanceTypeReq.Memory > 0 {
		describeRecommendInstanceTypeRequest.Memory = tea.Float32(instanceTypeReq.Memory)
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		result, _e = client.DescribeRecommendInstanceTypeWithOptions(describeRecommendInstanceTypeRequest, getRunTime())
		if _e != nil {
			return _e
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
		return result, errors
	}
	return result, nil
}

func DescribeInstanceTypes(keyID, keySecret string, instanceType *types.DescribeInstanceTypeReq) (*ecs20140526.DescribeInstanceTypesResponse, *tea.SDKError) {
	var result *ecs20140526.DescribeInstanceTypesResponse
	client, err := newEcsClient(instanceType.RegionId, keyID, keySecret)
	if err != nil {
		return result, err
	}
	describeInstanceTypesRequest := &ecs20140526.DescribeInstanceTypesRequest{}
	if instanceType.CpuArchitecture != "" {
		describeInstanceTypesRequest.CpuArchitecture = tea.String(instanceType.CpuArchitecture)
	}
	if instanceType.InstanceCategory != "" {
		describeInstanceTypesRequest.InstanceCategory = tea.String(instanceType.InstanceCategory)
	}
	if instanceType.CpuCoreCount != 0 {
		describeInstanceTypesRequest.MinimumCpuCoreCount = tea.Int32(instanceType.CpuCoreCount)
		describeInstanceTypesRequest.MaximumCpuCoreCount = tea.Int32(instanceType.CpuCoreCount)
	}
	if instanceType.MemorySize != 0 {
		describeInstanceTypesRequest.MinimumMemorySize = tea.Float32(instanceType.MemorySize)
		describeInstanceTypesRequest.MaximumMemorySize = tea.Float32(instanceType.MemorySize)
	}
	if instanceType.MaxResults != 0 {
		describeInstanceTypesRequest.MaxResults = tea.Int64(instanceType.MaxResults)
	}
	if instanceType.NextToken != "" {
		describeInstanceTypesRequest.NextToken = tea.String(instanceType.NextToken)
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		result, _e = client.DescribeInstanceTypesWithOptions(describeInstanceTypesRequest, getRunTime())
		if _e != nil {
			return _e
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
		return result, errors
	}
	return result, nil
}

// CreateSecurityGroup Create Security Group
func CreateSecurityGroup(regionID, keyID, keySecret string) (string, *tea.SDKError) {
	var out string

	client, err := newEcsClient(regionID, keyID, keySecret)
	if err != nil {
		return out, err
	}

	createSecurityGroupRequest := &ecs20140526.CreateSecurityGroupRequest{
		RegionId: tea.String(regionID),
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		result, err := client.CreateSecurityGroupWithOptions(createSecurityGroupRequest, getRunTime())
		if err != nil {
			return err
		}

		out = *result.Body.SecurityGroupId
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

// DescribeImages Describe Images
func DescribeImages(regionID, keyID, keySecret, instanceType string) (*ecs20140526.DescribeImagesResponse, *tea.SDKError) {
	var result *ecs20140526.DescribeImagesResponse

	client, err := newEcsClient(regionID, keyID, keySecret)
	if err != nil {
		return result, err
	}

	createSecurityGroupRequest := &ecs20140526.DescribeImagesRequest{
		RegionId: tea.String(regionID),
	}
	if instanceType != "" {
		createSecurityGroupRequest.InstanceType = tea.String(instanceType)
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		result, _e = client.DescribeImagesWithOptions(createSecurityGroupRequest, getRunTime())
		if _e != nil {
			return _e
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
		return result, errors
	}
	return result, nil
}

// DescribeInstanceStatus query instance status
func DescribeInstanceStatus(regionID, keyID, keySecret string, InstanceId []string) (*ecs20140526.DescribeInstanceStatusResponse, *tea.SDKError) {
	var result *ecs20140526.DescribeInstanceStatusResponse

	client, err := newEcsClient(regionID, keyID, keySecret)
	if err != nil {
		return result, err
	}

	createSecurityGroupRequest := &ecs20140526.DescribeInstanceStatusRequest{
		RegionId:   tea.String(regionID),
		InstanceId: tea.StringSlice(InstanceId),
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		result, _e = client.DescribeInstanceStatusWithOptions(createSecurityGroupRequest, getRunTime())
		if _e != nil {
			return _e
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
		return result, errors
	}
	return result, nil
}

// DescribeInstances instance detail info
func DescribeInstances(regionID, keyID, keySecret string, InstanceIds []string) (*ecs20140526.DescribeInstancesResponse, *tea.SDKError) {
	var result *ecs20140526.DescribeInstancesResponse

	client, err := newEcsClient(regionID, keyID, keySecret)
	if err != nil {
		return result, err
	}
	instanceIdsByte, e := json.Marshal(InstanceIds)
	if e != nil {
		log.Error(e)
	}
	instanceIdSting := string(instanceIdsByte)
	createSecurityGroupRequest := &ecs20140526.DescribeInstancesRequest{
		RegionId:    tea.String(regionID),
		InstanceIds: tea.String(instanceIdSting),
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		result, _e = client.DescribeInstancesWithOptions(createSecurityGroupRequest, getRunTime())
		if _e != nil {
			return _e
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
		return result, errors
	}
	return result, nil
}

// DescribeAvailableResource Describe Resource
func DescribeAvailableResource(keyID, keySecret string, instanceType *types.DescribeInstanceTypeReq) (*ecs20140526.DescribeAvailableResourceResponse, *tea.SDKError) {
	var result *ecs20140526.DescribeAvailableResourceResponse

	client, err := newEcsClient(instanceType.RegionId, keyID, keySecret)
	if err != nil {
		return result, err
	}

	describeAvailableResourceRequest := &ecs20140526.DescribeAvailableResourceRequest{
		NetworkCategory:     tea.String("vpc"),
		ResourceType:        tea.String("instance"),
		RegionId:            tea.String(instanceType.RegionId),
		DestinationResource: tea.String("InstanceType"),
		InstanceChargeType:  tea.String("PrePaid"),
		Cores:               tea.Int32(instanceType.CpuCoreCount),
		Memory:              tea.Float32(instanceType.MemorySize),
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		result, _e = client.DescribeAvailableResourceWithOptions(describeAvailableResourceRequest, getRunTime())
		if _e != nil {
			return _e
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
		return result, errors
	}
	return result, nil
}

func DescribeAvailableResourceForDesk(keyID, keySecret string, desk *types.AvailableResourceReq) (*ecs20140526.DescribeAvailableResourceResponse, *tea.SDKError) {
	var result *ecs20140526.DescribeAvailableResourceResponse

	client, err := newEcsClient(desk.RegionId, keyID, keySecret)
	if err != nil {
		return result, err
	}

	describeAvailableResourceRequest := &ecs20140526.DescribeAvailableResourceRequest{
		RegionId:            tea.String(desk.RegionId),
		DestinationResource: tea.String(desk.DestinationResource),
		InstanceChargeType:  tea.String("PrePaid"),
		InstanceType:        tea.String(desk.InstanceType),
		ResourceType:        tea.String("instance"),
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		result, _e = client.DescribeAvailableResourceWithOptions(describeAvailableResourceRequest, getRunTime())
		if _e != nil {
			return _e
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
		return result, errors
	}
	return result, nil
}

// CreateKeyPair Create key pair
func CreateKeyPair(regionID, keyID, keySecret, KeyPairName string) (*types.CreateKeyPairResponse, *tea.SDKError) {
	var out *types.CreateKeyPairResponse

	client, err := newEcsClient(regionID, keyID, keySecret)
	if err != nil {
		return out, err
	}

	createKeyPairRequest := &ecs20140526.CreateKeyPairRequest{
		RegionId:    tea.String(regionID),
		KeyPairName: tea.String(KeyPairName),
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		result, _e := client.CreateKeyPairWithOptions(createKeyPairRequest, getRunTime())
		if _e != nil {
			return _e
		}
		keyInfo := result.Body
		out = &types.CreateKeyPairResponse{
			KeyPairID:      *keyInfo.KeyPairId,
			KeyPairName:    *keyInfo.KeyPairName,
			PrivateKeyBody: *keyInfo.PrivateKeyBody,
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

// AttachKeyPair Attach KeyPair
func AttachKeyPair(regionID, keyID, keySecret, KeyPairName string, instanceIds []string) ([]*types.AttachKeyPairResponse, *tea.SDKError) {
	var out []*types.AttachKeyPairResponse

	client, err := newEcsClient(regionID, keyID, keySecret)
	if err != nil {
		return out, err
	}
	instanceIdsByte, e := json.Marshal(instanceIds)
	if e != nil {
		log.Error(e)
	}
	instanceIdSting := string(instanceIdsByte)
	attachKeyPairRequest := &ecs20140526.AttachKeyPairRequest{
		RegionId:    tea.String(regionID),
		KeyPairName: tea.String(KeyPairName),
		// InstanceIds should be []string
		InstanceIds: tea.String(instanceIdSting),
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		result, _e := client.AttachKeyPairWithOptions(attachKeyPairRequest, getRunTime())
		if _e != nil {
			return _e
		}
		for _, i := range result.Body.Results.Result {
			instanceInfo := &types.AttachKeyPairResponse{
				Code:       *i.Code,
				InstanceId: *i.InstanceId,
				Message:    *i.Message,
				Success:    *i.Success,
			}
			out = append(out, instanceInfo)
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

// RebootInstance  Reboot Instance
func RebootInstance(regionID, keyID, keySecret, instanceId string) *tea.SDKError {
	client, err := newEcsClient(regionID, keyID, keySecret)
	if err != nil {
		return err
	}

	rebootInstanceRequest := &ecs20140526.RebootInstanceRequest{
		InstanceId: tea.String(instanceId),
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		_, _e = client.RebootInstanceWithOptions(rebootInstanceRequest, getRunTime())
		if _e != nil {
			return _e
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

// RenewInstance renew instance
func RenewInstance(keyID, keySecret string, renewInstanceRequest *types.RenewInstanceRequest) *tea.SDKError {
	client, err := newEcsClient(renewInstanceRequest.RegionId, keyID, keySecret)
	if err != nil {
		return err
	}

	rebootInstanceRequest := &ecs20140526.RenewInstanceRequest{
		InstanceId: tea.String(renewInstanceRequest.InstanceId),
		Period:     tea.Int32(renewInstanceRequest.Period),
		PeriodUnit: tea.String(renewInstanceRequest.PeriodUnit),
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		_, _e = client.RenewInstanceWithOptions(rebootInstanceRequest, getRunTime())
		if _e != nil {
			return _e
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

func ModifyInstanceAutoRenewAttribute(keyID, keySecret string, renewInstanceRequest *types.SetRenewOrderReq) *tea.SDKError {
	client, err := newEcsClient(renewInstanceRequest.RegionID, keyID, keySecret)
	if err != nil {
		return err
	}
	rebootInstanceRequest := &ecs20140526.ModifyInstanceAutoRenewAttributeRequest{
		InstanceId: tea.String(renewInstanceRequest.InstanceId),
		RegionId:   tea.String(renewInstanceRequest.RegionID),
		PeriodUnit: tea.String(renewInstanceRequest.PeriodUnit),
		Duration:   tea.Int32(renewInstanceRequest.Period),
		AutoRenew:  tea.Bool(false),
	}
	if renewInstanceRequest.Renew == 1 {
		rebootInstanceRequest.AutoRenew = tea.Bool(true)
	}
	if renewInstanceRequest.PeriodUnit != "" {
		rebootInstanceRequest.PeriodUnit = tea.String(renewInstanceRequest.PeriodUnit)
	}
	if renewInstanceRequest.PeriodUnit != "" {
		rebootInstanceRequest.Duration = tea.Int32(renewInstanceRequest.Period)
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		_, _e = client.ModifyInstanceAutoRenewAttributeWithOptions(rebootInstanceRequest, getRunTime())
		if _e != nil {
			return _e
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

func DescribeInstanceAutoRenewAttribute(keyID, keySecret string, renewInstanceRequest *types.SetRenewOrderReq) (*ecs20140526.DescribeInstanceAutoRenewAttributeResponse, *tea.SDKError) {
	client, err := newEcsClient(renewInstanceRequest.RegionID, keyID, keySecret)
	if err != nil {
		return nil, err
	}
	var result *ecs20140526.DescribeInstanceAutoRenewAttributeResponse
	rebootInstanceRequest := &ecs20140526.DescribeInstanceAutoRenewAttributeRequest{
		InstanceId: tea.String(renewInstanceRequest.InstanceId),
		RegionId:   tea.String(renewInstanceRequest.RegionID),
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		result, _e = client.DescribeInstanceAutoRenewAttributeWithOptions(rebootInstanceRequest, getRunTime())
		if _e != nil {
			return _e
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
		return nil, errors
	}
	return result, nil
}

// CreateLaunchTemplate Create Launch Template
func CreateLaunchTemplate(keyID, keySecret, regionID, vpcID string) (string, *tea.SDKError) {
	var out string
	client, err := newEcsClient(regionID, keyID, keySecret)
	if err != nil {
		return out, err
	}

	request := &ecs20140526.CreateLaunchTemplateRequest{
		RegionId:           tea.String(regionID),
		VpcId:              tea.String(vpcID),
		LaunchTemplateName: tea.String(regionID),
		NetworkType:        tea.String("vpc"),
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()

		result, _e := client.CreateLaunchTemplateWithOptions(request, getRunTime())
		if _e != nil {
			return _e
		}

		out = *result.Body.LaunchTemplateId

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

// DescribeLaunchTemplates Describe Launch Templates
func DescribeLaunchTemplates(keyID, keySecret, regionID string) (string, *tea.SDKError) {
	var out string
	client, err := newEcsClient(regionID, keyID, keySecret)
	if err != nil {
		return out, err
	}

	request := &ecs20140526.DescribeLaunchTemplatesRequest{
		RegionId: tea.String(regionID),
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()

		result, _e := client.DescribeLaunchTemplatesWithOptions(request, getRunTime())
		if _e != nil {
			return _e
		}

		list := result.Body.LaunchTemplateSets.LaunchTemplateSet
		if len(list) > 0 {
			out = *list[0].LaunchTemplateId
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

// DeleteLaunchTemplate delete Launch Templates
func DeleteLaunchTemplate(keyID, keySecret, regionID, launchTemplateID string) *tea.SDKError {
	client, err := newEcsClient(regionID, keyID, keySecret)
	if err != nil {
		return err
	}

	request := &ecs20140526.DeleteLaunchTemplateRequest{
		RegionId:         tea.String(regionID),
		LaunchTemplateId: tea.String(launchTemplateID),
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()

		_, _e = client.DeleteLaunchTemplateWithOptions(request, getRunTime())
		if _e != nil {
			return _e
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
