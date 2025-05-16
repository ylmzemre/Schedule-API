package main

import (
	"github.com/ylmzemre/Schelude-API/config"
	"github.com/ylmzemre/Schelude-API/models"
)

func main() {
	db, err := config.GetDB()
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.Lessons{})
	if err != nil {
		panic(err)
	}
}
