package handler

import (
	"backend-api-commerce/helper"
	"backend-api-commerce/service"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type PaymentWebhookRequest struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
}


func PaymentWebhook(w http.ResponseWriter, r *http.Request) {

	paymentSecret := os.Getenv("PAYMENT_SECRET")
	fmt.Println("SECRET:", paymentSecret)
	if paymentSecret == "" {
		http.Error(w, "payment secret not configured", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// ambil signature dari header
	signature := r.Header.Get("X-Signature")
	if signature == "" {
		http.Error(w, "missing signature", 401)
		return
	}

	// baca raw body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// verify signature
	if !helper.VerifySignature(paymentSecret, body, signature) {
		http.Error(w, "invalid signature", 401)
		return
	}

	// decode payload
	var req struct {
		OrderID int    `json:"order_id"`
		Status  string `json:"status"`
	}

	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// process payment
	err = service.ConfirmPayment(req.OrderID, req.Status)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "payment updated",
	})
}