package db

import (
	"context"
	"fmt"

	"github.com/LMF709268224/titan-vps/api/types"
	"github.com/jmoiron/sqlx"
)

// SaveOrderInfo save order information
func (n *SQLDB) SaveOrderInfo(rInfo *types.OrderRecord) error {
	// update record table
	query := fmt.Sprintf(
		`INSERT INTO %s (order_id, from_addr, to_addr, value, created_height, done_height, state, done_state, vps_id, msg) 
		        VALUES (:order_id, :from_addr, :to_addr, :value, :created_height, :done_height, :state, :done_state, :vps_id, :msg)
				ON DUPLICATE KEY UPDATE state=:state, done_height=:done_height, done_state=:done_state, done_time=NOW(),
				from_addr=:from_addr, to_addr=:to_addr, value=:value, vps_id=:vps_id, created_height=:created_height, msg=:msg`, orderRecordTable)
	_, err := n.db.NamedExec(query, rInfo)

	return err
}

// LoadOrderRecord load order record information
func (n *SQLDB) LoadOrderRecord(orderID string) (*types.OrderRecord, error) {
	var info types.OrderRecord
	query := fmt.Sprintf("SELECT * FROM %s WHERE order_id=?", orderRecordTable)
	err := n.db.Get(&info, query, orderID)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// OrderExists checks if an order exists in the state machine table of the specified server.
func (n *SQLDB) OrderExists(orderID string) (bool, error) {
	var total int64
	countSQL := fmt.Sprintf(`SELECT count(order_id) FROM %s WHERE order_id=? `, orderRecordTable)
	if err := n.db.Get(&total, countSQL, orderID); err != nil {
		return false, err
	}

	return total > 0, nil
}

// LoadOrderRecords load the order records from the incoming scheduler
func (n *SQLDB) LoadOrderRecords(statuses []int64, limit, offset int) (*sqlx.Rows, error) {
	if limit > loadOrderRecordsDefaultLimit || limit == 0 {
		limit = loadOrderRecordsDefaultLimit
	}
	sQuery := fmt.Sprintf(`SELECT * FROM %s WHERE state in (?) order by order_id asc LIMIT ? OFFSET ?`, orderRecordTable)
	query, args, err := sqlx.In(sQuery, statuses, limit, offset)
	if err != nil {
		return nil, err
	}

	query = n.db.Rebind(query)
	return n.db.QueryxContext(context.Background(), query, args...)
}

// LoadOrderCount count order
func (n *SQLDB) LoadOrderCount() (int, error) {
	var size int
	cmd := fmt.Sprintf("SELECT count(order_id) FROM %s", orderRecordTable)
	err := n.db.Get(&size, cmd)
	if err != nil {
		return 0, err
	}
	return size, nil
}

// LoadAllOrderRecords loads all order records
func (n *SQLDB) LoadAllOrderRecords(limit, offset int, statuses []int64) (*sqlx.Rows, error) {
	sQuery := fmt.Sprintf(`SELECT * FROM %s WHERE state in (?) order by order_id asc limit ? offset ?`, orderRecordTable)
	query, args, err := sqlx.In(sQuery, statuses, limit, offset)
	if err != nil {
		return nil, err
	}

	query = n.db.Rebind(query)
	return n.db.QueryxContext(context.Background(), query, args...)
}
