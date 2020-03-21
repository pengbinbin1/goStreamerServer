package data

import (
	"log"
)

func ReadVideo(cnt int) ([]string, error) {
	stmtOuts, err := dbConn.Prepare("SELECT vid FROM video_del LIMIT ?")
	if err != nil {
		log.Println("dnConn prepare failed:", err)
		return nil, err
	}

	var res []string
	rows, err := stmtOuts.Query(cnt)
	if err != nil {
		log.Println("dnConn query failed:", err)
		return nil, err
	}

	for rows.Next() {
		var vid string
		err := rows.Scan(&vid)
		if err != nil {
			log.Println("rows scan failed:", err)
			return res, err
		}
		res = append(res, vid)
	}
	defer stmtOuts.Close()
	return res, nil
}

func DelVideo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_del WHERE vid = ?")
	if err != nil {
		log.Println("dnConn prepare failed:", err)
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		log.Println("dnConn exec failed:", err)
		return err
	}
	defer stmtDel.Close()
	return nil
}

func InsertDelVideo(vid string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO video_del (vid) VALUES (?)")
	if err != nil {
		log.Println("dnConn prepare failed:", err)
		return err
	}

	_, err = stmtIns.Exec(vid)
	if err != nil {
		log.Println("dnConn exec failed:", err)
		return err
	}
	return err
}
