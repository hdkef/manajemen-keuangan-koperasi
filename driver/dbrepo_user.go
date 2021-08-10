package driver

import (
	"context"
	"database/sql"
	"fmt"
	"manajemen-keuangan-koperasi/konstanta"
	"manajemen-keuangan-koperasi/models"
)

var statementInsertUser string = fmt.Sprintf("INSERT INTO %s (member_id, username, passwd, role) VALUES (?,?,?,?)", konstanta.TABLEALLUSER)

func (DB *DBDriver) InsertUser(MemID string, Username string, Passwd string, Role string) (sql.Result, error) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return DB.DB.ExecContext(ctx, statementInsertUser, MemID, Username, Passwd, Role)
}

var preparedStatementFindOneUser string = fmt.Sprintf("SELECT id, member_id, username, passwd, role FROM %s", konstanta.TABLEALLUSER)

func (DB *DBDriver) FindOneUser(bywhat string, value string) (models.User, error) {

	statement := fmt.Sprintf("%s WHERE %s=?", preparedStatementFindOneUser, bywhat)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	row := DB.DB.QueryRowContext(ctx, statement, value)

	tmpUsr := models.User{}

	err := row.Scan(&tmpUsr.ID, &tmpUsr.MemID, &tmpUsr.Username, &tmpUsr.Pass, &tmpUsr.Role)
	if err != nil {
		return models.User{}, err
	}

	return tmpUsr, nil
}
