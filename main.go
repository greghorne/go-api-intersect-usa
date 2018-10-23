package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"os"
	"database/sql"
	"encoding/json"
	// "fm:w
	t"
)

// ============================================================
var const_connect_string = 	" host="     + os.Getenv("GO_HOST") + 
	    					" database=" + os.Getenv("GO_DATABASE") + 
							" user="     + os.Getenv("GO_USER") + 
							" password=" + os.Getenv("GO_PASSWORD") + 
							" port="     + os.Getenv("GO_PORT") + 
	 						" sslmode=require"

var const_db_type = "postgres"
// ============================================================


// ============================================================
func main() {
	
	router := mux.NewRouter()
	router.HandleFunc("/v1/intersects-usa/{lng}/{lat}", v1IntersectsUSA).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))

}
// ============================================================


// ============================================================
func xyIntersectsUSA(xLng string, yLat string) (bIntersects bool) {

	// connect to db
	db, err := sql.Open(const_db_type, const_connect_string)
	
	defer db.Close()
	if err != nil { log.Fatal(err) }
	
	// exec pg funtion select z_tl_2016_us_state(lng, lat)
	var strQuery = "select z_tl_2016_us_state($1, $2);"
	rows, err   := db.Query(strQuery, xLng, yLat)
	
	defer rows.Close()
	if err != nil { log.Fatal(err) }

	bIntersects = false
	for rows.Next() {
		bIntersects = true
	} 

	return

}
// ============================================================


// ============================================================
func v1IntersectsUSA (w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	bIntersects := xyIntersectsUSA(params["lng"], params["lat"])
    jsonResult  := map[string]bool{"intersects": bIntersects}

	json.NewEncoder(w).Encode(jsonResult)

}
// ============================================================