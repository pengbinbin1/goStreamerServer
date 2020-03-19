package data

import (
	"database/sql"
	"goStreamerServer/api/defs"
	"log"
	"strconv"
	"sync"
)

func InsertSession(sessionId string, ttl int64, username string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("INSERT INTO sessions (session_id,TTL,user_name) VALUES(?,?,?)")
	if err != nil {
		log.Println("dbConn prepare failed:", err)
		return err
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(sessionId, ttlstr, username)
	if err != nil {
		log.Println("dbConn exec failed:", err)
		return err
	}

	return nil
}

func GetOneSession(sid string) (*defs.SimpleSession, error) {
	stmtOuts, err := dbConn.Prepare("SELECT TTL,user_name from sessions WHERE session_id = ?")
	if err != nil {
		log.Println("dbConn prepare failed:", err)
		return nil, err
	}
	var userName string
	var ttlstr string
	err = stmtOuts.QueryRow(sid).Scan(&userName, &ttlstr)
	if err != nil && err != sql.ErrNoRows {
		log.Println("dbConn scan failed:", err)
		return nil, err
	}
	var ttl int64
	ttl, err = strconv.ParseInt(ttlstr, 10, 64)
	if err != nil {
		log.Println("parseInt failed:", err)
	}

	res := &defs.SimpleSession{UserName: userName, TTL: ttl}
	defer stmtOuts.Close()
	return res, nil

}

func ListAllSessions() (*sync.Map, error) {

	var m *sync.Map
	stmtOuts, err := dbConn.Prepare("SELECT session_id,TTL,user_name from session")
	if err != nil {
		log.Println("dbConn prepare failed:", err)
		return nil, err
	}

	rows, err := stmtOuts.Query()
	if err != nil {
		log.Println("dbConn exec failed:", err)
		return nil, err
	}

	for rows.Next() {
		var session_id string
		var ttlstr string
		var user_name string

		err := rows.Scan(&session_id, &ttlstr, &user_name)
		if err != nil {
			log.Println("dbConn scan failed:", err)
			return nil, err
		}

		ttl, _ := strconv.ParseInt(ttlstr, 10, 64)
		element := &defs.SimpleSession{UserName: user_name, TTL: ttl}
		m.Store(session_id, element)
	}
	defer stmtOuts.Close()
	return m, nil
}

func DeletSession(sid string) error {
	stmtIns, err := dbConn.Prepare("DELETE FROM session WHERE session_id = ?")
	if err != nil {
		log.Println("dbConn prepare failed:", err)
		return err
	}

	_, err = stmtIns.Exec(sid)
	if err != nil {
		log.Println("dbConn exec failed:", err)
		return err
	}
	return nil
}
