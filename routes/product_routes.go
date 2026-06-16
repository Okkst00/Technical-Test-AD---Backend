package routes

import (
	"backend-api-commerce/handler"
	"backend-api-commerce/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func ProductRoutes(router *mux.Router) {

	router.Handle(
		"/products",
		http.HandlerFunc(handler.GetProducts),
	).Methods("GET")

	router.Handle(
		"/products/{id}",
		http.HandlerFunc(handler.GetProductByID),
	).Methods("GET")

	router.Handle(
		"/products",
		middleware.JWTMiddleware(
			middleware.RoleMiddleware("admin")(
				http.HandlerFunc(handler.CreateProduct),
			),
		),
	).Methods("POST")

	router.Handle(
		"/products/{id}",
		middleware.JWTMiddleware(
			middleware.RoleMiddleware("admin")(
				http.HandlerFunc(handler.UpdateProduct),
			),
		),
	).Methods("PUT")

	router.Handle(
		"/products/{id}",
		middleware.JWTMiddleware(
			middleware.RoleMiddleware("admin")(
				http.HandlerFunc(handler.DeleteProduct),
			),
		),
	).Methods("DELETE")
}