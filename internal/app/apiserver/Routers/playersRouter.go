package router

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Player struct {
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	Matches    int    `json:"matches"`
	Wins       int    `json:"wins"`
	IsOnline   bool   `json:"isOnline"`
}

// список игроков /players
func ConfigurePlayersRouter(router *mux.Router) *mux.Router {
	subrouter := router.PathPrefix("/players").Subrouter()
	subrouter.HandleFunc("", handlePlayersList()).Methods("POST", "GET")
	return subrouter
}

// хэндл получения списка игроков (нужно подключить к бд)
func handlePlayersList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a, err := strconv.Atoi(r.URL.Query().Get("a"))
		if err != nil {
			log.Fatal(err)
		}
		b, err := strconv.Atoi(r.URL.Query().Get("b"))
		if err != nil {
			log.Fatal(err)
		}
		file, err := os.ReadFile("jsons/Players.json")
		if err != nil {
			log.Fatal(err)
		}
		val := []Player{}
		err = json.Unmarshal(file, &val)
		if err != nil {
			log.Fatal(err)
		}
		data, err := json.Marshal(val[a:b])
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(data)
		if err != nil {
			log.Fatal(err)
		}
	}
}
