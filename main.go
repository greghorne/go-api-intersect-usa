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
	// "strconv"
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


func xyValid(xlng string, ylat string) (valid bool) {

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

	valid = false

	// ===============
	db, err := sql.Open("postgres", strConnect)
	if err != nil {
		log.Fatal(err)
	}
	// ===============

	// ===============
	var strQuery = "select z_tl_2016_us_state($1, $2);"
	// var strQuery = "select z_tl_2016_us_state(48, 15);"  Yemen
	rows, err := db.Query(strQuery, xlng, ylat)
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}
	// ===============

	// ===============
	// counter:=0
	for rows.Next() {
		  valid = true
		  return
	}
	// fmt.Println(strconv.Itoa(counter))

	return	
}




func v1GetXY (w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	lng := params["lng"]
	lat := params["lat"]
	// fmt.Println(lng + ", " + lat)
	fmt.Println(xyValid(lng, lat))
	
	// pass data back to client as JSON
	// json.NewEncoder(w).Encode(info)

	// host     := os.Getenv("GO_HOST")
	// database := os.Getenv("GO_DATABASE")
	// user     := os.Getenv("GO_USER")
	// password := os.Getenv("GO_PASSWORD")
	// port     := os.Getenv("GO_PORT")

	// strConnect := "host=" + host + 
	// 			  " database=" + database + 
	// 			  " user=" + user + 
	// 			  " password=" + password + 
	// 			  " port=" + port + 
	// 			  " sslmode=require"

	// fmt.Println(strConnect)
	// conn := conn(strConnect)
	// fmt.Println(conn)

	// if (conn != nil) {
	// 	rows , err := conn.Query('select z_tl_2016_us_state($1, $2)', lng, lat)
		
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	return
	// } else {
	// 	fmt.Println("error: unable to connect to database")
	// }

	// db, err := sql.Open("postgres", strConnect)
	// fmt.Println(db)
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	row := db.Query('select z_tl_2016_us_state($1, $2)', -95, 36)
	// }
	// fmt.Println(conn)
	

}
