package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Gurv33r/RPG_Blog/backend/api"
	"github.com/Gurv33r/go-env"
)

func main() {
	fmt.Println("Booting up Server")
	env.LoadFrom("./env/server.env")
	router := api.Router()
	srv := &http.Server{
		Handler: router,
		Addr:    os.Getenv("ADDR"),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
