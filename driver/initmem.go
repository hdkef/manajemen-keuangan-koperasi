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
	err = createTableMemReqMurobahah(tx)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	err = createTableMemMurobahah(tx)
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
	err = createTableAllInfo(tx)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	err = createTableAgentHistory(tx)
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
	statement := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (id int AUTO_INCREMENT, member_id VARCHAR(20) UNIQUE NOT NULL, username VARCHAR(25) UNIQUE NOT NULL, passwd VARCHAR(250) NOT NULL, role ENUM('Admin-Input','Admin-Super','member') DEFAULT 'member' NOT NULL, isagent ENUM('Y','N') DEFAULT 'N' NOT NULL, TEL VARCHAR(16),PRIMARY KEY (id))", konstanta.TABLEALLUSER)
	_, err := tx.Exec(statement)
	if err != nil {

		fmt.Println(err)
		return err
	}
	return nil
}

func createTableMemJournal(tx *sql.Tx) error {
	statement := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id int AUTO_INCREMENT, uid int NOT NULL, date DATE DEFAULT (CURRENT_DATE) NOT NULL, type ENUM('IP','IW','SS+','SS-','SHU','Bonus','MRBH+','MRBH-') NOT NULL, amount FLOAT(14,2) UNSIGNED NOT NULL, info VARCHAR(250), approvedby int NOT NULL, PRIMARY KEY (id), FOREIGN KEY (uid) REFERENCES %s (id), FOREIGN KEY (approvedby) REFERENCES %s (id))", konstanta.TABLEMEMJOURNAL, konstanta.TABLEALLUSER, konstanta.TABLEALLUSER)
	_, err := tx.Exec(statement)
	if err != nil {

		return err
	}
	return nil
}

func createTableMemMurobahah(tx *sql.Tx) error {
	statement := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id int AUTO_INCREMENT, uid int NOT NULL ,date DATE DEFAULT (CURRENT_DATE) NOT NULL, due_date DATE NOT NULL, initial FLOAT(14,2) UNSIGNED NOT NULL, paid FLOAT(14,2) UNSIGNED NOT NULL, document VARCHAR(150) NOT NULL, info VARCHAR(250), approvedby int NOT NULL, PRIMARY KEY (id), FOREIGN KEY (uid) REFERENCES %s (id) , FOREIGN KEY (approvedby) REFERENCES %s (id) )", konstanta.TABLEMEMMUROBAHAH, konstanta.TABLEALLUSER, konstanta.TABLEALLUSER)
	_, err := tx.Exec(statement)
	if err != nil {

		return err
	}
	return nil
}

func createTableMemReqMurobahah(tx *sql.Tx) error {
	statement := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id int AUTO_INCREMENT, agent_id int NOT NULL, buyer_id int NOT NULL ,date DATE DEFAULT (CURRENT_DATE) NOT NULL, due_date DATE NOT NULL, amount FLOAT(14,2) UNSIGNED NOT NULL, document VARCHAR(150) NOT NULL, info VARCHAR(250), PRIMARY KEY (id), FOREIGN KEY (agent_id) REFERENCES %s (id), FOREIGN KEY (buyer_id) REFERENCES %s (id) )", konstanta.TABLEMEMREQMUROBAHAH, konstanta.TABLEALLUSER, konstanta.TABLEALLUSER)
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

func createTableAllInfo(tx *sql.Tx) error {
	statement := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (id int AUTO_INCREMENT, uid int NOT NULL, date DATE DEFAULT(CURRENT_DATE) NOT NULL, Info VARCHAR(100) NOT NULL, PRIMARY KEY (id), FOREIGN KEY (uid) REFERENCES %s(id))", konstanta.TABLEALLINFO, konstanta.TABLEALLUSER)
	_, err := tx.Exec(statement)
	if err != nil {

		return err
	}
	return nil
}

func createTableAgentHistory(tx *sql.Tx) error {
	statement := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (id int AUTO_INCREMENT, uid int NOT NULL, murobahah_id int NOT NULL, PRIMARY KEY (id), FOREIGN KEY (uid) REFERENCES %s(id), FOREIGN KEY (murobahah_id) REFERENCES %s(id))", konstanta.TABLEAGENTHISTORY, konstanta.TABLEALLUSER, konstanta.TABLEMEMMUROBAHAH)
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
	statement1 := fmt.Sprintf("INSERT INTO %s (member_id,username,passwd,role,isagent,tel) VALUES (?,?,?,?,?,?)", konstanta.TABLEALLUSER)
	res, err := tx.Exec(statement1, "A0", os.Getenv("SUPERADMIN"), string(hashedPassbyte), "Admin-Super", "Y", "000000000000")
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
