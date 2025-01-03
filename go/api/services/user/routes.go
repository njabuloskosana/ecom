package user

// Services will handle the actual requests we avoid adding it in the api.go for
// readability and management
import (
	"ecom/go/api/configs"
	"ecom/go/api/services/auth"
	"ecom/go/api/types"
	"ecom/go/api/utils"
	"fmt"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// Hanldes dependencies
type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

// Handles the routes
func (h *Handler) RegisterRoutes(router *mux.Router) {
	// definition of route services
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")

}

// Actual service functions
func (h *Handler) handleLogin(q http.ResponseWriter, r *http.Request) {
	// get json payload and check if user exisits and if not we create the new user
	var payload types.LoginUserDto
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(q, http.StatusBadRequest, err)
		return
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(q, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	returnedUser, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(q, http.StatusBadRequest, fmt.Errorf("user doesnt exists"))
		return
	}

	err = auth.ComparePasswords(returnedUser.Password, payload.Password)
	if err != nil {
		utils.WriteError(q, http.StatusBadRequest, fmt.Errorf("invalid password"))
		return
	}

	// generate token
	secret := []byte(configs.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, returnedUser.ID)
	if err != nil {
		utils.WriteError(q, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(q, http.StatusOK, map[string]string{"token": token})

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// get json payload and check if user exisits and if not we create the new user
	var payload types.RegisterUserDto
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user already exists"))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(&types.User{
		FirstName:   payload.FirstName,
		LastName:    payload.LastName,
		Email:       payload.Email,
		Password:    hashedPassword,
		CreatedTime: utils.GetCurrentTime(),
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "user created"})

}
