package data

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:peng123456@tcp(localhost:3306)/videoserver")
	if err != nil {
		panic(err.Error())
	}

}
