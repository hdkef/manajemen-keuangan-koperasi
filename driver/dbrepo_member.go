package driver

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"manajemen-keuangan-koperasi/konstanta"
	"manajemen-keuangan-koperasi/models"
	"time"
)

type MemReqOption struct {
	Doc     string
	DueDate time.Time
	Info    string
}

var statementInsertMemReq string = fmt.Sprintf("INSERT INTO %s (uid,date,type,amount,info) VALUES (?,?,?,?,?)", konstanta.TABLEMEMREQ)

func (DB *DBDriver) InsertMemReq(uid float64, type_ string, amount float64, info string) (sql.Result, error) {

	ctx, cancel := context.WithCancel(context.Background())

	date := time.Now()

	defer cancel()
	return DB.DB.ExecContext(ctx, statementInsertMemReq, uid, date, type_, amount, info)
}

func (DB *DBDriver) InsertMemReqTx(tx *sql.Tx, uid float64, type_ string, amount float64, info string) (sql.Result, error) {

	ctx, cancel := context.WithCancel(context.Background())

	date := time.Now()

	defer cancel()
	return tx.ExecContext(ctx, statementInsertMemReq, uid, date, type_, amount, info)
}

var statementInsertMemMurobahahTx string = fmt.Sprintf("INSERT INTO %s (uid,date,initial,paid,document,due_date,info,approvedby) VALUES (?,?,?,?,?,?,?,?)", konstanta.TABLEMEMMUROBAHAH)

func (DB *DBDriver) InsertMemMurobahahTx(tx *sql.Tx, uid float64, amount float64, option MemReqOption, approvedby float64) (sql.Result, error) {

	ctx, cancel := context.WithCancel(context.Background())

	date := time.Now()

	defer cancel()
	return tx.ExecContext(ctx, statementInsertMemMurobahahTx, uid, date, amount, 0, option.Doc, option.DueDate, option.Info, approvedby)
}

var statementInsertMemReqMurobahah = fmt.Sprintf("INSERT INTO %s (agent_id,buyer_id,date,amount,due_date,document,info) VALUES (?,?,?,?,?,?,?)", konstanta.TABLEMEMREQMUROBAHAH)

func (DB *DBDriver) InsertMemReqMurobahah(agentid float64, buyerid float64, duedate string, amount float64, doc string, info string) (sql.Result, error) {
	ctx, cancel := context.WithCancel(context.Background())

	date := time.Now()

	defer cancel()

	return DB.DB.ExecContext(ctx, statementInsertMemReqMurobahah, agentid, buyerid, date, amount, duedate, doc, info)
}

var statementInsertMemJournalTx string = fmt.Sprintf("INSERT INTO %s (uid,date,type,amount,info,approvedby) VALUES (?,?,?,?,?,?)", konstanta.TABLEMEMJOURNAL)

func (DB *DBDriver) InsertMemJournalTx(tx *sql.Tx, uid float64, type_ string, amount float64, info string, approvedby float64) (sql.Result, error) {

	ctx, cancel := context.WithCancel(context.Background())

	date := time.Now()

	defer cancel()
	return tx.ExecContext(ctx, statementInsertMemJournalTx, uid, date, type_, amount, info, approvedby)

}

var statementFindLimitedMemJournalTx string = fmt.Sprintf("SELECT id,date,type,amount,info,approvedby FROM %s WHERE uid=? LIMIT ?", konstanta.TABLEMEMJOURNAL)

func (DB *DBDriver) FindLimitedMemJournalByUIDTx(tx *sql.Tx, uid float64, limit int) ([]models.MemTransaction, error) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	row, err := tx.QueryContext(ctx, statementFindLimitedMemJournalTx, uid, limit)
	if err != nil {
		return nil, err
	}

	var journals []models.MemTransaction

	for row.Next() {
		var tmp models.MemTransaction
		err = row.Scan(&tmp.ID, &tmp.Date, &tmp.Type, &tmp.Amount, &tmp.Info, &tmp.ApprovedBy)
		if err != nil {
			return nil, err
		}
		journals = append(journals, tmp)
	}

	return journals, nil
}

