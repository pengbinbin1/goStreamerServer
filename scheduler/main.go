package main

import (
	"goStreamerServer/scheduler/taskrunner"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type MiddleWare struct {
	r *httprouter.Router
}

func NewMiddleWare(r *httprouter.Router) http.Handler {
	mid := &MiddleWare{}
	mid.r = r
	return mid
}
func (md *MiddleWare) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	md.r.ServeHTTP(w, r)
}

func RegisterHandler() *httprouter.Router {

	r := httprouter.New()
	r.GET("/video/:vid", DeletVideoRecord)
	return r
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	go taskrunner.Start()

	r := RegisterHandler()

	log.Fatal(http.ListenAndServe(":8003", r))

}
