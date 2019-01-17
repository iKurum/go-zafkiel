package controller

import (
	"net/http"
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
func Startup() {
	//启动静态服务
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	pageController.registerRoutes()
}
