package db

import (
	"fmt"

	"github.com/LMF709268224/titan-vps/api/types"
	"time"
)

// SaveMyInstancesInfo  save instance information
func (n *SQLDB) SaveMyInstancesInfo(rInfo *types.MyInstance) error {
	query := fmt.Sprintf(
		`INSERT INTO %s (instance_id, order_id, user_id, private_key_status, instance_name, instance_system, location,  price,state,internet_charge_type) 
		        VALUES (:instance_id, :order_id, :user_id, :private_key_status, :instance_name, :instance_system, :location, :price,:state,:internet_charge_type)`, myInstancesTable)
	_, err := n.db.NamedExec(query, rInfo)

	return err
}

// SaveInstancesInfo save order information
func (n *SQLDB) SaveInstancesInfo(rInfo *types.DescribeInstanceTypeFromBase) error {
	query := fmt.Sprintf(
		`INSERT INTO %s (instance_type_id, region_id, memory_size,cpu_architecture,instance_category,cpu_core_count,available_zone,instance_type_family,physical_processor_model,price,original_price,status) 
		        VALUES (:instance_type_id, :region_id, :memory_size,:cpu_architecture,:instance_category,:cpu_core_count,:available_zone,:instance_type_family,:physical_processor_model,:price,:original_price,:status)
				ON DUPLICATE KEY UPDATE price=:price,status=:status,original_price=:original_price,updated_time=NOW()`, instanceDefaultTable)

	_, err := n.db.NamedExec(query, rInfo)

	return err
}

func (n *SQLDB) InstancesDefaultExists(instanceTypeID, regionID string) (bool, error) {
	var total int64
	timeString := time.Now().Format(time.DateOnly)
	fmt.Println(timeString)
	countSQL := fmt.Sprintf(`SELECT count(1) FROM %s WHERE instance_type_id=? and region_id=? and updated_time>?`, instanceDefaultTable)
	if err := n.db.Get(&total, countSQL, instanceTypeID, regionID, timeString); err != nil {
		return false, err
	}

	return total > 0, nil
}

func (n *SQLDB) UpdateInstanceDefaultStatus(instanceTypeID, regionID string) error {
	query := fmt.Sprintf(`UPDATE %s SET status='' and updated_time=NOW() WHERE instance_type_id=? and region_id=?`, instancesDetailsTable)
	_, err := n.db.Exec(query, instanceTypeID, regionID)
	if err != nil {
		return err
	}

	return err
}

// LoadMyInstancesInfo   load  my server information
func (n *SQLDB) LoadMyInstancesInfo(userID string, limit, offset int64) (*types.MyInstanceResponse, error) {
	out := new(types.MyInstanceResponse)
	var infos []*types.MyInstance
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=?  order by created_time desc LIMIT ? OFFSET ?", myInstancesTable)
	if limit > loadInstancesDefaultLimit {
		limit = loadInstancesDefaultLimit
	}
	err := n.db.Select(&infos, query, userID, limit, offset)
	if err != nil {
		return nil, err
	}

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE user_id=?", myInstancesTable)
	var count int
	err = n.db.Get(&count, countQuery, userID)
	if err != nil {
		return nil, err
	}

	out.Total = count
	out.List = infos

	return out, nil
}

func (n *SQLDB) LoadInstanceDetailsInfo(userID, instanceId string) (*types.InstanceDetails, error) {
	var info types.InstanceDetails
	query := fmt.Sprintf("SELECT region_id,instance_id,instance_name,instance_type,image_id,security_group_id,instance_charge_type,internet_charge_type,bandwidth_out,bandwidth_in,system_disk_size,ip_address,system_disk_category,created_time,memory,memory_used,cores,cores_used,os_type,data_disk  FROM %s WHERE user_id=? and instance_id=?", instancesDetailsTable)
	err := n.db.Get(&info, query, userID, instanceId)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func (n *SQLDB) LoadInstanceDefaultInfo(req *types.InstanceTypeFromBaseReq) (*types.InstanceTypeResponse, error) {
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
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s and status!=''", instanceDefaultTable, query)
	var count int
	err := n.db.Get(&count, countQuery, args...)
	if err != nil {
		return nil, err
	}
	querySql := fmt.Sprintf("SELECT * FROM %s WHERE %s and status!='' LIMIT %d OFFSET %d ", instanceDefaultTable, query, req.Limit, req.Offset)
	err = n.db.Select(&info, querySql, args...)
	if err != nil {
		return nil, err
	}
	out.Total = count
	out.List = info
	return out, nil
}

func (n *SQLDB) LoadInstanceCpuInfo(req *types.InstanceTypeFromBaseReq) ([]*int32, error) {
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
	querySql := fmt.Sprintf("SELECT distinct cpu_core_count FROM %s WHERE %s order by cpu_core_count asc", instanceDefaultTable, query)
	err := n.db.Select(&info, querySql, args...)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (n *SQLDB) LoadInstanceMemoryInfo(req *types.InstanceTypeFromBaseReq) ([]*float32, error) {
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
	querySql := fmt.Sprintf("SELECT distinct memory_size FROM %s WHERE %s order by memory_size asc", instanceDefaultTable, query)
	err := n.db.Select(&info, querySql, args...)
	if err != nil {
		return nil, err
	}
	return info, nil
}
