package driver

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"manajemen-keuangan-koperasi/konstanta"
	"manajemen-keuangan-koperasi/models"
)

var statementInsertUser string = fmt.Sprintf("INSERT INTO %s (member_id, username, passwd, role, isagent) VALUES (?,?,?,?,?)", konstanta.TABLEALLUSER)

func (DB *DBDriver) InsertUserTx(tx *sql.Tx, MemID string, Username string, Passwd string, Role string, IsAgent string) (sql.Result, error) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return DB.DB.ExecContext(ctx, statementInsertUser, MemID, Username, Passwd, Role, IsAgent)
}

var statementFindOneUserByUsername string = fmt.Sprintf("SELECT id, member_id, username, passwd, role, isagent FROM %s WHERE Username=?", konstanta.TABLEALLUSER)

func (DB *DBDriver) FindOneUserByUsername(value string) (models.User, error) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	row := DB.DB.QueryRowContext(ctx, statementFindOneUserByUsername, value)

	tmpUsr := models.User{}

	err := row.Scan(&tmpUsr.ID, &tmpUsr.MemID, &tmpUsr.Username, &tmpUsr.Pass, &tmpUsr.Role, &tmpUsr.IsAgent)
	if err != nil {
		return models.User{}, err
	}

	return tmpUsr, nil
}

var statementCreateZeroBalance string = fmt.Sprintf("INSERT INTO %s (uid,IP,IW,SS,SHU,Bonus) VALUES (?,?,?,?,?,?)", konstanta.TABLEMEMBALANCE)

func (DB *DBDriver) CreateZeroBalance(tx *sql.Tx, uid float64) (sql.Result, error) {
	return tx.Exec(statementCreateZeroBalance, uid, 0, 0, 0, 0, 0)
}

var statementFindOneUserByUID string = fmt.Sprintf("SELECT id, member_id, username, passwd, role, isagent FROM %s WHERE id=?", konstanta.TABLEALLUSER)

func (DB *DBDriver) FindOneUserByUIDTx(tx *sql.Tx, value float64) (models.User, error) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	row := tx.QueryRowContext(ctx, statementFindOneUserByUID, value)

	tmpUsr := models.User{}

	err := row.Scan(&tmpUsr.ID, &tmpUsr.MemID, &tmpUsr.Username, &tmpUsr.Pass, &tmpUsr.Role, &tmpUsr.IsAgent)
	if err != nil {
		return models.User{}, err
	}

	return tmpUsr, nil
}

var statementFindAllUserByUsername = fmt.Sprintf("SELECT id, member_id, username, role, isagent FROM %s WHERE username=?", konstanta.TABLEALLUSER)
var statementFindAllUserByRole = fmt.Sprintf("SELECT id, member_id, username, role, isagent FROM %s WHERE role=?", konstanta.TABLEALLUSER)
var statementFindAllUserByMemID = fmt.Sprintf("SELECT id, member_id, username, role, isagent FROM %s WHERE member_id=?", konstanta.TABLEALLUSER)

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
	default:
		return nil, errors.New("no filter match")
	}

	if err != nil {
		return nil, err
	}

	var users []models.User

	for row.Next() {
		var tmpUsr models.User
		err = row.Scan(&tmpUsr.ID, &tmpUsr.MemID, &tmpUsr.Username, &tmpUsr.Role, &tmpUsr.IsAgent)
		if err != nil {
			return nil, err
		}
		users = append(users, tmpUsr)
	}

	return users, nil
}
