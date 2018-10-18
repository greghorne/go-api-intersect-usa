package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("main...")

	router := mux.NewRouter()
	router.HandleFunc("/v1/getinfo/{xy}", GetXY.Methods("GET"))
	log.Fatal(http.ListenAndServer(":8000", router))
}