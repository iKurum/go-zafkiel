package main

import (
	"log"
	"net/http"
	"os"
	"zafkiel/package/controller"
	"zafkiel/package/model"

	"github.com/gorilla/context"
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

	port := os.Getenv("PORT")
	log.Println("Running on port: ", port)
	http.ListenAndServe(":"+port, context.ClearHandler(http.DefaultServeMux))
}
