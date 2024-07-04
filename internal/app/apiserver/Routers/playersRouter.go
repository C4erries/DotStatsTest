package router

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func ConfigurePlayersRouter(router *mux.Router) *mux.Router {
	subrouter := router.PathPrefix("/players").Subrouter()
	subrouter.HandleFunc("", handlePlayersList()).Methods("POST", "GET")
	return subrouter
}

func handlePlayersList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, err := os.ReadFile("jsons/Players.json")
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(file)
		if err != nil {
			log.Fatal(err)
		}
	}
}
