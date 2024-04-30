package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/biboyboy04/EyeNako-Server/services/auth"
	"github.com/biboyboy04/EyeNako-Server/types"
	"github.com/biboyboy04/EyeNako-Server/utils"
	"github.com/go-playground/validator/v10"

	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request){

	// get JSON payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return 
	}
	
	// Check if user exist
	_, err := h.store.GetUserByEmail(payload.Email)
	
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		 return 
	}


	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	err = h.store.CreateUser(types.User{
		Username: payload.Username,
		Password: hashedPassword,
		Email: payload.Email,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}


func(h *Handler) handleGetEmail(w http.ResponseWriter, r *http.Request){
	var payload types.User

	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// check value of payload, but ithinkg its a json
	u, err := h.store.GetUserByEmail("user@example.com")
	if err != nil {
		log.Fatal(err)
	} 
	fmt.Println(u, "user outpuit")
	
}