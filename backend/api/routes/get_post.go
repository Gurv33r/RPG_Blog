package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Gurv33r/RPG_Blog/backend/database"
	"github.com/gorilla/mux"
)

func GetPost(w http.ResponseWriter, r *http.Request) {
	// record the request onto the server logs.
	err := record(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// grab date from uri
	date := mux.Vars(r)["date"]
	if !isValid(date) {
		http.Error(w, "Path not found", http.StatusNotFound)
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
