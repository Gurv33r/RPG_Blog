package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Gurv33r/RPG_Blog/backend/database"
	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
)

var err error
var db *pg.DB

func allPosts(w http.ResponseWriter, r *http.Request) {
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

func getPost(w http.ResponseWriter, r *http.Request) {
	// grab date from uri
	date := mux.Vars(r)["date"]
	// query db
	result := &database.Post{}
	db = database.NewConn() // establish connection
	// pass the query
	err := db.Model(result).
		Where("date = ?", date).
		Select()
	db.Close() //close the connection
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// send json response
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(result)
}

func newPost(w http.ResponseWriter, r *http.Request) {
	// decode request into Post struct
	var post database.Post
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// store post data into db
	db = database.NewConn()            //establish connection
	_, err := db.Model(&post).Insert() // pass query, will send back result, but ignore it
	db.Close()                         // close connection
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// send back acceptance code and jsonized post
	w.WriteHeader(200)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	date := mux.Vars(r)["date"]
	db = database.NewConn()
	_, err = db.Model(&database.Post{}).
		Where("date = ?", date).
		Delete()
	db.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
}

func editPost(w http.ResponseWriter, r *http.Request) {
	// grab new content
	var post database.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	post.UpdatedAt = time.Now()
	// grab date to edit
	date := mux.Vars(r)["date"]
	// update db
	db := database.NewConn() // establish connection
	// pass update query
	_, err = db.Model(&post).
		Column("content").
		Column("updated_at").
		Where("date = ?", date).
		Update()
	db.Close() // close connection
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// respond with status code 200
	w.WriteHeader(200)
}
