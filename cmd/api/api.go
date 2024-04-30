package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/biboyboy04/EyeNako-Server/services/user"
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
	// make all router starts with the prefix given
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	fmt.Printf("Server running on %s... \n", s.addr)
	return http.ListenAndServe(s.addr, router)
}