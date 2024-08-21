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
	Period      int // the days i spent during works.
	Role        string
	Company     string
	Description string
	ResumeRefer uint
}

type Resume struct {
	gorm.Model
	Experiences []Experience `gorm:"foreignKey:ResumeRefer"`
	Skills      []string     `gorm:"type:text[]"`
	Languages   []string     `gorm:"type:text[]"`
}

func ConnectDB() *gorm.DB {

	dsn := "host=localhost user=postgres password=root1234 dbname=resume port=8801 sslmode=disable TimeZone=Asia/Seoul"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error occrued with : ", err)
	}

	return db
}
