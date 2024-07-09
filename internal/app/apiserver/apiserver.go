package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/c4erries/server/internal/app/store/sqlstore"
	"github.com/gorilla/sessions"
	"github.com/rs/cors"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	s := newServer(store, sessionStore)
	// используется rs/cors, только так cors разрешает post запрос с фронтенда
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{config.FrontendUrl}, //адресса, имеющие доступ к серверу
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true, //для cookie (вроде)
	}).Handler(s.router)

	return http.ListenAndServe(config.BindAddr, corsHandler)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
