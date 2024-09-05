package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Feedback struct
type Feedback struct {
	gorm.Model //Index PK, CreatedAT, UpdatedAT, DeletedAT
	Name       string
	Email      string
	Message    string
}

func (f *Feedback) Upload() {
	db := ConnectDB()
	db.Create(f)
}

type Experience struct {
	ID          uint `gorm:"primaryKey"`
	Period      int  // the days i spent during works.
	Role        string
	Company     string
	Description string
	ResumeID    uint
}

type Experiences struct {
	Exps []Experience `json:"exps"`
}

type Resume struct {
	gorm.Model
	Experiences Experience `gorm:"foreignKey:ResumeID"`
	Skills      string     `gorm:"foreignKey:ResumeID"`
	Languages   string     `gorm:"foreignKey:ResumeID"`
}

type ResumeJ struct {
	Experiences []Experience `json:"Experiences"`
	Skills      []Skill      `json:"Skills"`
	Languages   []Languages  `json:"Languages"`
}

// need to change when the resume page going out to user so the data should have ID and Desc.
type Skill struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	ResumeID    uint
}

type Languages struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Proficiency string
	ResumeID    uint
}

type Project struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `'json:"name"`
	ShortDesc string `json:"shortdesc"`
	LongDesc  string `json:"longdesc"`
	ImgUrl    string `json:"imgurl"`
	DetailUrl string `json:"detailurl"`
}

func (p *Project) Upload() {
	db := ConnectDB()
	db.Create(p)
}

func ConnectDB() *gorm.DB {
	//if you use separated pods, you need to activate a service controller with Type Clutser IP
	//the with Selector/metadata name. and the host should be host= name of service then it will query to registered Domain name in Kube DNS
	//dsn := "host=localhost, port=5432" // for kubernetes Pod deployment workernode -1
	dsn := "host=172.17.0.4 user=postgres password=root1234 dbname=resume1 port=5432 sslmode=disable TimeZone=Asia/Seoul" // for docker infra deployment.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error occrued with : ", err)
	}

	db.AutoMigrate(&Feedback{}, &Experience{}, &Skill{}, &Languages{}, &Project{})

	return db
}
