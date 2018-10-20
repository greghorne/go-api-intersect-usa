// https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo

package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	// "encoding/json"
	_ "github.com/lib/pq"
	"os"
	"database/sql"
	"strconv"
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


func conn() {

	host     := os.Getenv("GO_HOST")
	database := os.Getenv("GO_DATABASE")
	user     := os.Getenv("GO_USER")
	password := os.Getenv("GO_PASSWORD")
	port     := os.Getenv("GO_PORT")

	strConnect := "host=" + host + " database=" + database + " user=" + user + " password=" + password + " port=" + port + " sslmode=require"
	fmt.Println(strConnect)
	return
	// db, err := sql.Open("postgres", strConnect)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// return db
	
}

// func doQuery(db, lng, lat) {
// 	row, err := db.Query('select z_tl_2016_us_state($1, $2)', lng, lat)
// 	return row
// }


func v1GetXY (w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	lng := strconv.ParseFloat(params["lng"], 64)
	lat := strconv.ParseFloat(params["lat"], 64)
	fmt.Println(lng + ", " + lat)
	
	// pass data back to client as JSON
	// json.NewEncoder(w).Encode(info)

	host     := os.Getenv("GO_HOST")
	database := os.Getenv("GO_DATABASE")
	user     := os.Getenv("GO_USER")
	password := os.Getenv("GO_PASSWORD")
	port     := os.Getenv("GO_PORT")

	strConnect := "host=" + host + " database=" + database + " user=" + user + " password=" + password + " port=" + port + " sslmode=require"
	fmt.Println(strConnect)

	db, err := sql.Open("postgres", strConnect)
	fmt.Println(db)
	if err != nil {
		log.Fatal(err)
	} else {
		row, err := db.Query('select z_tl_2016_us_state($1, $2)', lng, lat)
	}
	// fmt.Println(conn)
	

}
