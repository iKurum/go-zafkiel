package controller

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/sessions"
)

var (
	pageController page
	templates      map[string]*template.Template
	sessionName    string
	flashName      string
	store          *sessions.CookieStore
	pageLimit      int
)

func init() {
	templates = populateTemplates()
	store = sessions.NewCookieStore([]byte("kurumi-zafkiel-very-nice"))
	sessionName = "zafkiel"
	flashName = "go-flash"
	pageLimit = 6
}

//Startup func
//registerRoutes
func Startup() {
	port := os.Getenv("PORT")
	if port != "" {
		log.Println("Running on port: ", port)
		http.ListenAndServe(":"+port, pageController.registerRoutes())
	} else {
		log.Println("Running on port: ", 9090)
		http.ListenAndServe(":9090", pageController.registerRoutes())
	}
}
