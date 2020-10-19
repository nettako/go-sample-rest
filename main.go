package main

import (
	"fmt"
	"log"
	"net/http"

	"./service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", service.ReturnAllUsers).Methods("GET")
	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe("127.0.0.1:1234", router))
}
