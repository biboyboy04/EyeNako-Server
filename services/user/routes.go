package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/biboyboy04/EyeNako-Server/config"
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
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request){
		// get JSON payload
		var payload types.LoginUserPayload

		// Attach request data to payload
		if err := utils.ParseJSON(r, &payload); err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}
	
		// Data Validation 
		if err := utils.Validate.Struct(payload); err != nil {
			errors := err.(validator.ValidationErrors)
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
			return 
		}
		
		// Check if user exist
		u, err := h.store.GetUserByEmail(payload.Email)
		
		// User not found err handler
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
			return
		}
	
		// Check if correct pass
		if !auth.ComparePasswords(u.Password, []byte(payload.Password)) {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
			return 
		}
		secret := []byte(config.Envs.JWTSecret)
		token, err := auth.CreateJWT(secret, u.ID)

		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request){

	var payload types.RegisterUserPayload

	// Attach request data to payload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Data Validation 
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return 
	}
	
	// Check if user exist
	_, err := h.store.GetUserByEmail(payload.Email)
	
	// If email exists already, error
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