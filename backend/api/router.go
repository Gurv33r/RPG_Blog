package api

import (
	"github.com/Gurv33r/RPG_Blog/backend/api/routes"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/all", routes.AllPosts).Methods("GET")
	router.HandleFunc("/{date}", routes.GetPost).Methods("GET")
	router.HandleFunc("/new", routes.NewPost).Methods("POST")
	router.HandleFunc("/edit/{date}", routes.EditPost).Methods("POST", "PUT")
	router.HandleFunc("/remove/{date}", routes.DeletePost).Methods("DELETE")
	return router
}
