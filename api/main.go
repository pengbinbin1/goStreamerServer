package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.r.ServeHTTP(w, r)
}
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreaterUser)
	router.GET("/user/:username", Login)
	router.GET("/user/:username/videos", ListVideo)
	router.GET("/user/:username/videos/:vid", GetOneVideo)
	router.DELETE("/user/:username/videos/:vid", DelOneVideo)

	return router
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	fmt.Println("enter main")
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)

	http.ListenAndServe(":8000", mh)
}
