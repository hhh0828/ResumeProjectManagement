package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Indexhandler(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "./home/index.html")

}

func ResumePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./home/resume.html")

	///여기서부터

}

func ProjectPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./home/projects.html")

}

func Editproject(w http.ResponseWriter, r *http.Request) {
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
func UploadResume(w http.ResponseWriter, r *http.Request) {
	//DB conn Update
	db := ConnectDB()
	var resj ResumeJ
	var exps Experiences

	json.NewDecoder(r.Body).Decode(&resj)

	//given data from FE
	exps.Exps = resj.Experiences

	//the model i need to create are - Exps,
	var dbmodelexps Experiences
	var skillangexps Skillang
	db.Model(&dbmodelexps).Updates(exps.Exps)
	db.Model(&skillangexps)
	//Comparing with previous data that user has requested.

}

func ReturnResume(w http.ResponseWriter, r *http.Request) {
	//DB conn required - Read, Create,  Update
	//Object
	db := ConnectDB()

	var Exps Experiences

	for i := 1; i < 4; i++ {
		origind := new(Experience) //trash --
		db.First(&origind, i)
		Exps.Exps = append(Exps.Exps, *origind)
	}

	sendingdata, err := json.Marshal(Exps)

	if err != nil {
		log.Fatal("yes")
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
	db.Save(&modelproject)

	thanks := project.Name + "에 대한 편집이 완료되었습니다."
	data, _ := json.Marshal(thanks)
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
	db := ConnectDB()
	db.AutoMigrate(&Feedback{})
	db.Create(&feedback)

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
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func ReturnProject(w http.ResponseWriter, r *http.Request) {

	fmt.Println("logged well")
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