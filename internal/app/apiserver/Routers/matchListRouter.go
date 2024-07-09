package router

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// список матчей игрока  /matches/{playerName}
func ConfigureMatchListSubRouter(router *mux.Router) *mux.Router {
	subrouter := router.PathPrefix("/matches").Subrouter()
	subrouter.HandleFunc("/{playerName}", handlePostMatchList())
	return subrouter
}

// хэндл матчей игрока, которые есть на сервере(в бд) (нужно подключить к бд и сделать проверку прав на доступ)
func handlePostMatchList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		playerNickName := vars["playerName"]
		//запрос в бд на этого пользователя и его список матчей
		file, err := os.ReadFile("jsons/matches/" + playerNickName + ".json")
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(file)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	}
}
