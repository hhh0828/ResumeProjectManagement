package main

import (
	"encoding/json"
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

func SendingFeedback(w http.ResponseWriter, r *http.Request) {
	//DB conn required - Create
	//Object
	var feedback Feedback
	err := json.NewDecoder(r.Body).Decode(&feedback)
	if err != nil {
		log.Fatal("error occured with Decoding the Feedback data with message :", err)
	}
	db := ConnectDB()
	db.AutoMigrate(&Feedback{})
	db.Create(&feedback)
}

func ReturnProject(w http.ResponseWriter, r *http.Request) {
	//DB conn required - Create, Read, Update
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
