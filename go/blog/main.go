package main

import (
	"blog/router"
	"log"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	//路由
	router.Route()

	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
