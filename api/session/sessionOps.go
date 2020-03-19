package session

import (
	data "goStreamerServer/api/database"
	"goStreamerServer/api/defs"
	"goStreamerServer/utils"
	"log"
	"sync"
	"time"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func LoadSessionFromDB() error {
	recorder, err := data.ListAllSessions()
	if err != nil {
		log.Println("get all session from db faild:", err)
		return err
	}
	recorder.Range(func(k, v interface{}) bool {
		res := v.(*defs.SimpleSession)
		sessionMap.Store(k, res)
		return true
	})
	return nil
}

func GenerateSession(userName string) string {
	uuidStr := utils.NewUUID()
	ts := time.Now().Nanosecond() / 1000000

	ttl_t := ts + 30*60*1000 //超时时间为30分钟)
	ttl := (int64)(ttl_t)
	element := &defs.SimpleSession{UserName: userName, TTL: ttl}
	sessionMap.Store(uuidStr, element)
	data.InsertSession(uuidStr, ttl, userName)
	return uuidStr
}
func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	data.DeletSession(sid)
}
func IsSessionExpired(sid string) (string, bool) {
	element, ok := sessionMap.Load(sid)
	if ok {
		ts := (int64)(time.Now().Nanosecond() / 1000000)
		ttl := element.(*defs.SimpleSession).TTL
		if ttl < ts {
			deleteExpiredSession(sid)
			return "", true
		} else {

			return element.(*defs.SimpleSession).UserName, false
		}
	} else {
		return "", true
	}

}
