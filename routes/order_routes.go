package routes

import (
	"backend-api-commerce/handler"
	"backend-api-commerce/middleware"
	"net/http"

	"github.com/gorilla/mux"
)


func OrderRoutes(router *mux.Router) {

	// MEMBER + ADMIN (lihat history masing-masing)
	router.Handle(
		"/orders/history",
		middleware.JWTMiddleware(
			middleware.RoleMiddleware("admin", "member")(
				http.HandlerFunc(handler.GetOrderHistory),
			),
		),
	).Methods("GET")

	// CREATE ORDER (admin only)
	router.Handle(
		"/orders",
		middleware.JWTMiddleware(
			middleware.RoleMiddleware("admin")(
				http.HandlerFunc(handler.CreateOrder),
			),
		),
	).Methods("POST")

	// GET ALL ORDERS (admin only)
	router.Handle(
		"/orders",
		middleware.JWTMiddleware(
			middleware.RoleMiddleware("admin")(
				http.HandlerFunc(handler.GetOrders),
			),
		),
	).Methods("GET")

	router.Handle(
		"/orders/{id}",
		middleware.JWTMiddleware(
			http.HandlerFunc(handler.GetOrderByID),
		),
	).Methods("GET")
}