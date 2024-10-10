package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	IsLogged string
	Con      string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	var data PageData
	cookie, err := r.Cookie("token")
	if err != nil || !ValidateToken(cookie.Value) {
		fmt.Println("error while getting cookiee")
		//cookie.Value = "error"
		data = PageData{
			IsLogged: "/loginpage",
			Con:      "Login",
		}
	}
	if err == nil && ValidateToken(cookie.Value) {
		fmt.Println("login state still validated")
		data = PageData{
			IsLogged: "/logout",
			Con:      "Logout",
		}
	}

	t, err := template.ParseFiles("./home/index.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}
	fmt.Println(data)
	err1 := t.Execute(w, data)
	if err1 != nil {
		log.Printf("Error executing template: %v", err1)
		http.Error(w, "Template execution error", http.StatusInternalServerError)
	}
}
func ResumePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./home/resume.html")

	///여기서부터

}

func ProjectPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./home/projects.html")

}

func Projectuploadpage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./home/projectuploadpage.html")
}

func Editproject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-store")
	http.ServeFile(w, r, "./home/editproject.html")

}

func Contactpage(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "./home/contact.html")
}

func PrintPDF(w http.ResponseWriter, r *http.Request) {
	//PDF Creator API required
}

func UpdateResume(w http.ResponseWriter, r *http.Request) {

	var updatedresume Resume
	err := json.NewDecoder(r.Body).Decode(&updatedresume)
	if err != nil {
		log.Fatal("error occured with Decoding the Resume with message :", err)
	}

	updatedresume.ID = 1
	db := ConnectDB()
	db.AutoMigrate(&Resume{}, &Experience{})

	var resume Resume
	db.First(&resume, updatedresume.ID)

	db.Model(&resume).Updates(updatedresume)
}

// ocne user clicked the upload button, the page will guide you the uploader page.
// and show the page that is filled with some data user provided.
// getting object from uploader-
// connect DB and Update all the data with
// overwrite all of them right after the user click the save button
func UploadPage(w http.ResponseWriter, r *http.Request) {
	db := ConnectDB()
	var savedexpdata []Experience

	var goingexpdata Experiences
	db.Find(&savedexpdata)
	goingexpdata.Exps = savedexpdata

	data, err := json.Marshal(goingexpdata)
	if err != nil {
		log.Fatal("fata error occured with", err)
	}
	//fill the page with the data that is made as writable.
	w.Write(data)

}

// getting object >> Exps, Skills, Langs
// the request body must have the Pkey for updating the database properly.
func UploadResumeExp(w http.ResponseWriter, r *http.Request) {
	//DB conn Update
	db := ConnectDB()
	var exp Experience

	json.NewDecoder(r.Body).Decode(&exp)

	//given data from FE

	db.Create(&exp)

	data, _ := json.Marshal("uploaded")
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	//the model i need to create are - Exps,
	//Comparing with previous data that user has requested.

}

