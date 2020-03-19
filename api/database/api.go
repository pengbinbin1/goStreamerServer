package data

import (
	"database/sql"
	"goStreamerServer/api/defs"
	"goStreamerServer/utils"
	"log"
	"time"

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
		log.Println("exec faild:", err)
		return err
	}

	defer stmtInts.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOuts, err := dbConn.Prepare("SELECT pwd FROM users where user_name = ?")
	if err != nil {
		log.Println("err:", err)
		return "", err
	}
	var pwd string
	err = stmtOuts.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		log.Println("sacn failed:", err)
		return "", err
	}
	defer stmtOuts.Close()
	return pwd, nil
}
func DelUser(loginName string, pwd string) error {
	stmtDels, err := dbConn.Prepare("DELETE FROM users WHERE user_name = ? AND pwd = ? ")
	if err != nil {
		log.Println("delete from users failed:", err)
		return err
	}
	_, err = stmtDels.Exec(loginName, pwd)
	if err != nil {
		log.Println("exec faild:", err)
		return err
	}
	defer stmtDels.Close()
	return nil
}

func AddVideo(authID int, name string) (*defs.VideoInfo, error) {
	uuid := utils.NewUUID()
	t := time.Now()
	displaytime := t.Format("Jan 02 2006, 15:04:05")
	stmtInts, err := dbConn.Prepare("INSERT INTO video_info (id,author_id,name,display_ctime) VALUES(?,?,?,?)")

	if err != nil {
		log.Println("dbConn prepare failed:", err)
		return nil, err
	}
	_, err = stmtInts.Exec(uuid, authID, name, displaytime)
	if err != nil {
		log.Println("dbConn Exec failed:", err)
		return nil, err
	}
	defer stmtInts.Close()

	res := &defs.VideoInfo{ID: uuid, AuthorID: authID, Name: name, DisplayTime: displaytime}

	return res, nil
}

func DeleVideo(vid string) error {
	stmtDle, err := dbConn.Prepare("DELETE FROM video_info WHERE id = ?")
	if err != nil {
		log.Println("dbConn prepare failed:", err)
		return err
	}
	_, err = stmtDle.Exec(vid)
	if err != nil {
		log.Println("dbConn Exec failed:", err)
		return err
	}
	defer stmtDle.Close()
	return nil
}

func GetOneVideo(vid string) (*defs.VideoInfo, error) {
	stmtSelect, err := dbConn.Prepare("SELECT name ,author_id,display_ctime FROM video_info WHERE id = ?")
	if err != nil {
		log.Println("dbConn prepare failed:", err)
		return nil, err
	}
	var name string
	var displaytime string
	var authID int
	err = stmtSelect.QueryRow(vid).Scan(&name, &authID, &displaytime)
	if err != nil {
		log.Println("dbConn scan failed:", err)
		return nil, err
	}
	defer stmtSelect.Close()
	res := &defs.VideoInfo{ID: vid, AuthorID: authID, Name: name, DisplayTime: displaytime}
	return res, nil
}

func AddNewComments(vid string, authID int, content string) error {
	tempUUID := utils.NewUUID()

	stmtInts, err := dbConn.Prepare("INSERT INTO comments (id,video_id,author_id,content) VALUES(?,?,?,?) ")
	if err != nil {
		log.Println("dbConn prepare failed:", err)
		return err
	}

	_, err = stmtInts.Exec(tempUUID, vid, authID, content)
	if err != nil {
		log.Println("dbConn exec failed:", err)
		return err
	}
	defer stmtInts.Close()
	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(`SELECT comments.id,users.user_name,comments.content FROM comments
								INNER JOIN users ON comments.author_id = users.id
								WHERE comments.video_id =? AND comments.time>FROM_UNIXTIME(?) AND comments.time<=FROM_UNIXTIME(?)`)
	if err != nil {
		log.Println("dbConn prepare failed:", err)
		return nil, err
	}
	var res []*defs.Comment

	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		log.Println("query failed:", err)
		return res, err
	}

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			log.Println("scan failed:", err)
			return res, err
		}

		recorder := &defs.Comment{ID: id, VideoID: vid, AuthorName: name, Content: content}
		res = append(res, recorder)
	}
	defer stmtOut.Close()
	return res, nil
}
