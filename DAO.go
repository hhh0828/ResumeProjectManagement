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

type Experience struct {
	ID          uint
	Period      int // the days i spent during works.
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
	Experiences []Experience `gorm:"foreignKey:ResumeID"`
	Skills      []string     `gorm:"type:text[]"`
	Languages   []string     `gorm:"type:text[]"`
}

type Skill struct {
	ID          uint
	Name        string
	Description string
}

type Languages struct {
	ID          uint
	Name        string
	Proficiency string
}

type Project struct {
	ID        uint
	Name      string `'json:"name"`
	ShortDesc string `json:"shortdesc"`
	LongDesc  string `json:"longdesc"`
	ImgUrl    string `json:"imgurl"`
	DetailUrl string `json:"detailurl"`
}

func ConnectDB() *gorm.DB {

	dsn := "host=172.17.0.2 user=postgres password=root1234 dbname=resume1 port=5432 sslmode=disable TimeZone=Asia/Seoul"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error occrued with : ", err)
	}
	db.AutoMigrate(&Feedback{}, &Resume{}, &Experience{}, &Skill{}, &Languages{}, &Project{})

	return db
}