var statementSSPos string = fmt.Sprintf("UPDATE %s SET %s = %s + ? WHERE uid = ?", konstanta.TABLEMEMBALANCE, "SS", "SS")
var statementSSNeg string = fmt.Sprintf("UPDATE %s SET %s = %s - ? WHERE uid = ?", konstanta.TABLEMEMBALANCE, "SS", "SS")
var statementIP string = fmt.Sprintf("UPDATE %s SET %s = %s + ? WHERE uid = ?", konstanta.TABLEMEMBALANCE, "IP", "IP")
var statementIW string = fmt.Sprintf("UPDATE %s SET %s = %s + ? WHERE uid = ?", konstanta.TABLEMEMBALANCE, "IW", "IW")
var statementSHU string = fmt.Sprintf("UPDATE %s SET %s = %s + ? WHERE uid = ?", konstanta.TABLEMEMBALANCE, "SHU", "SHU")
var statementBonus string = fmt.Sprintf("UPDATE %s SET %s = %s + ? WHERE uid = ?", konstanta.TABLEMEMBALANCE, "Bonus", "Bonus")

func (DB *DBDriver) ModifyMemBalanceTx(tx *sql.Tx, type_ string, amount float64, uid float64) (sql.Result, error) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	switch type_ {
	case konstanta.TypeIP:
		return tx.ExecContext(ctx, statementIP, amount, uid)
	case konstanta.TypeIW:
		return tx.ExecContext(ctx, statementIW, amount, uid)
	case konstanta.TypeSSPos:
		return tx.ExecContext(ctx, statementSSPos, amount, uid)
	case konstanta.TypeSSNeg:
		return tx.ExecContext(ctx, statementSSNeg, amount, uid)
	case konstanta.TypeSHU:
		return tx.ExecContext(ctx, statementSHU, amount, uid)
	case konstanta.TypeBonus:
		return tx.ExecContext(ctx, statementBonus, amount, uid)
	default:
		return nil, errors.New("ERROR")
	}
}

var statementFindMemBalanceByUID string = fmt.Sprintf("SELECT IP,IW,SS,SHU,Bonus FROM %s WHERE uid=?", konstanta.TABLEMEMBALANCE)

func (DB *DBDriver) FindMemBalanceByUIDTx(tx *sql.Tx, uid float64) (models.MemBalance, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var tmp models.MemBalance

	err := tx.QueryRowContext(ctx, statementFindMemBalanceByUID, uid).Scan(&tmp.IP, &tmp.IW, &tmp.SS, &tmp.Bonus, &tmp.Bonus)
	if err != nil {
		return models.MemBalance{}, err
	}
	return tmp, nil
}

var statementFindMemSSBalance = fmt.Sprintf("SELECT SS from %s WHERE uid = ?", konstanta.TABLEMEMBALANCE)

func (DB *DBDriver) FindMemSSBalanceTx(tx *sql.Tx, uid float64) (float64, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var amt float64

	err := tx.QueryRowContext(ctx, statementFindMemSSBalance, uid).Scan(&amt)
	if err != nil {
		return 0, err
	}
	return amt, nil
}

var statementFindMemReqTx string = "SELECT alluser.id, alluser.member_id, alluser.username, member_req.id, member_req.date, member_req.type, member_req.amount, member_req.info from member_req JOIN alluser ON member_req.uid = alluser.id"

func (DB *DBDriver) FindMemReqTx(tx *sql.Tx) ([]models.MemReq, error) {

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var memreq []models.MemReq

	row, err := tx.QueryContext(ctx, statementFindMemReqTx)
	if err != nil {
		return memreq, err
	}

	for row.Next() {
		var tmpmemreq models.MemReq
		err = row.Scan(&tmpmemreq.User.ID, &tmpmemreq.User.MemID, &tmpmemreq.User.Username, &tmpmemreq.ID, &tmpmemreq.Date, &tmpmemreq.Type, &tmpmemreq.Amount, &tmpmemreq.Info)
		if err != nil {
			return memreq, err
		}
		memreq = append(memreq, tmpmemreq)
	}
	return memreq, nil
}

