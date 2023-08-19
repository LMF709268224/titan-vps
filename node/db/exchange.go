package db

import (
	"database/sql"
	"fmt"

	"github.com/LMF709268224/titan-vps/api/types"
)

// SaveRechargeRecordAndUserBalance save recharge information
func (n *SQLDB) SaveRechargeRecordAndUserBalance(rInfo *types.RechargeRecord, balance, oldBalance string) error {
	tx, err := n.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		err = tx.Rollback()
		if err != nil && err != sql.ErrTxDone {
			log.Errorf("SaveRechargeRecordAndUserBalance Rollback err:%s", err.Error())
		}
	}()

	query := fmt.Sprintf(
		`INSERT INTO %s (order_id, from_addr, to_addr, value, created_height, done_height, state,  user_id) 
		        VALUES (:order_id, :from_addr, :to_addr, :value, :created_height, :done_height, :state, :user_id)`, rechargeRecordTable)
	_, err = tx.NamedExec(query, rInfo)
	if err != nil {
		return err
	}

	query = fmt.Sprintf(`UPDATE %s SET balance=? WHERE user_id=? AND balance=?`, userTable)
	_, err = tx.Exec(query, balance, rInfo.UserID, oldBalance)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// RechargeRecordExists checks if an order exists
func (n *SQLDB) RechargeRecordExists(orderID string) (bool, error) {
	var total int64
	countSQL := fmt.Sprintf(`SELECT count(order_id) FROM %s WHERE order_id=? `, rechargeRecordTable)
	if err := n.db.Get(&total, countSQL, orderID); err != nil {
		return false, err
	}

	return total > 0, nil
}

// UpdateRechargeRecord update recharge record information
func (n *SQLDB) UpdateRechargeRecord(info *types.RechargeRecord, newState types.RechargeState) error {
	query := fmt.Sprintf(`UPDATE %s SET state=?, done_time=NOW(), done_height=? WHERE order_id=? AND state=?`, rechargeRecordTable)
	_, err := n.db.Exec(query, newState, info.DoneHeight, info.OrderID, info.State)

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

// SaveWithdrawInfoAndUserBalance save withdraw information
func (n *SQLDB) SaveWithdrawInfoAndUserBalance(rInfo *types.WithdrawRecord, balance, oldBalance string) error {
	tx, err := n.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		err = tx.Rollback()
		if err != nil && err != sql.ErrTxDone {
			log.Errorf("SaveWithdrawInfoAndUserBalance Rollback err:%s", err.Error())
		}
	}()

	query := fmt.Sprintf(
		`INSERT INTO %s (order_id, from_addr, to_addr, value, created_height, done_height, state, withdraw_addr, withdraw_hash,  user_id) 
		        VALUES (:order_id, :from_addr, :to_addr, :value, :created_height, :done_height, :state, :withdraw_addr, :withdraw_hash, :user_id)`, withdrawRecordTable)
	_, err = tx.NamedExec(query, rInfo)
	if err != nil {
		return err
	}

	query = fmt.Sprintf(`UPDATE %s SET balance=? WHERE user_id=? AND balance=?`, userTable)
	_, err = tx.Exec(query, balance, rInfo.UserID, oldBalance)
	if err != nil {
		return err
	}

	return tx.Commit()
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
	    done_height=?, withdraw_hash=?, executor=? WHERE order_id=? AND state=?`, withdrawRecordTable)
	_, err := n.db.Exec(query, newState, info.Value, info.From, info.DoneHeight, info.WithdrawHash, info.Executor, info.OrderID, info.State)

	return err
}

// LoadWithdrawRecords load the withdraw records from the incoming scheduler
func (n *SQLDB) LoadWithdrawRecords(limit, offset int64) (*types.WithdrawResponse, error) {
	out := new(types.WithdrawResponse)

	var infos []*types.WithdrawRecord
	query := fmt.Sprintf("SELECT * FROM %s order by created_time desc LIMIT ? OFFSET ?", withdrawRecordTable)
	if limit > loadOrderRecordsDefaultLimit {
		limit = loadOrderRecordsDefaultLimit
	}

	err := n.db.Select(&infos, query, limit, offset)
	if err != nil {
		return nil, err
	}

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM %s ", withdrawRecordTable)
	var count int
	err = n.db.Get(&count, countQuery)
	if err != nil {
		return nil, err
	}

	out.Total = count
	out.List = infos

	return out, nil
}

// LoadWithdrawRecordsByUser load records
func (n *SQLDB) LoadWithdrawRecordsByUser(userID string, limit, offset int64) (*types.WithdrawResponse, error) {
	out := new(types.WithdrawResponse)

	var infos []*types.WithdrawRecord
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=? order by created_time desc LIMIT ? OFFSET ?", withdrawRecordTable)
	if limit > loadOrderRecordsDefaultLimit {
		limit = loadOrderRecordsDefaultLimit
	}

	err := n.db.Select(&infos, query, userID, limit, offset)
	if err != nil {
		return nil, err
	}

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE user_id=?", withdrawRecordTable)
	var count int
	err = n.db.Get(&count, countQuery, userID)
	if err != nil {
		return nil, err
	}

	out.Total = count
	out.List = infos

	return out, nil
}

// LoadRechargeRecordsByUser load records
func (n *SQLDB) LoadRechargeRecordsByUser(userID string, limit, offset int64) (*types.RechargeResponse, error) {
	out := new(types.RechargeResponse)

	var infos []*types.RechargeRecord
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=? order by created_time desc LIMIT ? OFFSET ?", rechargeRecordTable)
	if limit > loadOrderRecordsDefaultLimit {
		limit = loadOrderRecordsDefaultLimit
	}

	err := n.db.Select(&infos, query, userID, limit, offset)
	if err != nil {
		return nil, err
	}

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE user_id=?", rechargeRecordTable)
	var count int
	err = n.db.Get(&count, countQuery, userID)
	if err != nil {
		return nil, err
	}

	out.Total = count
	out.List = infos

	return out, nil
}
