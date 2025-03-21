package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ndraoo/restapi-go/service/user"
)

type APIserver struct {
	addr string
	db   *sql.DB
}

func NewAPIserver(addr string, db *sql.DB) *APIserver {
	return &APIserver{addr: addr, db: db}
}

func (s *APIserver) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	log.Println("server is running on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
