package api

import "github.com/gorilla/mux"

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/all", allPosts).Methods("GET")
	router.HandleFunc("/{date}", getPost).Methods("GET")
	router.HandleFunc("/new", newPost).Methods("POST")
	router.HandleFunc("/edit/{date}", editPost).Methods("POST", "PUT")
	router.HandleFunc("/remove/{date}", deletePost).Methods("DELETE")
	return router
}
