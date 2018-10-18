// https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo

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
	router.HandleFunc("/v1/getinfo/{xy}", V1GetXY).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func V1GetXY (w http.ResponseWriter, r *http.Request) {}