package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type MiddleWareHandler struct {
	r *httprouter.Router
	l *ConnLimter
}

func NewMiddleWareHanlder(n int, r *httprouter.Router) *MiddleWareHandler {
	mh := &MiddleWareHandler{}
	mh.r = r
	mh.l = NewConnLimeter(n)
	return mh
}

func (mh *MiddleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !mh.l.GetConn() {
		sendErrMessage(w, http.StatusTooManyRequests, "too many connections")
		return
	}
	mh.r.ServeHTTP(w, r)
	defer mh.l.ReleaseConn()

}
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/video/:vid", streamHandler)
	router.POST("/video/:vid", uploadHandler)
	return router
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	r := RegisterHandlers()
	mh := NewMiddleWareHanlder(3, r)
	http.ListenAndServe(":8001", mh)
}
