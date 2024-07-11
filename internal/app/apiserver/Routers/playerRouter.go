package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/c4erries/server/internal/app/store"
	"github.com/gorilla/mux"
)

// данные о игроке /player/{playerName}
func ConfigurePlayerProfileRouter(router *mux.Router, store store.Store) {
	subrouter := router.PathPrefix("/player").Subrouter()
	store.User()
	subrouter.HandleFunc("/{playerName}", playerProfileHandler(store))
}

// хэндл получения данных о пользователе (нужно подключить бд и настроить проверку прав на доступ к запросу)
func playerProfileHandler(s store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["playerName"]
		u, err := s.User().FindByNickname(name)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		js, err := json.Marshal(u)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(js)
	}
}
