package apiserver

import (
	"github.com/sirupsen/logrus"
	"io"
	"net/http"

	"github.com/gorilla/mux"
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

	if err := s.ConfigureLogger(); err != nil {
		return err
	}

	s.logger.Info("Starting api server...")
	s.ConfigureRouter()
	return http.ListenAndServe(s.config.BindAddres, s.router)
}

func (s *APIServer) ConfigureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}
func (s *APIServer) ConfigureRouter() {
	s.router.HandleFunc("/", s.HandleHello())
}

func (s *APIServer) HandleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}