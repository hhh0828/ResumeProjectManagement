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
	//testmodel - authmiddelware - need to have new env for testing the code below.

	//mux.HandleFunc("/returnresume", Authmiddelware(Editproject))
	mux.HandleFunc("/loginpage", LoginPage)
	mux.HandleFunc("/contactpage", Contactpage)
	mux.HandleFunc("/resumepage", ResumePage)
	mux.HandleFunc("/projectspage", ProjectPage)

	mux.HandleFunc("/download", PrintPDF)

	//Resume page
	mux.HandleFunc("/returnresume", ReturnResume)
	mux.HandleFunc("/updateresume", UpdateResume)
	mux.HandleFunc("/returnskillang", Returnskillang)
	mux.HandleFunc("/uploadresume", UploadResumeExp)

	//Feedback
	mux.HandleFunc("/submit", SendingFeedback)
	//Project
	mux.HandleFunc("/returnproject", ReturnProject)
	mux.HandleFunc("/returnprojectone", Returnprojectone)
	mux.HandleFunc("/editproject", Editproject)
	mux.HandleFunc("/requestprojectedit", RequestProjectEdit)
	mux.HandleFunc("/projectuploadpage", Projectuploadpage)
	mux.HandleFunc("/deleteproject", DeleteProject)
	mux.HandleFunc("/imageurlsaverequest", ImageurlSaveRequest) // deprecated soon - since it could be adjusted with a CSS and Html.
	mux.HandleFunc("/uploadproject", UploadProject)

	//User and Login.
	mux.HandleFunc("/requestlogin", LoginRequest)
	mux.HandleFunc("/joinus", JoinasMember)

	//Static Fileserver - Css / JS push
	staticFileServer := http.FileServer(http.Dir("./home"))
	mux.Handle("/", http.StripPrefix("/", staticFileServer))

	fmt.Println("handlerset")

	return mux
}
