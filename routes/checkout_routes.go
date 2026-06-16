package routes

import (
	"backend-api-commerce/handler"
	"backend-api-commerce/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func CheckoutRoutes(router *mux.Router) {

	router.Handle(
		"/checkout/calculate",
		middleware.JWTMiddleware(
			middleware.RoleMiddleware("admin", "member")(
				http.HandlerFunc(handler.CheckoutCalculate),
			),
		),
	).Methods("POST")
}