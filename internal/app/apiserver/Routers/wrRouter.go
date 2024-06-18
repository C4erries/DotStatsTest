package router

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func ConfigureStatsSubRouter(router *mux.Router) *mux.Router {
	subrouter := router.PathPrefix("/stats").Subrouter()
	subrouter.HandleFunc("/wr", handlePostWRStats()).Methods("POST", "GET")
	return subrouter
}

func handlePostWRStats() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, err := os.ReadFile("jsons/HeroesWr.json")
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(file)
		if err != nil {
			log.Fatal(err)
		}
	}
}
