package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db: db,
	}
}

func (s *APIServer) Run () error {
	router := mux.NewRouter()
	
	fmt.Printf("Server running on %s... \n", s.addr)
	return http.ListenAndServe(s.addr, router)
}