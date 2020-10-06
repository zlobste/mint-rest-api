package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	
	"github.com/zlobste/mint-rest-api/internal/app/apiserver/middlewares"
	"github.com/zlobste/mint-rest-api/internal/app/store"
)

type server struct {
	router  *mux.Router
	logger  *logrus.Logger
	store   store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}
	s.ConfigureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) ConfigureRouter() {

	s.router.Use(middlewares.ContentTypeMiddleware)
	s.router.HandleFunc("/register", s.SignUp()).Methods("POST")
	s.router.HandleFunc("/login", s.SignIn()).Methods("POST")
	subrouter := s.router.PathPrefix("/api").Subrouter()
	subrouter.Use(middlewares.TokenAuthMiddleware)
	subrouter.HandleFunc("/user", s.GetUser()).Methods("GET")
}