var statementDeleteMemReq string = fmt.Sprintf("DELETE FROM %s WHERE id = ?", konstanta.TABLEMEMREQ)

func (DB *DBDriver) DeleteMemReq(id float64) (sql.Result, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	return DB.DB.ExecContext(ctx, statementDeleteMemReq, id)
}

func (DB *DBDriver) DeleteMemReqTx(tx *sql.Tx, id float64) (sql.Result, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	return tx.ExecContext(ctx, statementDeleteMemReq, id)
}

var statementFindMemReqMurobahah string = fmt.Sprintf("SELECT member_req_murobahah.id,  agent.ID, agent.Username, buyer.ID,buyer.Username, member_req_murobahah.date, member_req_murobahah.due_date, member_req_murobahah.amount, member_req_murobahah.info, member_req_murobahah.document FROM %s JOIN %s AS agent ON agent_id = agent.id JOIN %s AS buyer ON buyer_id = buyer.id", konstanta.TABLEMEMREQMUROBAHAH, konstanta.TABLEALLUSER, konstanta.TABLEALLUSER)

func (DB *DBDriver) FindMemReqMurobahahTx(tx *sql.Tx) ([]models.Murobahah, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	row, err := tx.QueryContext(ctx, statementFindMemReqMurobahah)
	if err != nil {
		return nil, err
	}

	var murobahahs []models.Murobahah

	for row.Next() {
		var tmp models.Murobahah
		err = row.Scan(&tmp.ID, &tmp.Agent.ID, &tmp.Agent.Username, &tmp.Buyer.ID, &tmp.Buyer.Username, &tmp.Date, &tmp.DueDate, &tmp.Amount, &tmp.Info, &tmp.Doc)
		if err != nil {
			return nil, err
		}
		murobahahs = append(murobahahs, tmp)
	}
	return murobahahs, nil
}

var statementDeleteMemReqMurobahah string = fmt.Sprintf("DELETE FROM %s WHERE id=?", konstanta.TABLEMEMREQMUROBAHAH)

func (DB *DBDriver) DeleteMemReqMurobahahTx(tx *sql.Tx, id float64) (sql.Result, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	return tx.ExecContext(ctx, statementDeleteMemReqMurobahah, id)
}

var statementInsertAllInfoTx string = fmt.Sprintf("INSERT INTO %s (uid,date,info) VALUES (?,?,?)", konstanta.TABLEALLINFO)

func (DB *DBDriver) InsertAllInfoTx(tx *sql.Tx, uid float64, info string) (sql.Result, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	date := time.Now()

	return tx.ExecContext(ctx, statementInsertAllInfoTx, uid, date, info)
}

var statementFindAllInfo string = fmt.Sprintf("SELECT id,date,info FROM %s WHERE uid=?", konstanta.TABLEALLINFO)

func (DB *DBDriver) FindAllInfoTx(tx *sql.Tx, uid float64) ([]models.AllInfo, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	row, err := tx.QueryContext(ctx, statementFindAllInfo, uid)
	if err != nil {
		return nil, err
	}

	var infos []models.AllInfo

	for row.Next() {
		var tmp models.AllInfo
		err = row.Scan(&tmp.ID, &tmp.Date, &tmp.Info)
		if err != nil {
			return nil, err
		}
		infos = append(infos, tmp)
	}
	return infos, nil
}

var statementDeleteAllInfo string = fmt.Sprintf("DELETE FROM %s WHERE id=?", konstanta.TABLEALLINFO)

func (DB *DBDriver) DeleteAllInfoTx(tx *sql.Tx, id float64) (sql.Result, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	return tx.ExecContext(ctx, statementDeleteAllInfo, id)
}
