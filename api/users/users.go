package userService

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/alexgemas17/api-rest-go/api/auth"
	"github.com/alexgemas17/api-rest-go/api/envs"
	"github.com/alexgemas17/api-rest-go/api/store"
	"github.com/alexgemas17/api-rest-go/api/types"
	"github.com/alexgemas17/api-rest-go/api/utils"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var errEmailRequired = errors.New("email is required")
var errFirstNameRequired = errors.New("first name is required")
var errLastNameRequired = errors.New("last name is required")
var errPasswordRequired = errors.New("password is required")

type UserService struct {
	store store.Store
}

func NewUserService(s store.Store) *UserService {
	return &UserService{store: s}
}

func (s *UserService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users/register", s.handleUserRegister).Methods("POST")
	r.HandleFunc("/users/login", s.handleUserLogin).Methods("GET")
}

func (s *UserService) handleUserRegister(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, types.ErrorResponse{Error: "Error reading request body"})
		return
	}

	defer r.Body.Close()

	// data validation
	var payload *types.User
	err = json.Unmarshal(body, &payload)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, types.ErrorResponse{Error: "Error request payload"})
		return
	}

	if err := validateUserPayload(payload); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, types.ErrorResponse{Error: err.Error()})
		return
	}

	// hash the password
	hashedPassword, err := HashPassword(payload.Password)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, types.ErrorResponse{Error: "Error generating hashed password"})
		return
	}
	payload.Password = hashedPassword

	// registry user
	u, err := s.store.CreateUser(payload)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, types.ErrorResponse{Error: "Error creating user"})
		return
	}

	token, err := createAndSetAuthCookie(u.ID, w)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, types.ErrorResponse{Error: "Error creating session"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, token)
}

func (s *UserService) handleUserLogin(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, types.ErrorResponse{Error: "Error reading request body"})
		return
	}

	defer r.Body.Close()

	// data validation
	var payload *types.LoginPayload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, types.ErrorResponse{Error: "Error request payload"})
		return
	}

	if err := validateUserLoginPayload(payload); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, types.ErrorResponse{Error: err.Error()})
		return
	}

	u, err := s.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, types.ErrorResponse{Error: "Invalid username/password"})
		return
	}

	err = ComparePassword(payload.Password, u.Password)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, types.ErrorResponse{Error: "Invalid username/password"})
		return
	}

	token, err := createAndSetAuthCookie(u.ID, w)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, types.ErrorResponse{Error: "Error creating session"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, token)
}

func validateUserPayload(user *types.User) error {
	if user.Email == "" {
		return errEmailRequired
	}

	if user.FirstName == "" {
		return errFirstNameRequired
	}

	if user.LastName == "" {
		return errLastNameRequired
	}

	if user.Password == "" {
		return errPasswordRequired
	}

	return nil
}

func validateUserLoginPayload(user *types.LoginPayload) error {
	if user.Email == "" {
		return errEmailRequired
	}

	if user.Password == "" {
		return errPasswordRequired
	}

	return nil
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePassword(password string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}

func createAndSetAuthCookie(userId int64, w http.ResponseWriter) (string, error) {
	secret := []byte(envs.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, userId)
	if err != nil {
		return "", err
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "Authorization",
		Value: token,
	})

	return token, nil
}
