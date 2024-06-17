package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Rune struct {
	Name   string `json:"Name"`
	During int8   `json:"During"`
}

func ConfigureStatsSubRouter(router *mux.Router) *mux.Router {
	subrouter := router.PathPrefix("/stats").Subrouter()
	subrouter.HandleFunc("/wr", handlePostWRStats()).Methods("POST", "GET")
	return subrouter
}

func handlePostWRStats() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		haste := Rune{"Haste", 9}
		h, err := json.Marshal(haste)
		if err != nil {
			log.Fatal(err)
			return
		}
		_, err = w.Write(h)
		if err != nil {
			log.Fatal(err)
		}
	}
}
