package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type PgInfo struct {
	Hostname string
	Database string
	Username string
	Password string
	Port     string
}

var DB *gorm.DB

func (i *PgInfo) ConnectDatabase() {
	dsn := "host=" + i.Hostname + " user=" + i.Username + " password=" + i.Password + " dbname=" + i.Database + " port=" + i.Port + " sslmode=disable TimeZone=Asia/Jakarta"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect to database PortgreSQL")
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database PortgreSQL")
	}

	err = database.AutoMigrate()
	if err != nil {
		return
	}

	DB = database
}
