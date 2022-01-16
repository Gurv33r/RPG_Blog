package routes

import (
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func isValid(input string) bool {
	const iso = "2006-01-02"
	_, err := time.Parse(iso, input)
	return err == nil
}

func record(r *http.Request) error {
	reqdump, err := httputil.DumpRequest(r, true)
	if err != nil {
		return err
	}
	log.Println(string(reqdump))
	return nil

}
