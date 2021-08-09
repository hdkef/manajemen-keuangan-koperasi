package driver

import (
	"context"
	"database/sql"
	"fmt"
	"manajemen-keuangan-koperasi/konstanta"
	"time"
)

type MemReqOption struct {
	Doc     string
	DueDate time.Time
	Info    string
}

func (DB *DBDriver) InsertMemReq(memid interface{}, type_ string, amount float64, option MemReqOption) (sql.Result, error) {

	statement := fmt.Sprintf("INSERT INTO %s (mem_id,date,type,amount,document,due_date,info) VALUES (?,?,?,?,?,?,?)", konstanta.TABLEMEMREQ)
	ctx, cancel := context.WithCancel(context.Background())

	date := time.Now()

	defer cancel()
	return DB.DB.ExecContext(ctx, statement, memid, date, type_, amount, option.Doc, option.DueDate, option.Info)
}

func (DB *DBDriver) InsertMemDebtTx(tx *sql.Tx, memid interface{}, amount float64, option MemReqOption, approvedby interface{}) (sql.Result, error) {

	statement := fmt.Sprintf("INSERT INTO %s (mem_id,date,initial,paid,document,due_date,info,approvedby) VALUES (?,?,?,?,?,?,?,?)", konstanta.TABLEMEMDEBT)
	ctx, cancel := context.WithCancel(context.Background())

	date := time.Now()

	defer cancel()
	return tx.ExecContext(ctx, statement, memid, date, amount, 0, option.Doc, option.DueDate, option.Info, approvedby)
}

func (DB *DBDriver) InsertMemJournalTx(tx *sql.Tx, memid interface{}, type_ string, amount float64, info string, approvedby interface{}) (sql.Result, error) {

	statement := fmt.Sprintf("INSERT INTO %s (mem_id,date,type,amount,info,approvedby) VALUES (?,?,?,?,?,?)", konstanta.TABLEMEMJOURNAL)
	ctx, cancel := context.WithCancel(context.Background())

	date := time.Now()

	defer cancel()
	return tx.ExecContext(ctx, statement, memid, date, type_, amount, info, approvedby)

}

func (DB *DBDriver) ModifyMemBalanceTx(tx *sql.Tx, type_ string, amount float64) {

}
