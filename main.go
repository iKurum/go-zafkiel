package main

import (
	"go-zafkiel/package/controller"
	"go-zafkiel/package/model"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	//Setup DB
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	//Setup controller
	controller.Startup()
}
