package main

import (
	"encoding/json"
	"goStreamerServer/api/defs"
	"io"
	"net/http"
)

func SendErrResponse(w http.ResponseWriter, errResp defs.ErrResponse) {
	w.WriteHeader(errResp.HttpErr)
	resStr, _ := json.Marshal(&errResp.InnerErr)
	io.WriteString(w, string(resStr))
}
func SendNormalResponse(w http.ResponseWriter, sc int, resp string) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
