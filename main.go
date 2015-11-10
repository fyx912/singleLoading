package main

import (
	"log"
	"net/http"
	"singleLoading/controller"
)

func main() {
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/logout", controller.Logout)
	go log.Fatal(http.ListenAndServe(":9999", nil))
}
