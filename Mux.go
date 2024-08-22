package main

import (
	"fmt"
	"net/http"
)

func NewHandlers() *http.ServeMux {
	mux := http.NewServeMux()
	//mux.HandleFunc("/", Indexhandler)
	mux.HandleFunc("/download", PrintPDF)
	mux.HandleFunc("/returnresume", ReturnResume)
	mux.HandleFunc("/updateresume", UpdateResume)
	mux.HandleFunc("/feedback", SendingFeedback)
	mux.HandleFunc("/project", ReturnProject)
	mux.HandleFunc("/resumepage", ResumePage)
	mux.HandleFunc("/returnskillang", Returnskillang)
	//파일서버
	staticFileServer := http.FileServer(http.Dir("./home"))
	mux.Handle("/", http.StripPrefix("/", staticFileServer))

	fmt.Println("handlerset")

	return mux
}
