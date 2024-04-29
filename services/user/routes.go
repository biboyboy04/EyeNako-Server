package user

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", h.handleGet).Methods("GET")
}

func(h *Handler) handleGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sucessfully Get User")
}