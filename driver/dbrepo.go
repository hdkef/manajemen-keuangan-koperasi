package driver

import (
	"context"
	"database/sql"
	"fmt"
	"manajemen-keuangan-koperasi/konstanta"
)

func (DB *DBDriver) InsertUser(MemID string, Username string, Passwd string, Role string) (sql.Result, error) {

	statement := fmt.Sprintf("INSERT INTO %s (member_id, username, passwd, role) VALUES (?,?,?,?)", konstanta.TABLEALLUSER)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return DB.DB.ExecContext(ctx, statement, MemID, Username, Passwd, Role)
}
