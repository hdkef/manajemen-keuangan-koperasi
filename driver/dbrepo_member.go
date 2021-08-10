package driver

import (
	"context"
	"database/sql"
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

var statementInsertMemReq string = fmt.Sprintf("INSERT INTO %s (mem_id,date,type,amount,document,due_date,info) VALUES (?,?,?,?,?,?,?)", konstanta.TABLEMEMREQ)

func (DB *DBDriver) InsertMemReq(memid interface{}, type_ string, amount float64, option MemReqOption) (sql.Result, error) {

	ctx, cancel := context.WithCancel(context.Background())

	date := time.Now()

	defer cancel()
	return DB.DB.ExecContext(ctx, statementInsertMemReq, memid, date, type_, amount, option.Doc, option.DueDate, option.Info)
}

var statementInsertMemDebtTx string = fmt.Sprintf("INSERT INTO %s (mem_id,date,initial,paid,document,due_date,info,approvedby) VALUES (?,?,?,?,?,?,?,?)", konstanta.TABLEMEMDEBT)

func (DB *DBDriver) InsertMemDebtTx(tx *sql.Tx, memid interface{}, amount float64, option MemReqOption, approvedby interface{}) (sql.Result, error) {

	ctx, cancel := context.WithCancel(context.Background())

	date := time.Now()

	defer cancel()
	return tx.ExecContext(ctx, statementInsertMemDebtTx, memid, date, amount, 0, option.Doc, option.DueDate, option.Info, approvedby)
}

var statementInsertMemJournalTx string = fmt.Sprintf("INSERT INTO %s (mem_id,date,type,amount,info,approvedby) VALUES (?,?,?,?,?,?)", konstanta.TABLEMEMJOURNAL)

func (DB *DBDriver) InsertMemJournalTx(tx *sql.Tx, memid interface{}, type_ string, amount float64, info string, approvedby interface{}) (sql.Result, error) {

	ctx, cancel := context.WithCancel(context.Background())

	date := time.Now()

	defer cancel()
	return tx.ExecContext(ctx, statementInsertMemJournalTx, memid, date, type_, amount, info, approvedby)

}

func (DB *DBDriver) ModifyMemBalanceTx(tx *sql.Tx, type_ string, amount float64) {

}

var statementFindMemReqTx string = "SELECT alluser.id, alluser.member_id, alluser.username, member_req.id, member_req.date, member_req.type, member_req.amount, member_req.document, member_req.due_date, member_req.info from member_req JOIN alluser ON member_req.mem_id = alluser.id"

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
