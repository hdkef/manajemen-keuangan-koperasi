package driver

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DBDriver struct {
	DB *sql.DB
}

var DBUSER string
var DBPASS string
var DBNAME string
var DBPORT string

func init() {
	_ = godotenv.Load()
}

func DBConn() *DBDriver {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable",
		DBUSER, DBPASS, DBNAME, DBPORT)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	for {
		err := db.Ping()
		if err != nil {
			return &DBDriver{
				DB: db,
			}
		}
		fmt.Println("ping db...")
		time.Sleep(5000 * time.Millisecond)
	}
}
