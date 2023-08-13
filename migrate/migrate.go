package main

import (
	"gorestapi/db"
	"gorestapi/model"
	"log"
)

func main() {
	dbConn := db.NewDB()
	defer db.Close(dbConn)

	dbConn.AutoMigrate(&model.User{}, &model.Task{})

	log.Println("Migrate successfully")
}
