package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/Gurv33r/RPG_Blog/backend/database"
	"github.com/gorilla/mux"
)

func GetPost(w http.ResponseWriter, r *http.Request) {
	reqdump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	log.Println(string(reqdump))
	// grab date from uri
	date := mux.Vars(r)["date"]
	if !validate(date) {
		http.Error(w, "Date uri format is YYYY-MM-DD", http.StatusBadRequest)
		return
	}
	// query db
	result := &database.Post{}
	db := database.NewConn() // establish connection
	// pass the query
	err = db.Model(result).
		Where("date = ?", date).
		Select()
	db.Close() //close the connection
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// send json response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
