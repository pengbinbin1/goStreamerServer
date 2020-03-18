package main

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CreaterUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "enter create user")
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
