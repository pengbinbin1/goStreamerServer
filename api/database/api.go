package data

import (
	"database/sql"
	"fmt"
	"goStreamerServer/api/defs"
	"goStreamerServer/utils"

	_ "github.com/go-sql-driver/mysql"
)

func opneConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:peng123456@tcp(localhost:3306)/videoserver/charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func AddUserCredential(loginName string, pwd string) error {
	stmtInts, err := dbConn.Prepare("INSERT INTO users (user_name,pwd) VALUES (?,?)")
	if err != nil {
		return err
	}
	_, err = stmtInts.Exec(loginName, pwd)
	if err != nil {
		fmt.Println("exec faild:", err)
		return err
	}

	defer stmtInts.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOuts, err := dbConn.Prepare("SELECT pwd FROM users where user_name = ?")
	if err != nil {
		fmt.Println("err:", err)
		return "", err
	}
	var pwd string
	err = stmtOuts.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("sacn failed:", err)
		return "", err
	}
	defer stmtOuts.Close()
	return pwd, nil
}
func DelUser(loginName string, pwd string) error {
	stmtDels, err := dbConn.Prepare("DELETE FROM users WHERE user_name = ? AND pwd = ? ")
	if err != nil {
		fmt.Println("delete from users failed:", err)
		return err
	}
	_, err = stmtDels.Exec(loginName, pwd)
	if err != nil {
		fmt.Println("exec faild:", err)
		return err
	}
	defer stmtDels.Close()
	return nil
}

func AddVideo() (*defs.VideoInfo, error) {
	uuid := utils.NewUUID()
	fmt.Println("uuid:", uuid)
	return nil, nil
}
