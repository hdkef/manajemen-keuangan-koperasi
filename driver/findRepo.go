package driver

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"manajemen-keuangan-koperasi/konstanta"
	"manajemen-keuangan-koperasi/models"
)

var statementFindOneUserByUsername string = fmt.Sprintf("SELECT id, member_id, username, passwd, role, isagent, tel FROM %s WHERE Username=?", konstanta.TABLEALLUSER)

func (DB *DBDriver) FindOneUserByUsername(value string) (models.User, error) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	row := DB.DB.QueryRowContext(ctx, statementFindOneUserByUsername, value)

	tmpUsr := models.User{}

	err := row.Scan(&tmpUsr.ID, &tmpUsr.MemID, &tmpUsr.Username, &tmpUsr.Pass, &tmpUsr.Role, &tmpUsr.IsAgent, &tmpUsr.Tel)
	if err != nil {
		return models.User{}, err
	}

	return tmpUsr, nil
}

var statementFindOneUserByUID string = fmt.Sprintf("SELECT id, member_id, username, passwd, role, isagent, tel FROM %s WHERE id=?", konstanta.TABLEALLUSER)

func (DB *DBDriver) FindOneUserByUIDTx(tx *sql.Tx, value float64) (models.User, error) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	row := tx.QueryRowContext(ctx, statementFindOneUserByUID, value)

	tmpUsr := models.User{}

	err := row.Scan(&tmpUsr.ID, &tmpUsr.MemID, &tmpUsr.Username, &tmpUsr.Pass, &tmpUsr.Role, &tmpUsr.IsAgent, &tmpUsr.Tel)
	if err != nil {
		return models.User{}, err
	}

	return tmpUsr, nil
}

var statementFindAllUserByUsername = fmt.Sprintf("SELECT id, member_id, username, role, isagent, tel FROM %s WHERE username LIKE ?", konstanta.TABLEALLUSER)
var statementFindAllUserByRole = fmt.Sprintf("SELECT id, member_id, username, role, isagent, tel FROM %s WHERE role LIKE ?", konstanta.TABLEALLUSER)
var statementFindAllUserByMemID = fmt.Sprintf("SELECT id, member_id, username, role, isagent, tel FROM %s WHERE member_id LIKE ?", konstanta.TABLEALLUSER)
var statementFindAllUserByIsAgent = fmt.Sprintf("SELECT id, member_id, username, role, isagent, tel FROM %s WHERE isagent LIKE ?", konstanta.TABLEALLUSER)

func (DB *DBDriver) FindAllUserByFilter(filter string, key string) ([]models.User, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var row *sql.Rows
	var err error

	switch filter {
	case konstanta.QueryUsername:
		row, err = DB.DB.QueryContext(ctx, statementFindAllUserByUsername, key)
	case konstanta.QueryMemID:
		row, err = DB.DB.QueryContext(ctx, statementFindAllUserByMemID, key)
	case konstanta.QueryRole:
		row, err = DB.DB.QueryContext(ctx, statementFindAllUserByRole, key)
	case konstanta.QueryIsAgent:
		row, err = DB.DB.QueryContext(ctx, statementFindAllUserByIsAgent, key)
	default:
		return nil, errors.New("no filter match")
	}

	if err != nil {
		return nil, err
	}

	var users []models.User

	for row.Next() {
		var tmpUsr models.User
		err = row.Scan(&tmpUsr.ID, &tmpUsr.MemID, &tmpUsr.Username, &tmpUsr.Role, &tmpUsr.IsAgent, &tmpUsr.Tel)
		if err != nil {
			return nil, err
		}
		users = append(users, tmpUsr)
	}

	return users, nil
}

var statementFindLimitedMemJournalTx string = fmt.Sprintf("SELECT id,date,type,amount,info,approvedby FROM %s WHERE uid=? ORDER BY id DESC LIMIT ?", konstanta.TABLEMEMJOURNAL)

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

