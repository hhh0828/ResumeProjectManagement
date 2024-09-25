package main

import (
	"fmt"
	"net/http"
)

func NewHandlers() *http.ServeMux {
	mux := http.NewServeMux()
	//HandlerFunc의 역할...
	//serveHTTP를 반환하는 Handler로 바꿔줌 Handlefunc을
	//mux.Handle("/updateres", authMiddleware(http.HandlerFunc(UpdateResume)))
	//mux.HandleFunc("/", Indexhandler)
	mux.HandleFunc("/returnresume", Authmiddelware(Editproject))
	mux.HandleFunc("/download", PrintPDF)
	mux.HandleFunc("/returnresume", ReturnResume)
	mux.HandleFunc("/updateresume", UpdateResume)
	mux.HandleFunc("/submit", SendingFeedback)
	mux.HandleFunc("/projectspage", ProjectPage)
	mux.HandleFunc("/returnproject", ReturnProject)
	mux.HandleFunc("/resumepage", ResumePage)
	mux.HandleFunc("/returnskillang", Returnskillang)
	mux.HandleFunc("/contactpage", Contactpage)
	mux.HandleFunc("/uploadresume", UploadResumeExp)
	mux.HandleFunc("/returnprojectone", Returnprojectone)
	mux.HandleFunc("/editproject", Editproject)
	mux.HandleFunc("/requestprojectedit", RequestProjectEdit)
	mux.HandleFunc("/projectuploadpage", Projectuploadpage)
	mux.HandleFunc("/deleteproject", DeleteProject)
	mux.HandleFunc("/imageurlsaverequest", ImageurlSaveRequest)
	mux.HandleFunc("/uploadproject", UploadProject)
	mux.HandleFunc("/login", LoginRequest)
	mux.HandleFunc("/joinus", JoinasMember)
	//파일서버
	staticFileServer := http.FileServer(http.Dir("./home"))
	mux.Handle("/", http.StripPrefix("/", staticFileServer))

	fmt.Println("handlerset")

	return mux
}
