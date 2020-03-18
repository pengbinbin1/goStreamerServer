package main

import (
	"fmt"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreaterUser)
	router.GET("/user/:user_name", Login)
	router.GET("/user/:username/videos", ListVideo)
	router.GET("/user/:username/videos/:vid", GetOneVideo)
	router.DELETE("/user/:username/videos/:vid", DelOneVideo)

	return router
}

func main() {
	fmt.Println("enter main")
	r := RegisterHandlers()
	http.ListenAndServe(":8000", r)
}
