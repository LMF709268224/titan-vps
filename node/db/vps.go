package db

import (
	"fmt"

	"github.com/LMF709268224/titan-vps/api/types"
)

// LoadVpsInfo  load  vps information
func (n *SQLDB) LoadVpsInfo(vpsId string) (*types.CreateInstanceReq, error) {
	var info types.CreateInstanceReq
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=?", vpsInstanceTable)
	err := n.db.Get(&info, query, vpsId)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// VpsExists  checks if this vps info exists in the state machine table of the specified server.
func (n *SQLDB) VpsExists(vpsId string) (bool, error) {
	var total int64
	countSQL := fmt.Sprintf(`SELECT count(id) FROM %s WHERE id=? `, vpsInstanceTable)
	if err := n.db.Get(&total, countSQL, vpsId); err != nil {
		return false, err
	}

	return total > 0, nil
}

// SaveVpsInstance   saves vps info into the database.
func (n *SQLDB) SaveVpsInstance(rInfo *types.OrderRecord) error {
	query := fmt.Sprintf(
		`INSERT INTO %s (region_id, instance_type, dry_run, image_id, security_group_id, instanceCharge_type, period_unit, period, bandwidth_out,bandwidth_in) 
				VALUES (:region_id, :instance_type, :dry_run, :image_id, :security_group_id, :instanceCharge_type, :period_unit, :period, :bandwidth_out,bandwidth_in)`, vpsInstanceTable)

	_, err := n.db.NamedExec(query, rInfo)

	return err
}
