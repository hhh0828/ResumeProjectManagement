package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func Indexhandler(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "./home/index.html")

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
