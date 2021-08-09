package driver

import (
	"context"
	"database/sql"
	"fmt"
	"manajemen-keuangan-koperasi/konstanta"
)

func initMember(DB *sql.DB) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tx, err := DB.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	err = createTableUser(tx)
	if err != nil {
		panic(err)
	}
	// err = createTableMemJournal(tx)
	// if err != nil {
	// 	panic(err)
	// }
	// err = createTableMemReq(tx)
	// if err != nil {
	// 	panic(err)
	// }
	// err = createTableMemDebt(tx)
	// if err != nil {
	// 	panic(err)
	// }
	// err = createTableMemBalance(tx)
	// if err != nil {
	// 	panic(err)
	// }
	// err = createTableMemBalanceHistory(tx)
	// if err != nil {
	// 	panic(err)
	// }
}

func createTableUser(tx *sql.Tx) error {
	statement := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (id int AUTO_INCREMENT, member_id VARCHAR(20) UNIQUE NOT NULL, username VARCHAR(25) UNIQUE NOT NULL, passwd VARCHAR(250) NOT NULL, role VARCHAR(10) NOT NULL, CONSTRAINT PK_alluser PRIMARY KEY (id))", konstanta.TABLEALLUSER)
	_, err := tx.Exec(statement)
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return err
	}
	fmt.Println("no err")
	return nil
}

func createTableMemJournal(tx *sql.Tx) error {
	statement := fmt.Sprintf("CREATE DATABASE %s", konstanta.TABLEMEMJOURNAL)
	_, err := tx.Exec(statement)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func createTableMemDebt(tx *sql.Tx) error {
	statement := fmt.Sprintf("CREATE DATABASE %s", konstanta.TABLEMEMDEBT)
	_, err := tx.Exec(statement)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func createTableMemBalance(tx *sql.Tx) error {
	statement := fmt.Sprintf("CREATE DATABASE %s", konstanta.TABLEMEMBALANCE)
	_, err := tx.Exec(statement)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func createTableMemReq(tx *sql.Tx) error {
	statement := fmt.Sprintf("CREATE DATABASE %s", konstanta.TABLEMEMREQ)
	_, err := tx.Exec(statement)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func createTableMemBalanceHistory(tx *sql.Tx) error {
	statement := fmt.Sprintf("CREATE DATABASE %s", konstanta.TABLEMEMBALANCEHISTORY)
	_, err := tx.Exec(statement)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func insertSuperAdmin(tx *sql.Tx) error {
	statement := fmt.Sprintf("CREATE DATABASE %s", konstanta.TABLEMEMBALANCEHISTORY)
	_, err := tx.Exec(statement)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
