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

var ModifyMemberRoleByUID string = fmt.Sprintf("UPDATE %s SET role=? WHERE id=?", konstanta.TABLEALLUSER)
var ModifyMemberIsAgentByUID string = fmt.Sprintf("UPDATE %s SET isagent=? WHERE id=?", konstanta.TABLEALLUSER)

func (DB *DBDriver) ModifyMemberFieldByUIDTx(tx *sql.Tx, uid float64, type_ string, value string) (sql.Result, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	switch type_ {
	case konstanta.TYPEIsAgent:
		return tx.ExecContext(ctx, ModifyMemberIsAgentByUID, value, uid)
	case konstanta.TYPERole:
		return tx.ExecContext(ctx, ModifyMemberRoleByUID, value, uid)
	default:
		return nil, errors.New("ERROR")
	}
}
