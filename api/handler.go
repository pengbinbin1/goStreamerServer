package main

import (
	"encoding/json"
	data "goStreamerServer/api/database"
	"goStreamerServer/api/defs"
	"goStreamerServer/api/session"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CreaterUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	req, _ := ioutil.ReadAll(r.Body)
	user := &defs.UserCredential{}
	err := json.Unmarshal(req, user)
	if err != nil {
		log.Println("json unmarshal failed:", err)
		SendErrResponse(w, defs.ErrParseFailed)
	}

	err = data.AddUserCredential(user.UserName, user.Pwd)
	if err != nil {
		log.Println("add user failed:", err)
		SendErrResponse(w, defs.ErrDBErr)
	}

	sessionId := session.GenerateSession(user.UserName)

	signedUp := defs.SingUp{Success: true, SessionID: sessionId}

	su, err := json.Marshal(signedUp)
	if err != nil {
		log.Println("json marashal failed:", err)
		SendErrResponse(w, defs.ErrInnerFailed)
	}
	SendNormalResponse(w, 200, string(su))

}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user := p.ByName("user_name")
	io.WriteString(w, user)
}

func ListVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func GetOneVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
func DelOneVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
