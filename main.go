// https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo

package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"os"
	"database/sql"
	"encoding/json"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/v1/intersects-usa/{lng}/{lat}", v1GetXY).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}


func xyValid(xlng string, ylat string) (nValid int) {

	host     := os.Getenv("GO_HOST")
	database := os.Getenv("GO_DATABASE")
	user     := os.Getenv("GO_USER")
	password := os.Getenv("GO_PASSWORD")
	port     := os.Getenv("GO_PORT")

	strConnect := "host="      + host + 
				  " database=" + database + 
				  " user="     + user + 
				  " password=" + password + 
				  " port="     + port + 
				  " sslmode=require"

	nValid = 0  // false

	// connect to pg
	db, err := sql.Open("postgres", strConnect)
	
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	
	// exec pg funtion select z_tl_2016_us_state(lng, lat)
	var strQuery = "select z_tl_2016_us_state($1, $2);"
	rows, err := db.Query(strQuery, xlng, ylat)
	
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		  nValid = 1  // has row, true
		  return
	}
	return
}


func v1GetXY (w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	lng := params["lng"]
	lat := params["lat"]

	nValid := xyValid(lng, lat)
	json.NewEncoder(w).Encode(nValid)

}
