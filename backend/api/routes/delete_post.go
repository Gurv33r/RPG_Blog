package routes

import (
	"net/http"

	"github.com/Gurv33r/RPG_Blog/backend/database"
	"github.com/gorilla/mux"
)

func DeletePost(w http.ResponseWriter, r *http.Request) {
	// record the request onto the server logs.
	err := record(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	date := mux.Vars(r)["date"]
	if !isValid(date) {
		http.Error(w, "Path not found", http.StatusNotFound)
		return
	}
	db := database.NewConn()
	_, err = db.Model(&database.Post{}).
		Where("date = ?", date).
		Delete()
	db.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
