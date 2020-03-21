package main

import (
	"goStreamerServer/scheduler/data"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func DeletVideoRecord(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid")
	if len(vid) == 0 {
		log.Println("param vid is invalid")
		sendResponse(w, http.StatusBadRequest, "invalid param")
		return
	}
	err := data.InsertDelVideo(vid)
	if err != nil {
		log.Println("insert into db failed:", err)
		sendResponse(w, http.StatusInternalServerError, "internal error")
		return
	}
	sendResponse(w, 200, "success")
}
