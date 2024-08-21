package main

import (
	"fmt"
	"net/http"
)

func NewHandlers() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/download", PrintPDF)
	mux.HandleFunc("/resume", ReturnResume)
	mux.HandleFunc("/feedback", SendingFeedback)
	mux.HandleFunc("/index", Indexhandler)
	mux.HandleFunc("/project", ReturnProject)
	fmt.Println("handlerset")
	return mux
}
