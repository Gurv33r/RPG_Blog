package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Gurv33r/RPG_Blog/backend/database"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var err error
var db *gorm.DB

func init() {
	db = database.NewConn()
	db.AutoMigrate(&database.Post{})
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", getPosts).Methods("GET")
	router.HandleFunc("/new", newPost).Methods("POST")
	//router.HandleFunc("/edit/{date}", editPost).Methods("POST","PUT")
	router.HandleFunc("/remove/{date}", deletePost).Methods("DELETE")
	return router
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	// search db for posts
	var posts []database.Post // make post slice to receive posts in
	result := db.Find(&posts) // pass query to access all of them
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	// send posts back as JSON
	json.NewEncoder(w).Encode(&posts)
}

func newPost(w http.ResponseWriter, r *http.Request) {
	// decode request into Post struct
	var post database.Post
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("Adding", post, "to DB!")
	// store post data into db
	db.Create(&post)
	// send back acceptance code
	w.WriteHeader(200)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	date := mux.Vars(r)["date"]
	db.Delete(&database.Post{}, date)
	w.WriteHeader(200)
}
