package db

import (
	"fmt"

	"github.com/LMF709268224/titan-vps/api/types"
)

// SaveRechargeInfo save recharge information
func (n *SQLDB) SaveRechargeInfo(rInfo *types.RechargeRecord) error {
	// update record table
	query := fmt.Sprintf(
		`INSERT INTO %s (order_id, from_addr, to_addr, value, created_height, done_height, state, recharge_addr, recharge_hash, msg, user_addr, tx_hash) 
		        VALUES (:order_id, :from_addr, :to_addr, :value, :created_height, :done_height, :state, :recharge_addr, :recharge_hash, :msg, :user_addr, :tx_hash)`, rechargeRecordTable)
	_, err := n.db.NamedExec(query, rInfo)

	return err
}

// UpdateRechargeRecord update recharge record information
func (n *SQLDB) UpdateRechargeRecord(info *types.RechargeRecord, newState types.RechargeState) error {
	query := fmt.Sprintf(`UPDATE %s SET state=?, done_time=NOW(), done_height=?, recharge_hash=?, msg=? WHERE order_id=? AND state=?`, rechargeRecordTable)
	_, err := n.db.Exec(query, newState, info.DoneHeight, info.RechargeHash, info.Msg, info.OrderID, info.State)

	return err
}

// LoadRechargeRecord load recharge record information
func (n *SQLDB) LoadRechargeRecord(orderID string) (*types.RechargeRecord, error) {
	var info types.RechargeRecord
	query := fmt.Sprintf("SELECT * FROM %s WHERE order_id=?", rechargeRecordTable)
	err := n.db.Get(&info, query, orderID)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// LoadRechargeRecords load the recharge records from the incoming scheduler
func (n *SQLDB) LoadRechargeRecords(state types.RechargeState) ([]*types.RechargeRecord, error) {
	var infos []*types.RechargeRecord
	query := fmt.Sprintf("SELECT * FROM %s WHERE state=? ", rechargeRecordTable)

	err := n.db.Select(&infos, query, state)
	if err != nil {
		return nil, err
	}

	return infos, nil
}

// SaveWithdrawInfo save withdraw information
func (n *SQLDB) SaveWithdrawInfo(rInfo *types.WithdrawRecord) error {
	// update record table
	query := fmt.Sprintf(
		`INSERT INTO %s (order_id, from_addr, to_addr, value, created_height, done_height, state, withdraw_addr, withdraw_hash, msg, user_addr, tx_hash) 
		        VALUES (:order_id, :from_addr, :to_addr, :value, :created_height, :done_height, :state, :withdraw_addr, :withdraw_hash, :msg, :user_addr, :tx_hash)`, withdrawRecordTable)
	_, err := n.db.NamedExec(query, rInfo)

	return err
}

// LoadWithdrawRecord load withdraw record information
func (n *SQLDB) LoadWithdrawRecord(orderID string) (*types.WithdrawRecord, error) {
	var info types.WithdrawRecord
	query := fmt.Sprintf("SELECT * FROM %s WHERE order_id=?", withdrawRecordTable)
	err := n.db.Get(&info, query, orderID)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// UpdateWithdrawRecord update withdraw record information
func (n *SQLDB) UpdateWithdrawRecord(info *types.WithdrawRecord, newState types.WithdrawState) error {
	query := fmt.Sprintf(`UPDATE %s SET state=?, value=?, done_time=NOW(), from_addr=?,
	    done_height=?, tx_hash=?, withdraw_hash=?, msg=? WHERE order_id=? AND state=?`, withdrawRecordTable)
	_, err := n.db.Exec(query, newState, info.Value, info.From, info.DoneHeight, info.TxHash, info.WithdrawHash, info.Msg, info.OrderID, info.State)

	return err
}

// LoadWithdrawRecords load the withdraw records from the incoming scheduler
func (n *SQLDB) LoadWithdrawRecords(state types.WithdrawState) ([]*types.WithdrawRecord, error) {
	var infos []*types.WithdrawRecord
	query := fmt.Sprintf("SELECT * FROM %s WHERE state=? ", withdrawRecordTable)

	err := n.db.Select(&infos, query, state)
	if err != nil {
		return nil, err
	}

	return infos, nil
}

// LoadWithdrawRecordsByUser load records
func (n *SQLDB) LoadWithdrawRecordsByUser(userAddr string, limit, offset int64) (*types.WithdrawResponse, error) {
	out := new(types.WithdrawResponse)

	var infos []*types.WithdrawRecord
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_addr=? order by created_time desc LIMIT ? OFFSET ?", withdrawRecordTable)
	if limit > loadOrderRecordsDefaultLimit {
		limit = loadOrderRecordsDefaultLimit
	}

	err := n.db.Select(&infos, query, userAddr, limit, offset)
	if err != nil {
		return nil, err
	}

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE user_addr=?", withdrawRecordTable)
	var count int
	err = n.db.Get(&count, countQuery, userAddr)
	if err != nil {
		return nil, err
	}

	out.Total = count
	out.List = infos

	return out, nil
}

// LoadRechargeRecordsByUser load records
func (n *SQLDB) LoadRechargeRecordsByUser(userAddr string, limit, offset int64) (*types.RechargeResponse, error) {
	out := new(types.RechargeResponse)

	var infos []*types.RechargeRecord
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_addr=? order by created_time desc LIMIT ? OFFSET ?", rechargeRecordTable)
	if limit > loadOrderRecordsDefaultLimit {
		limit = loadOrderRecordsDefaultLimit
	}

	err := n.db.Select(&infos, query, userAddr, limit, offset)
	if err != nil {
		return nil, err
	}

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE user_addr=?", rechargeRecordTable)
	var count int
	err = n.db.Get(&count, countQuery, userAddr)
	if err != nil {
		return nil, err
	}

	out.Total = count
	out.List = infos

	return out, nil
}
