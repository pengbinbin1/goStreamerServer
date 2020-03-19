package main

import (
	"io"
)
import (
	"net/http"
)

func sendErrMessage(w http.ResponseWriter, sc int, errMsg string) {
	w.WriteHeader(sc)
	io.WriteString(w, errMsg)
}
