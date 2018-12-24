package login

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

//Login 登陆逻辑
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "POST" {
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			fmt.Println("token:", token)
		} else {
			fmt.Println("No token")
		}
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
	}

	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))

	t, err := template.ParseFiles("./src/zafkiel/views/index.html")
	if err != nil {
		log.Fatal("ParseFiles err:", err)
	}
	log.Println(t.Execute(w, token))
}
