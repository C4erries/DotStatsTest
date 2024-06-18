package apiserver

import (
	"io"
	"net/http"

	router "github.com/c4erries/server/internal/app/apiserver/Routers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("starting api server")
	corsObj := handlers.AllowedOrigins([]string{"*"})
	return http.ListenAndServe(s.config.BindAddr, handlers.CORS(corsObj)(s.router))
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hi", s.handleHello())
	router.ConfigureStatsSubRouter(s.router)
	router.ConfigureMatchListSubRouter(s.router)
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hi")
	}
}
