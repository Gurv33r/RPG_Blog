package routes

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Gurv33r/RPG_Blog/backend/database"
	"github.com/gorilla/mux"
)

func EditPost(w http.ResponseWriter, r *http.Request) {
	// record the request onto the server logs.
	err := record(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// grab new content
	var post database.Post
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// add update time
	post.UpdatedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	// grab date to edit
	date := mux.Vars(r)["date"]

	// validate date
	if !isValid(date) {
		http.Error(w, "Path not found", http.StatusNotFound)
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
