// https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo

package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

func main() {
	fmt.Println("main...")

	router := mux.NewRouter()
	router.HandleFunc("/v1/getinfo/{lng}/{lat}", v1GetXY).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}


type XYInfo struct {
	Country	string
	Name1	string
	Name2	string
}


func v1GetXY (w http.ResponseWriter, r *http.Request) {

	fmt.Println("=====================")
	params := mux.Vars(r)
	fmt.Println(params)
	fmt.Println(params["lng"])
	fmt.Println(params["lat"])
	fmt.Println("=====================")

	// create some data
	var info XYInfo
	info.Country = "USA"
	info.Name1	 = "Oklahoma"
	info.Name2   = "Tulsa"
	
	// pass data back to client as JSON
	json.NewEncoder(w).Encode(info)
	

}
