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
func (f *Feedback) Delete() {
	db := ConnectDB()
	db.Delete(f)
}

type Experience struct {
	ID          uint   `gorm:"primaryKey"`
	Period      int    `json:"period"` // the days i spent during works.
	Role        string `json:"role"`
	Company     string `json:"company"`
	Description string `json:"descritpion"`
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

func (p *Project) Delete() {
	db := ConnectDB()
	db.Delete(p)
}

func ConnectDB() *gorm.DB {
	//new server provisioned - postgre 172.17.0.2 / resumeapi 172.17.0.3 / fileserver 172.17.0.4
	//if you use separated pods, you need to activate a service controller with Type Clutser IP
	//the with Selector/metadata name. and the host should be host= name of service then it will query to registered Domain name in Kube DNS
	//dsn := "host=localhost, port=5432" // for kubernetes Pod deployment workernode -1
	dsn := "host=172.17.0.2 user=postgres password=root1234 dbname=resume1 port=5432 sslmode=disable TimeZone=Asia/Seoul" // for docker infra deployment.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error occrued with : ", err)
	}

	db.AutoMigrate(&Feedback{}, &Experience{}, &Skill{}, &Languages{}, &Project{})

	return db
}
