package driver

import (
	"context"
	"database/sql"
	"fmt"
	"manajemen-keuangan-koperasi/konstanta"
)

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

var statementDeleteMemReqMurobahah string = fmt.Sprintf("DELETE FROM %s WHERE id=?", konstanta.TABLEMEMREQMUROBAHAH)

func (DB *DBDriver) DeleteMemReqMurobahahTx(tx *sql.Tx, id float64) (sql.Result, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	return tx.ExecContext(ctx, statementDeleteMemReqMurobahah, id)
}

var statementDeleteAllInfo string = fmt.Sprintf("DELETE FROM %s WHERE id=?", konstanta.TABLEALLINFO)

func (DB *DBDriver) DeleteAllInfoTx(tx *sql.Tx, id float64) (sql.Result, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	return tx.ExecContext(ctx, statementDeleteAllInfo, id)
}
