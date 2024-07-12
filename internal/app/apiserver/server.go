package apiserver

//http -v --session=user http://localhost:8080/private/whoami

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	router "github.com/c4erries/server/internal/app/apiserver/Routers"
	"github.com/c4erries/server/internal/app/model"
	"github.com/c4erries/server/internal/app/store"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// Константные ключи/поля
const (
	sessionName        = "usersession"
	ctxKeyUser  ctxKey = iota
	ctxKeyRequestID
)

// Особые типы ошибок
var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
	errNotAuthenticated         = errors.New("no authenticated")
)

// Тип для ключей контекста
type ctxKey int8

// Файл, определяющий сервер
type server struct {
	router       *mux.Router
	store        store.Store
	sessionStore sessions.Store
}

// Конструктор сервера (Хранилище(БД) -> сервер)
func newServer(store store.Store, sessionStore sessions.Store) *server {
	s := &server{
		router:       mux.NewRouter(),
		store:        store,
		sessionStore: sessionStore,
	}

	s.configureRouter()

	return s
}

// Ф-ция делегирования запросов в Роутер (используется для учучшения условий тестирования)
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// Конфигурация роутера (запросов)
func (s *server) configureRouter() {
	s.router.Use(s.setRequestID)
	//две строки ниже что-то делают ? вроде нет, а должны
	//s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	//s.router.Use(handlers.CORS(handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})))
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
	s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods("POST")
	s.router.HandleFunc("/allusers", s.handleUsersListAll()).Methods("GET")
	router.ConfigureMatchListSubRouter(s.router)
	router.ConfigurePlayersRouter(s.router)
	router.ConfigurePlayerProfileRouter(s.router, s.store)
	private := s.router.PathPrefix("/private").Subrouter()
	private.Use(s.authenticateUser)
	private.HandleFunc("/whoami", s.handleWhoami()).Methods("GET")
	private.HandleFunc("/unauth", s.handleSessionTerminate()).Methods("POST")
}

// MIDW Присвоение запросам ID
func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))
	})
}

// MIDW Проверка аутентификации
func (s *server) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		id, ok := session.Values["user_id"]
		if !ok {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		u, err := s.store.User().Find(id.(int))
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		//Подтверждение и передача управления следующему оператору
		//Context используется, чтобы далее проверка пользователя не требовалась
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))
	})
}

// Выход из сессии
func (s *server) handleSessionTerminate() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		session.Values["user_id"] = -1
		if err := s.sessionStore.Save(r, w, session); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		c := &http.Cookie{
			Name:     "usersession",
			Value:    "",
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: true,
		}

		http.SetCookie(w, c)
		s.respond(w, r, http.StatusOK, nil)

	})
}

// Доступная только для авторизованных пользователей ф-ция
func (s *server) handleWhoami() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*model.User))
	}
}

// Хэндл запроса на создание пользователей (Использует служебные методы ошибки и ответа)
func (s *server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Nickname string `json:"nickname"`
		Email    string `json:"email"`
		Password string `json:"password"`
		PlayerID int    `json:"playerid"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Nickname: req.Nickname,
			Email:    req.Email,
			Password: req.Password,
			PlayerID: req.PlayerID,
		}

		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u.Sanitize()
		s.respond(w, r, http.StatusCreated, u)
	}
}

// Хэндл запроса на список всех пользователей
func (s *server) handleUsersListAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		Us, err := s.store.User().ListAll()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		data, err := json.Marshal(Us)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}

		w.Write(data)
	}
}

// Хэндл запроса на вход (новая сессия) пользователя
func (s *server) handleSessionsCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
			return
		}

		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		session.Values["user_id"] = u.ID
		if err := s.sessionStore.Save(r, w, session); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

// Служебный метод ошибки
func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

// Служебный метод ответа
func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
