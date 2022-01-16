package routes

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Gurv33r/RPG_Blog/backend/database"
)

func NewPost(w http.ResponseWriter, r *http.Request) {
	// record the request onto the server logs.
	err := record(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// decode request into Post struct
	var post database.Post
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// store post data into db
	db := database.NewConn()          //establish connection
	_, err = db.Model(&post).Insert() // pass query, will send back result, but ignore it
	db.Close()                        // close connection
	if err != nil {
		if strings.Contains(err.Error(), "ERROR #23505") { // duplicate entry case
			http.Error(w, "The request had a date already stored in the database.\nThis endpoint is for new blog posts only.\nUpdate a post with the /edit endpoint.", http.StatusBadRequest)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	// send back jsonized posts
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}
