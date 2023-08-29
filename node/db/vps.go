package db

import (
	"fmt"
	"time"

	"github.com/LMF709268224/titan-vps/api/types"
)

// LoadVpsInfo loads VPS information by VPS ID.
func (d *SQLDB) LoadVpsInfo(vpsID int64) (*types.CreateInstanceReq, error) {
	var info types.CreateInstanceReq
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=?", userInstancesTable)
	err := d.db.Get(&info, query, vpsID)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// LoadVpsInfoByInstanceID loads VPS information by instance ID.
func (d *SQLDB) LoadVpsInfoByInstanceID(instanceID string) (*types.CreateInstanceReq, error) {
	var info types.CreateInstanceReq
	query := fmt.Sprintf("SELECT * FROM %s WHERE instance_id=?", userInstancesTable)
	err := d.db.Get(&info, query, instanceID)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// VpsExists checks if VPS info exists in the database.
func (d *SQLDB) VpsExists(vpsID int64) (bool, error) {
	var total int64
	countSQL := fmt.Sprintf(`SELECT count(id) FROM %s WHERE id=? `, userInstancesTable)
	if err := d.db.Get(&total, countSQL, vpsID); err != nil {
		return false, err
	}

	return total > 0, nil
}

// VpsDeviceExists checks if VPS device info exists in the database.
func (d *SQLDB) VpsDeviceExists(instanceID int64) (bool, error) {
	var total int64
	countSQL := fmt.Sprintf(`SELECT count(instance_id) FROM %s WHERE instance_id=? `, userInstancesTable)
	if err := d.db.Get(&total, countSQL, instanceID); err != nil {
		return false, err
	}

	return total > 0, nil
}

// SaveVpsInstance saves VPS instance information into the database.
func (d *SQLDB) SaveVpsInstance(rInfo *types.CreateOrderReq) (int64, error) {
	query := fmt.Sprintf(
		`INSERT INTO %s (region_id,instance_id,user_id,order_id, instance_type, dry_run, image_id,
			    security_group_id, instance_charge_type,internet_charge_type, period_unit, period, bandwidth_out,bandwidth_in,
			    ip_address,trade_price,system_disk_category,system_disk_size,os_type,data_disk,renew, access_key) 
				VALUES (:region_id,:instance_id,:user_id,:order_id, :instance_type, :dry_run, :image_id, 
				:security_group_id, :instance_charge_type,:internet_charge_type, :period_unit, :period, :bandwidth_out,:bandwidth_in,
				:ip_address,:trade_price,:system_disk_category,:system_disk_size,:os_type,:data_disk,:renew, :access_key)`, userInstancesTable)

	result, err := d.db.NamedExec(query, rInfo)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// UpdateVpsInstance updates VPS instance information in the database.
func (d *SQLDB) UpdateVpsInstance(info *types.CreateInstanceReq) error {
	query := fmt.Sprintf(`UPDATE %s SET ip_address=?, instance_id=?, user_id=?,os_type=?,cores=?,memory=?,expired_time=?,
	    security_group_id=? WHERE order_id=?`, userInstancesTable)
	_, err := d.db.Exec(query, info.IpAddress, info.InstanceId, info.UserID, info.OSType, info.Cores, info.Memory, info.ExpiredTime, info.SecurityGroupId, info.OrderID)

	return err
}

// RenewVpsInstance updates VPS instance renewal information in the database.
func (d *SQLDB) RenewVpsInstance(info *types.CreateInstanceReq) error {
	query := fmt.Sprintf(`UPDATE %s SET period_unit=?, period=?, trade_price=?,renew=? WHERE instance_id=?`, userInstancesTable)
	_, err := d.db.Exec(query, info.PeriodUnit, info.Period, info.TradePrice, info.Renew, info.InstanceId)
	if err != nil {
		return err
	}
	return nil
}

// UpdateRenewInstanceStatus updates VPS instance renewal status in the database.
func (d *SQLDB) UpdateRenewInstanceStatus(info *types.SetRenewOrderReq) error {
	query := fmt.Sprintf(`UPDATE %s SET renew=? WHERE instance_id=?`, userInstancesTable)
	_, err := d.db.Exec(query, info.Renew, info.InstanceId)
	if err != nil {
		return err
	}
	return nil
}

// UpdateVpsInstanceName updates VPS instance name in the database.
func (d *SQLDB) UpdateVpsInstanceName(instanceID, instanceName, userID string) error {
	query := fmt.Sprintf(`UPDATE %s SET instance_name=? WHERE instance_id=? and user_id=?`, userInstancesTable)
	_, err := d.db.Exec(query, instanceName, instanceID, userID)
	if err != nil {
		return err
	}
	//query = fmt.Sprintf(`UPDATE %s SET instance_name=? WHERE instance_id=? and user_id=?`, myInstancesTable)
	//_, err = d.db.Exec(query, instanceName, instanceID, userID)

	return err
}

// SaveInstancesInfo saves order information.
func (d *SQLDB) SaveInstancesInfo(rInfo *types.DescribeInstanceTypeFromBase) error {
	query := fmt.Sprintf(
		`INSERT INTO %s (instance_type_id, region_id, memory_size,cpu_architecture,instance_category,cpu_core_count,available_zone,instance_type_family,physical_processor_model,price,original_price,status) 
		        VALUES (:instance_type_id, :region_id, :memory_size,:cpu_architecture,:instance_category,:cpu_core_count,:available_zone,:instance_type_family,:physical_processor_model,:price,:original_price,:status)
				ON DUPLICATE KEY UPDATE price=:price,status=:status,original_price=:original_price,updated_time=NOW()`, instanceBaseInfoTable)

	_, err := d.db.NamedExec(query, rInfo)

	return err
}

// InstancesDefaultExists checks if instance defaults exist for a specific instance type and region.
func (d *SQLDB) InstancesDefaultExists(instanceTypeID, regionID string) (bool, error) {
	var total int64
	timeString := time.Now().Format("2006-01-02")
	countSQL := fmt.Sprintf(`SELECT count(1) FROM %s WHERE instance_type_id=? and region_id=? and updated_time>?`, instanceBaseInfoTable)
	if err := d.db.Get(&total, countSQL, instanceTypeID, regionID, timeString); err != nil {
		return false, err
	}

	return total > 0, nil
}

// UpdateInstanceDefaultStatus updates the status of instance defaults for a specific instance type and region.
func (d *SQLDB) UpdateInstanceDefaultStatus(instanceTypeID, regionID string) error {
	query := fmt.Sprintf(`UPDATE %s SET status='' and updated_time=NOW() WHERE instance_type_id=? and region_id=?`, instanceBaseInfoTable)
	_, err := d.db.Exec(query, instanceTypeID, regionID)
	if err != nil {
		return err
	}

	return err
}

// LoadInstancesInfoOfUser loads user server information.
func (d *SQLDB) LoadInstancesInfoOfUser(userID string, limit, offset int64) (*types.MyInstanceResponse, error) {
	out := new(types.MyInstanceResponse)
	var infos []*types.MyInstance
	query := fmt.Sprintf("SELECT id,instance_id,instance_name,os_type,region_id,trade_price,created_time,bandwidth_out,expired_time FROM %s WHERE user_id=?  order by created_time desc LIMIT ? OFFSET ?", userInstancesTable)
	if limit > loadOrderRecordsDefaultLimit {
		limit = loadOrderRecordsDefaultLimit
	}
	err := d.db.Select(&infos, query, userID, limit, offset*limit)
	if err != nil {
		return nil, err
	}

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE user_id=?", userInstancesTable)
	var count int
	err = d.db.Get(&count, countQuery, userID)
	if err != nil {
		return nil, err
	}

	out.Total = count
	out.List = infos

	return out, nil
}

// LoadInstanceDetailsInfo loads details of a specific instance.
func (d *SQLDB) LoadInstanceDetailsInfo(userID, instanceID string) (*types.InstanceDetails, error) {
	var info types.InstanceDetails
	query := fmt.Sprintf("SELECT region_id,instance_id,instance_name,instance_type,image_id,security_group_id,instance_charge_type,internet_charge_type,bandwidth_out,bandwidth_in,system_disk_size,ip_address,system_disk_category,created_time,memory,memory_used,cores,cores_used,os_type,data_disk  FROM %s WHERE user_id=? and instance_id=?", userInstancesTable)
	err := d.db.Get(&info, query, userID, instanceID)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// LoadInstanceDefaultInfo loads instance type defaults based on specified criteria.
func (d *SQLDB) LoadInstanceDefaultInfo(req *types.InstanceTypeFromBaseReq) (*types.InstanceTypeResponse, error) {
	out := new(types.InstanceTypeResponse)
	var info []*types.DescribeInstanceTypeFromBase
	var query string
	var args []interface{}

	query = "region_id=?"
	args = append(args, req.RegionId)
	if req.InstanceCategory != "" {
		query += " and instance_category=?"
		args = append(args, req.InstanceCategory)
	}
	if req.MemorySize != 0 {
		query += " and memory_size=?"
		args = append(args, req.MemorySize)
	}
	if req.CpuCoreCount != 0 {
		query += " and cpu_core_count=?"
		args = append(args, req.CpuCoreCount)
	}
	if req.CpuArchitecture != "" {
		query += " and cpu_architecture=?"
		args = append(args, req.CpuArchitecture)
	}

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s and status!=''", instanceBaseInfoTable, query)
	var count int
	err := d.db.Get(&count, countQuery, args...)
	if err != nil {
		return nil, err
	}

	querySQL := fmt.Sprintf("SELECT * FROM %s WHERE %s and status!='' LIMIT %d OFFSET %d ", instanceBaseInfoTable, query, req.Limit, req.Offset)
	err = d.db.Select(&info, querySQL, args...)
	if err != nil {
		return nil, err
	}

	out.Total = count
	out.List = info
	return out, nil
}

// LoadInstanceCPUInfo loads distinct CPU core counts based on specified criteria.
func (d *SQLDB) LoadInstanceCPUInfo(req *types.InstanceTypeFromBaseReq) ([]*int32, error) {
	var info []*int32
	var query string
	var args []interface{}
	query = "region_id=?"
	args = append(args, req.RegionId)
	if req.InstanceCategory != "" {
		query += " and instance_category=?"
		args = append(args, req.InstanceCategory)
	}
	if req.CpuCoreCount != 0 {
		query += " and cpu_core_count=?"
		args = append(args, req.CpuCoreCount)
	}
	if req.CpuArchitecture != "" {
		query += " and cpu_architecture=?"
		args = append(args, req.CpuArchitecture)
	}

	querySQL := fmt.Sprintf("SELECT distinct cpu_core_count FROM %s WHERE %s order by cpu_core_count asc", instanceBaseInfoTable, query)
	err := d.db.Select(&info, querySQL, args...)
	if err != nil {
		return nil, err
	}
	return info, nil
}

// LoadInstanceMemoryInfo loads distinct memory sizes based on specified criteria.
func (d *SQLDB) LoadInstanceMemoryInfo(req *types.InstanceTypeFromBaseReq) ([]*float32, error) {
	var info []*float32
	var query string
	var args []interface{}
	query = "region_id=?"
	args = append(args, req.RegionId)
	if req.InstanceCategory != "" {
		query += " and instance_category=?"
		args = append(args, req.InstanceCategory)
	}
	if req.CpuCoreCount != 0 {
		query += " and cpu_core_count=?"
		args = append(args, req.CpuCoreCount)
	}
	if req.CpuArchitecture != "" {
		query += " and cpu_architecture=?"
		args = append(args, req.CpuArchitecture)
	}

	querySQL := fmt.Sprintf("SELECT distinct memory_size FROM %s WHERE %s order by memory_size asc", instanceBaseInfoTable, query)
	err := d.db.Select(&info, querySQL, args...)
	if err != nil {
		return nil, err
	}
	return info, nil
}
