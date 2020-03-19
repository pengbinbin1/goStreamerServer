package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid")
	video, err := os.Open(VIDEO_DIRECTORY + vid)
	if err != nil {
		log.Println("open file failed:", err)
		sendErrMessage(w, http.StatusInternalServerError, "Internal server err")
		return
	}

	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)
	defer video.Close()
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_VIDEO_SIZE)
	err := r.ParseMultipartForm(MAX_VIDEO_SIZE)
	if err != nil {
		log.Println("parse multi form failed:", err)
		sendErrMessage(w, http.StatusBadRequest, "parse form failed")
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		log.Println("form file failed:", err)
		sendErrMessage(w, http.StatusInternalServerError, "form file failed")
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("read all  file failed:", err)
		sendErrMessage(w, http.StatusInternalServerError, "read all file failed")
		return
	}

	vid := p.ByName("vid")
	err = ioutil.WriteFile(VIDEO_DIRECTORY+vid, data, 0666)
	if err != nil {
		log.Println("write  file failed:", err)
		sendErrMessage(w, http.StatusInternalServerError, "write file failed")
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "upload success")
}
