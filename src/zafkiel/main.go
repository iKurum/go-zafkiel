package main

import (
	"log"
	"net/http"

	"zafkiel/package/login"
)

func main() {
	http.HandleFunc("/", login.Login)
	//启动静态服务
	http.Handle("/asset/", http.StripPrefix("/asset/", http.FileServer(http.Dir("./src/zafkiel/asset"))))

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe err:", err)
	}
}
