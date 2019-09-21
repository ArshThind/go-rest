package main

import (
	"log"
	"net/http"
	"user-service/github.com.arshthind/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/users/", handlers.UserHandler{})
	log.Println(http.ListenAndServe(":2300", mux))
}
