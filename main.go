package main

import (
	"net/http"
	"zafkiel/package/controller"
	"zafkiel/package/model"

	"github.com/gorilla/context"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	//Setup DB
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	//Setup controller
	controller.Startup()

	http.ListenAndServe(":9090", context.ClearHandler(http.DefaultServeMux))
}
