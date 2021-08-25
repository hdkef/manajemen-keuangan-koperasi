package driver

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type DBDriver struct {
	DB *sql.DB
}

var DBUSER string
var DBPASS string
var DBNAME string
var DBPORT string
var DBHOST string

func init() {
	_ = godotenv.Load()
	DBUSER = os.Getenv("DBUSER")
	DBPASS = os.Getenv("DBPASS")
	DBNAME = os.Getenv("DBNAME")
	DBPORT = os.Getenv("DBPORT")
	DBHOST = os.Getenv("DBHOST")
}

func DBConn() *DBDriver {
	dbinfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		DBUSER, DBPASS, DBHOST, DBPORT, DBNAME)

	db, err := sql.Open("mysql", dbinfo)
	if err != nil {
		panic(err)
	}
	for {
		err := db.Ping()
		if err == nil {
			initTable(db)
			return &DBDriver{
				DB: db,
			}
		}
		fmt.Println("ping db...")
		time.Sleep(5000 * time.Millisecond)
	}
}
