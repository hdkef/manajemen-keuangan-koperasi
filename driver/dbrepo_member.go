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

var statementInsertMemReq string = fmt.Sprintf("INSERT INTO %s (uid,date,type,amount,document,due_date,info) VALUES (?,?,?,?,?,?,?)", konstanta.TABLEMEMREQ)

func (DB *DBDriver) InsertMemReq(uid float64, type_ string, amount float64, option MemReqOption) (sql.Result, error) {

	ctx, cancel := context.WithCancel(context.Background())

	date := time.Now()

	defer cancel()
	return DB.DB.ExecContext(ctx, statementInsertMemReq, uid, date, type_, amount, option.Doc, option.DueDate, option.Info)
}

var statementInsertMemDebtTx string = fmt.Sprintf("INSERT INTO %s (uid,date,initial,paid,document,due_date,info,approvedby) VALUES (?,?,?,?,?,?,?,?)", konstanta.TABLEMEMDEBT)

func (DB *DBDriver) InsertMemDebtTx(tx *sql.Tx, uid float64, amount float64, option MemReqOption, approvedby float64) (sql.Result, error) {

	ctx, cancel := context.WithCancel(context.Background())

	date := time.Now()

	defer cancel()
	return tx.ExecContext(ctx, statementInsertMemDebtTx, uid, date, amount, 0, option.Doc, option.DueDate, option.Info, approvedby)
}

var statementInsertMemJournalTx string = fmt.Sprintf("INSERT INTO %s (uid,date,type,amount,info,approvedby) VALUES (?,?,?,?,?,?)", konstanta.TABLEMEMJOURNAL)

func (DB *DBDriver) InsertMemJournalTx(tx *sql.Tx, uid float64, type_ string, amount float64, info string, approvedby float64) (sql.Result, error) {

	ctx, cancel := context.WithCancel(context.Background())

	date := time.Now()

	defer cancel()
	return tx.ExecContext(ctx, statementInsertMemJournalTx, uid, date, type_, amount, info, approvedby)

}

var statementSS string = fmt.Sprintf("UPDATE %s SET %s = %s + ? WHERE uid = ?", konstanta.TABLEMEMBALANCE, "SS", "SS")
var statementIP string = fmt.Sprintf("UPDATE %s SET %s = %s + ? WHERE uid = ?", konstanta.TABLEMEMBALANCE, "IP", "IP")
var statementIW string = fmt.Sprintf("UPDATE %s SET %s = %s + ? WHERE uid = ?", konstanta.TABLEMEMBALANCE, "IW", "IW")

func (DB *DBDriver) ModifyMemBalanceTx(tx *sql.Tx, type_ string, amount float64, uid float64) (sql.Result, error) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	switch type_ {
	case konstanta.TypeIP:
		return tx.ExecContext(ctx, statementIP, amount, uid)
	case konstanta.TypeIW:
		return tx.ExecContext(ctx, statementIW, amount, uid)
	case konstanta.TypeSSPos:
		return tx.ExecContext(ctx, statementSS, amount, uid)
	default:
		return nil, errors.New("ERROR")
	}
}

var statementFindMemReqTx string = "SELECT alluser.id, alluser.member_id, alluser.username, member_req.id, member_req.date, member_req.type, member_req.amount, member_req.document, member_req.due_date, member_req.info from member_req JOIN alluser ON member_req.uid = alluser.id"

func (DB *DBDriver) FindMemReq() ([]models.MemReq, error) {

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var memreq []models.MemReq

	row, err := DB.DB.QueryContext(ctx, statementFindMemReqTx)
	if err != nil {
		return memreq, err
	}

	for row.Next() {
		var tmpmemreq models.MemReq
		err = row.Scan(&tmpmemreq.User.ID, &tmpmemreq.User.MemID, &tmpmemreq.User.Username, &tmpmemreq.ID, &tmpmemreq.Date, &tmpmemreq.Type, &tmpmemreq.Amount, &tmpmemreq.Doc, &tmpmemreq.DueDate, &tmpmemreq.Info)
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
