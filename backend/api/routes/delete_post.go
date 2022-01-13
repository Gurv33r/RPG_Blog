package routes

import (
	"log"
	"net/http"

	"github.com/Gurv33r/RPG_Blog/backend/database"
	"github.com/gorilla/mux"
)

func DeletePost(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	date := mux.Vars(r)["date"]
	if !validate(date) {
		http.Error(w, "Date uri format is YYYY-MM-DD", http.StatusBadRequest)
		return
	}
	db := database.NewConn()
	_, err := db.Model(&database.Post{}).
		Where("date = ?", date).
		Delete()
	db.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