func ReturnResume(w http.ResponseWriter, r *http.Request) {
	//DB conn required - Read, Create,  Update
	//Object
	db := ConnectDB()

	var experience []Experience
	//wrapping to EXP slice
	var Exps Experiences

	db.Find(&experience)
	Exps.Exps = experience

	sendingdata, err := json.Marshal(Exps)
	if err != nil {
		log.Fatal("error occured with DB EXP", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(sendingdata)

}

func RequestProjectEdit(w http.ResponseWriter, r *http.Request) {

	var project Project
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		log.Fatal("error occured with Decoding the Feedback data with message :", err)
	}
	db := ConnectDB()

	var modelproject Project

	db.First(&modelproject, project.ID)
	modelproject.Name = project.Name
	modelproject.ShortDesc = project.ShortDesc
	modelproject.LongDesc = project.LongDesc
	modelproject.DetailUrl = project.DetailUrl
	db.Save(&modelproject)
	fmt.Println("is this working ? ", modelproject.DetailUrl)

	thanks := project.Name + "에 대한 편집이 완료되었습니다."
	data, _ := json.Marshal(thanks)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func UploadProject(w http.ResponseWriter, r *http.Request) {

	var uploadedproject Project
	err := json.NewDecoder(r.Body).Decode(&uploadedproject)
	if err != nil {
		log.Fatal("error occured with user project upload Data decode", err)
	}

	Upload(&uploadedproject)
	fmt.Println("Project", uploadedproject.Name, "has been uploaded")
	fmt.Println(uploadedproject)

	data, err := json.Marshal(uploadedproject.Name + "has been uploaded Thank you sir!")
	if err != nil {
		log.Fatal("fatal error occured with your Data marshaling", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func ImageurlSaveRequest(w http.ResponseWriter, r *http.Request) {

	// Parse multipart form with max file size

	var project Project
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		log.Println("error occured", err)
	}

	Upload(&project)
	data, err := json.Marshal(project.Name + "has been uploaded Thank you sir!")
	if err != nil {
		log.Fatal("fatal error occured with your Data marshaling", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {

	var deprecateddata Project
	err := json.NewDecoder(r.Body).Decode(&deprecateddata)
	if err != nil {
		log.Fatal("error occured with user project upload Data decode", err)
	}

	db := ConnectDB()
	db.Delete(&Project{}, deprecateddata.ID)
	fmt.Println("Project", deprecateddata.ID, "has been deleted")
	//json 포맷으로바꾸공..
	dataforimage, _ := json.Marshal(deprecateddata)
	//바이트타입 버퍼로 만들어줌
	newrequest, _ := http.NewRequest("POST", "http://172.17.0.4:8770/ImageDelete", bytes.NewBuffer(dataforimage))
	if err != nil {
		log.Println("something wrong with creating http request", err)
	}
	newrequest.Header.Set("Content-Type", "application/json")
	clicon, _ := http.DefaultClient.Do(newrequest)
	fmt.Println(clicon.StatusCode)

	data, err := json.Marshal("your project has been removed")
	if err != nil {
		log.Fatal("fatal error occured with your Data marshaling", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func SendingFeedback(w http.ResponseWriter, r *http.Request) {
	//DB conn required - Create
	//Send the username if it works fine.
	//Object
	fmt.Println("working?")
	var feedback Feedback
	err := json.NewDecoder(r.Body).Decode(&feedback)
	if err != nil {
		log.Fatal("error occured with Decoding the Feedback data with message :", err)
	}
	fmt.Println(feedback)
	Upload(&feedback)

	data, _ := json.Marshal(feedback.Name)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

type Projectarray struct {
	Projects []Project `json:"projects"`
}

type ProjectPKid struct {
	Pkid uint `json:"projectId"`
}

func Returnprojectone(w http.ResponseWriter, r *http.Request) {

	var projectpkid ProjectPKid
	err := json.NewDecoder(r.Body).Decode(&projectpkid)
	if err != nil {
		log.Fatal("error occured with getting data from client ", err)
	}
	db := ConnectDB()
	fmt.Println(projectpkid)

	dbproject := new(Project)
	fmt.Println("current PK id is ", projectpkid.Pkid)
	db.First(&dbproject, projectpkid.Pkid)
	data, err := json.Marshal(dbproject)
	if err != nil {
		log.Fatal("error occured with Marshaling the Projectdata", err)
	}
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Content-Type", "application/json")

	w.Write(data)

}

func ReturnProject(w http.ResponseWriter, r *http.Request) {

	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		fmt.Println(forwarded, "this user has accessed to your server")
	}
	//fmt.Println("a user accessed to ReturnProject, the IP address is :", r.RemoteAddr)
	/*func() {
		not impletemented yet. need to write the logs that user accessed to your site.
		and see your project.
	}
	*/

	//DB conn required - Create, Read, Update
	var projects []Project

	var projectarr Projectarray
	db := ConnectDB()
	db.Find(&projects)

	projectarr.Projects = projects
	data, err := json.Marshal(projectarr)
	if err != nil {
		log.Fatal("error occured with Marshaling the Projectdata", err)
	}
	w.Header().Set("Content-Type", "application/json")

	w.Write(data)

}

type Skillang struct {
	Skillsname []string `json:"skills"`
	Langsname  []string `json:"languages"`
}

func Returnskillang(w http.ResponseWriter, r *http.Request) {
	var skills []Skill
	var langauges []Languages

	db := ConnectDB()
	db.Find(&skills)
	db.Find(&langauges)
	//언어이름모음
	var langsname []string
	//스킬이름모음
	var skillsname []string

	//스킬의 이름만
	for _, skill := range skills {
		skillsname = append(skillsname, skill.Name)
	}
	//언어의 이름만
	for _, lang := range langauges {
		langsname = append(langsname, lang.Name)
	}
	//가져온 스킬과 언어를 구조체로 묶어서 Json으로 보낼거임

	var skillang Skillang
	skillang.Langsname = langsname
	skillang.Skillsname = skillsname

	data, _ := json.Marshal(skillang)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
