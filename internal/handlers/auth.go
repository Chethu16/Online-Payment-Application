package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Chethu16/Online-Payment-Application/internal/models"
	"github.com/Chethu16/Online-Payment-Application/pkg/payload"
)

type authenticationHandler struct {
	authentication models.Authentication
}

func NewAuthenticationHandler(authentication models.Authentication) *authenticationHandler {
	return &authenticationHandler{
		authentication,
	}
}

func (ah *authenticationHandler) LoginRequest(w http.ResponseWriter, r *http.Request) {
	var user models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(payload.FailurePayload{Message: err.Error(), Status: "failed"})
		return
	}
	res, err := ah.authentication.Login(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(payload.FailurePayload{Message: err.Error(), Status: "failed"})
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (ah *authenticationHandler) RegisterRequest(w http.ResponseWriter, r *http.Request) {
	var user models.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(payload.FailurePayload{Message: err.Error(), Status: "failed"})
		return
	}
	res, err := ah.authentication.Register(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(payload.FailurePayload{Message: err.Error(), Status: "failed"})
		return
	}
	json.NewEncoder(w).Encode(res)
}
