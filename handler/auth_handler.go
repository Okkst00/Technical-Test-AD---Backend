package handler

import (
	"backend-api-commerce/model"
	"backend-api-commerce/service"
	"encoding/json"
	"net/http"
	"strings"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var req model.RegisterRequest

	json.NewDecoder(r.Body).Decode(&req)

	err := service.Register(req)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			400,
		)
		return
	}

	json.NewEncoder(w).Encode(
		map[string]string{
			"message": "register success",
		},
	)
}

func Login(w http.ResponseWriter, r *http.Request) {

	var req model.LoginRequest

	json.NewDecoder(r.Body).Decode(&req)

	token, user, err := service.Login(req)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			401,
		)
		return
	}

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"token": token,
			"user": map[string]interface{}{
				"id":    user.ID,
				"nama":  user.Nama,
				"email": user.Email,
				"role":  user.Role,
			},
		},
	)
}

func Logout(w http.ResponseWriter, r *http.Request) {

	auth := r.Header.Get("Authorization")

	if auth == "" {
		http.Error(
			w,
			"token required",
			401,
		)
		return
	}

	token := strings.Replace(
		auth,
		"Bearer ",
		"",
		1,
	)

	err := service.Logout(token)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			400,
		)
		return
	}

	json.NewEncoder(w).Encode(
		map[string]string{
			"message": "logout success",
		},
	)
}