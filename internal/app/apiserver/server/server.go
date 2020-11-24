package server

import (
	"database/sql"
	"net/http"
	
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	
	"github.com/zlobste/mint-rest-api/internal/app/store"
	"github.com/zlobste/mint-rest-api/internal/app/store/sqlstore"
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

func (server *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server.router.ServeHTTP(w, r)
}

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	store := sqlstore.New(db)
	server := newServer(store)
	return http.ListenAndServe(config.BindAddres, server.router)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL);
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, err
}
