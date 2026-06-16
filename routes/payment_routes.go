package routes

import (
	"backend-api-commerce/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func PaymentRoutes(router *mux.Router) {

	router.Handle(
		"/webhook/payment",
		http.HandlerFunc(handler.PaymentWebhook),
	).Methods("POST")

	router.Handle(
		"/payments/webhook",
		http.HandlerFunc(handler.PaymentWebhook),
	).Methods("POST")

}