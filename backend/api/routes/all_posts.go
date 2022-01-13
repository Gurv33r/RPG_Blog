package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/Gurv33r/RPG_Blog/backend/database"
)

func AllPosts(w http.ResponseWriter, r *http.Request) {
	reqdump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	log.Println(string(reqdump))
	// search db for posts
	var posts []database.Post       // make post slice to receive posts in
	db := database.NewConn()        // establish connection to db
	err = db.Model(&posts).Select() // pass query to access all of them
	db.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// send posts back as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&posts)
}
