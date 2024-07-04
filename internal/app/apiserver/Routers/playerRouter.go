package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ConfigurePlayerProfileRouter(router *mux.Router) {
	subrouter := router.PathPrefix("/player").Subrouter()
	subrouter.HandleFunc("/{playerName}", playerProfileHandler())
}

func playerProfileHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["playerName"]
		js, err := json.Marshal(name)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(js)
	}
}
