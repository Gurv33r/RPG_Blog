package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Gurv33r/RPG_Blog/backend/database"
	"github.com/Gurv33r/go-env"
	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
)

var err error
var db *pg.DB

func init() {
	log.Println("If you see me multiple times, I am redundant")
	env.LoadFrom("./env/db.env")
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
	var posts []database.Post       // make post slice to receive posts in
	db = database.NewConn()         // establish connection to db
	err = db.Model(&posts).Select() // pass query to access all of them
	db.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// send posts back as JSON
	w.WriteHeader(200)
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
	post.CreatedAt = time.Now()
	log.Println("Adding", post, "to DB!")
	// store post data into db
	db = database.NewConn()            //establish connection
	_, err := db.Model(&post).Insert() // pass query, will send back result, but ignore it
	db.Close()                         // close connection
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// send back acceptance code
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(post)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	date := mux.Vars(r)["date"]
	db = database.NewConn()
	_, err := db.Model(&database.Post{}).Where("date = ?", date).Delete()
	db.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
}
