package driver

import (
	"context"
	"database/sql"
	"fmt"
	"manajemen-keuangan-koperasi/konstanta"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func initMember(DB *sql.DB) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tx, err := DB.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	err = createTableUser(tx)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	err = createTableMemJournal(tx)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	err = createTableMemReq(tx)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	err = createTableMemDebt(tx)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	err = createTableMemBalance(tx)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	err = createTableMemBalanceHistory(tx)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	err = insertSuperAdmin(tx)
	if err != nil {
		tx.Rollback()
		//member_id is unique, if already exist will return error
		fmt.Println(err)
	}
}

func createTableUser(tx *sql.Tx) error {
	statement := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (id int AUTO_INCREMENT, member_id VARCHAR(20) UNIQUE NOT NULL, username VARCHAR(25) UNIQUE NOT NULL, passwd VARCHAR(250) NOT NULL, role ENUM('Admin-Input','Admin-Super','member') DEFAULT 'member' NOT NULL, PRIMARY KEY (id))", konstanta.TABLEALLUSER)
	_, err := tx.Exec(statement)
	if err != nil {

		fmt.Println(err)
		return err
	}
	return nil
}

func createTableMemJournal(tx *sql.Tx) error {
	statement := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id int AUTO_INCREMENT, uid int NOT NULL, date DATE DEFAULT (CURRENT_DATE) NOT NULL, type ENUM('IP','IW','SS+','SS-','SHU','Bonus','D+','D-') NOT NULL, amount FLOAT(14,2) UNSIGNED NOT NULL, info VARCHAR(250), approvedby int NOT NULL, PRIMARY KEY (id), FOREIGN KEY (uid) REFERENCES %s (id), FOREIGN KEY (approvedby) REFERENCES %s (id))", konstanta.TABLEMEMJOURNAL, konstanta.TABLEALLUSER, konstanta.TABLEALLUSER)
	_, err := tx.Exec(statement)
	if err != nil {

		return err
	}
	return nil
}

func createTableMemDebt(tx *sql.Tx) error {
	statement := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id int AUTO_INCREMENT, uid int NOT NULL ,date DATE DEFAULT (CURRENT_DATE) NOT NULL, due_date DATE NOT NULL, initial FLOAT(14,2) UNSIGNED NOT NULL, paid FLOAT(14,2) UNSIGNED NOT NULL, document VARCHAR(50) NOT NULL, info VARCHAR(250), approvedby int NOT NULL, PRIMARY KEY (id), FOREIGN KEY (uid) REFERENCES %s (id) , FOREIGN KEY (approvedby) REFERENCES %s (id) )", konstanta.TABLEMEMDEBT, konstanta.TABLEALLUSER, konstanta.TABLEALLUSER)
	_, err := tx.Exec(statement)
	if err != nil {

		return err
	}
	return nil
}

func createTableMemBalance(tx *sql.Tx) error {
	statement := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id int AUTO_INCREMENT, uid int NOT NULL, IP FLOAT(14,2) UNSIGNED NOT NULL, IW FLOAT(14,2) UNSIGNED NOT NULL, SS FLOAT(14,2) UNSIGNED NOT NULL, SHU FLOAT(14,2) UNSIGNED NOT NULL, Bonus FLOAT(14,2) UNSIGNED NOT NULL, PRIMARY KEY (id), FOREIGN KEY (uid) REFERENCES %s (id) )", konstanta.TABLEMEMBALANCE, konstanta.TABLEALLUSER)
	_, err := tx.Exec(statement)
	if err != nil {

		return err
	}
	return nil
}

func createTableMemReq(tx *sql.Tx) error {
	statement := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id int AUTO_INCREMENT, uid int NOT NULL, date DATE DEFAULT (CURRENT_DATE) NOT NULL, type ENUM('IP','IW','SS+','SS-') NOT NULL, amount FLOAT(14,2) UNSIGNED NOT NULL, info VARCHAR(250), PRIMARY KEY (id), FOREIGN KEY (uid) REFERENCES %s (id))", konstanta.TABLEMEMREQ, konstanta.TABLEALLUSER)
	_, err := tx.Exec(statement)
	if err != nil {

		return err
	}
	return nil
}

func createTableMemBalanceHistory(tx *sql.Tx) error {
	statement := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id int AUTO_INCREMENT, uid int NOT NULL, date DATE DEFAULT(CURRENT_DATE) NOT NULL, IP FLOAT(14,2) UNSIGNED NOT NULL, IW FLOAT(14,2) UNSIGNED NOT NULL, SS FLOAT(14,2) UNSIGNED NOT NULL, SHU FLOAT(14,2) UNSIGNED NOT NULL, Bonus FLOAT(14,2) UNSIGNED NOT NULL, PRIMARY KEY (id), FOREIGN KEY (uid) REFERENCES %s (id))", konstanta.TABLEMEMBALANCEHISTORY, konstanta.TABLEALLUSER)
	_, err := tx.Exec(statement)
	if err != nil {

		return err
	}
	return nil
}

func insertSuperAdmin(tx *sql.Tx) error {
	pass := os.Getenv("SUPERPASS")
	//hashing pass before insert into db
	hashedPassbyte, err := bcrypt.GenerateFromPassword([]byte(pass), 5)
	if err != nil {
		return err
	}
	statement1 := fmt.Sprintf("INSERT INTO %s (member_id,username,passwd,role) VALUES (?,?,?,?)", konstanta.TABLEALLUSER)
	res, err := tx.Exec(statement1, "A0", os.Getenv("SUPERADMIN"), string(hashedPassbyte), "Admin-Super")
	if err != nil {

		return err
	}
	id, err := res.LastInsertId()
	if err != nil {

		return err
	}
	statement2 := fmt.Sprintf("INSERT INTO %s (uid,IP,IW,SS,SHU,Bonus) VALUES (?,?,?,?,?,?)", konstanta.TABLEMEMBALANCE)
	_, err = tx.Exec(statement2, id, 0, 0, 0, 0, 0)
	if err != nil {

		return err
	}
	return nil
}
