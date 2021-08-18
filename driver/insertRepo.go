package driver

import (
	"context"
	"database/sql"
	"fmt"
	"manajemen-keuangan-koperasi/konstanta"
	"strings"
	"time"
)

var statementCreateZeroBalance string = fmt.Sprintf("INSERT INTO %s (uid,IP,IW,SS,SHU,Bonus) VALUES (?,?,?,?,?,?)", konstanta.TABLEMEMBALANCE)

func (DB *DBDriver) CreateZeroBalance(tx *sql.Tx, uid float64) (sql.Result, error) {
	return tx.Exec(statementCreateZeroBalance, uid, 0, 0, 0, 0, 0)
}

var statementInsertUser string = fmt.Sprintf("INSERT INTO %s (member_id, username, passwd, role, isagent) VALUES (?,?,?,?,?)", konstanta.TABLEALLUSER)

func (DB *DBDriver) InsertUserTx(tx *sql.Tx, MemID string, Username string, Passwd string, Role string, IsAgent string) (sql.Result, error) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return DB.DB.ExecContext(ctx, statementInsertUser, MemID, Username, Passwd, Role, IsAgent)
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

func (DB *DBDriver) InsertMemMurobahahTx(tx *sql.Tx, uid float64, amount float64, doc string, duedate string, info string, approvedby float64) (sql.Result, error) {

	ctx, cancel := context.WithCancel(context.Background())

	date := time.Now()

	defer cancel()
	return tx.ExecContext(ctx, statementInsertMemMurobahahTx, uid, date, amount, 0, doc, duedate, info, approvedby)
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

var statementRawInsertBatchAllInfo string = fmt.Sprintf("INSERT INTO %s (uid,date,info)", konstanta.TABLEALLINFO)

func (DB *DBDriver) InsertBatchAllInfoTx(tx *sql.Tx, uid []float64, info string) (sql.Result, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	date := time.Now()

	valueStrings := make([]string, 0, len(uid))
	valueArgs := make([]interface{}, 0, len(uid)*3)

	for _, v := range uid {
		valueStrings = append(valueStrings, "(?,?,?)")
		valueArgs = append(valueArgs, v)
		valueArgs = append(valueArgs, date)
		valueArgs = append(valueArgs, info)
	}

	var statementReadyInsertBatchAllInfo string = fmt.Sprintf("%s VALUES %s", statementRawInsertBatchAllInfo, strings.Join(valueStrings, ","))

	return tx.ExecContext(ctx, statementReadyInsertBatchAllInfo, valueArgs...)
}

var statementInsertAgentHistory string = fmt.Sprintf("INSERT INTO %s (uid,murobahah_id) VALUES (?,?)", konstanta.TABLEAGENTHISTORY)

func (DB *DBDriver) InsertAgentHistoryTx(tx *sql.Tx, agentid float64, murobahahid float64) (sql.Result, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	return tx.ExecContext(ctx, statementInsertAgentHistory, agentid, murobahahid)
}
