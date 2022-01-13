package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/Gurv33r/RPG_Blog/backend/database"
	"github.com/gorilla/mux"
)

func EditPost(w http.ResponseWriter, r *http.Request) {
	reqdump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	log.Println(string(reqdump))
	// grab new content
	var post database.Post
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// grab date to edit
	date := mux.Vars(r)["date"]

	// validate date
	if !validate(date) {
		http.Error(w, "Date uri format is YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	// update db
	db := database.NewConn() // establish connection
	// pass update query
	_, err = db.Model(&post).
		Column("content").
		Where("date = ?", date).
		Update()
	db.Close() // close connection
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