var statementFindMemReqMurobahah string = fmt.Sprintf("SELECT member_req_murobahah.id,  agent.ID, agent.Username, buyer.ID,buyer.Username, member_req_murobahah.date, member_req_murobahah.due_date, member_req_murobahah.amount, member_req_murobahah.info, member_req_murobahah.document FROM %s JOIN %s AS agent ON agent_id = agent.id JOIN %s AS buyer ON buyer_id = buyer.id", konstanta.TABLEMEMREQMUROBAHAH, konstanta.TABLEALLUSER, konstanta.TABLEALLUSER)

func (DB *DBDriver) FindMemReqMurobahahTx(tx *sql.Tx) ([]models.MurobahahReq, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	row, err := tx.QueryContext(ctx, statementFindMemReqMurobahah)
	if err != nil {
		return nil, err
	}

	var murobahahs []models.MurobahahReq

	for row.Next() {
		var tmp models.MurobahahReq
		err = row.Scan(&tmp.ID, &tmp.Agent.ID, &tmp.Agent.Username, &tmp.Buyer.ID, &tmp.Buyer.Username, &tmp.Date, &tmp.DueDate, &tmp.Amount, &tmp.Info, &tmp.Doc)
		if err != nil {
			return nil, err
		}
		murobahahs = append(murobahahs, tmp)
	}
	return murobahahs, nil
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

var statementFindMemMurobahah string = fmt.Sprintf("SELECT id,date,due_date,initial,paid,document,info FROM %s WHERE uid=?", konstanta.TABLEMEMMUROBAHAH)

func (DB *DBDriver) FindMemMurobahahTx(tx *sql.Tx, uid float64) ([]models.MemMurobahah, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var murobahahs []models.MemMurobahah

	row, err := tx.QueryContext(ctx, statementFindMemMurobahah, uid)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var tmp models.MemMurobahah
		err = row.Scan(&tmp.ID, &tmp.Date, &tmp.DueDate, &tmp.Initial, &tmp.Paid, &tmp.Doc, &tmp.Info)
		if err != nil {
			return nil, err
		}
		murobahahs = append(murobahahs, tmp)
	}

	return murobahahs, nil
}

func (DB *DBDriver) FindMemMurobahah(uid float64) ([]models.MemMurobahah, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var murobahahs []models.MemMurobahah

	row, err := DB.DB.QueryContext(ctx, statementFindMemMurobahah, uid)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var tmp models.MemMurobahah
		err = row.Scan(&tmp.ID, &tmp.Date, &tmp.DueDate, &tmp.Initial, &tmp.Paid, &tmp.Doc, &tmp.Info)
		if err != nil {
			return nil, err
		}
		murobahahs = append(murobahahs, tmp)
	}

	return murobahahs, nil
}

var statementFindMemMurobahahPayReq string = fmt.Sprintf("SELECT member_murobahah_payreq.id,member_murobahah_payreq.date,member_murobahah_payreq.murobahah_id,member_murobahah_payreq.amount,member_murobahah_payreq.info,murobahah.document FROM %s JOIN %s as murobahah ON murobahah.id = member_murobahah_payreq.murobahah_id", konstanta.TABLEMEMMUROBAHAHPAYREQ, konstanta.TABLEMEMMUROBAHAH)

func (DB *DBDriver) FindMemMurobahahPayReqTx(tx *sql.Tx) ([]models.PayMurobahahReq, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var reqs []models.PayMurobahahReq

	row, err := DB.DB.QueryContext(ctx, statementFindMemMurobahahPayReq)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var tmp models.PayMurobahahReq
		err = row.Scan(&tmp.ID, &tmp.Date, &tmp.MurobahahID, &tmp.Amount, &tmp.Info, &tmp.Doc)
		if err != nil {
			return nil, err
		}
		reqs = append(reqs, tmp)
	}

	return reqs, nil
}
