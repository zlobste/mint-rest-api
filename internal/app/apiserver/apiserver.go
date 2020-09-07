package apiserver

import (
	"database/sql"
	"net/http"
	
	"github.com/zlobste/mint-rest-api/internal/app/store/sqlstore"
)

func Start(config Config) error {
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
