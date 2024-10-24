package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewHandlers() *http.ServeMux {
	mux := http.NewServeMux()
	//HandlerFunc의 역할...
	//serveHTTP를 반환하는 Handler로 바꿔줌 Handlefunc을
	//mux.Handle("/updateres", authMiddleware(http.HandlerFunc(UpdateResume)))
	mux.HandleFunc("/index", IndexHandler)
	//testmodel - authmiddelware - need to have new env for testing the code below.

	mux.HandleFunc("/editproject", Authmiddelware(Editproject))

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
	//mux.HandleFunc("/editproject", Editproject)
	mux.HandleFunc("/requestprojectedit", RequestProjectEdit)
	mux.HandleFunc("/projectuploadpage", Projectuploadpage)
	mux.HandleFunc("/deleteproject", DeleteProject)
	mux.HandleFunc("/imageurlsaverequest", ImageurlSaveRequest) // deprecated soon - since it could be adjusted with a CSS and Html.
	mux.HandleFunc("/uploadproject", UploadProject)

	//User and Login.
	mux.HandleFunc("/requestlogin", LoginRequest)
	mux.HandleFunc("/joinus", JoinasMember)
	mux.HandleFunc("/logout", Logout)

	//Oauth
	mux.HandleFunc("/oauthsignin", OauthSignin)
	mux.HandleFunc("/navercallback", OauthCallback)
	//Static Fileserver - Css / JS push
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/index", http.StatusMovedPermanently)
	})

	staticFileServer := http.FileServer(http.Dir("./home"))
	mux.Handle("/home/", http.StripPrefix("/home/", staticFileServer))
	// 뒤에 인덱스페이지 요청으로 가게끔 해야함... 첫페이지가 인덱스임.
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("handlerset")

	return mux
}
