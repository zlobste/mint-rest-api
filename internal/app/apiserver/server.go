package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

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

	s.router.Use(s.ContentTypeMiddleware)
	s.router.HandleFunc("/register", s.SignUp()).Methods("POST")
	s.router.HandleFunc("/login", s.SignIn()).Methods("POST")
	subrouter := s.router.PathPrefix("/api").Subrouter()
	subrouter.Use(s.TokenAuthMiddleware)
	subrouter.HandleFunc("/user", s.GetUser()).Methods("GET")
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error){
	s.respond(w, r, code, map[string] string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
