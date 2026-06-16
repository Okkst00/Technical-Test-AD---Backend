package handler

import (
	"backend-api-commerce/model"
	"backend-api-commerce/service"
	"encoding/json"
	"net/http"
)

type CheckoutHandler struct{}


func CheckoutCalculate(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var req model.CalculateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	result, err := service.CalculateCheckout(req)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}


	json.NewEncoder(w).Encode(result)
}