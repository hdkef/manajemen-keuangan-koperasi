package driver

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"manajemen-keuangan-koperasi/konstanta"
)

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
