package routes

import (
	"backend-api-commerce/handler"

	"github.com/gorilla/mux"
)


func AuthRoutes(router *mux.Router){

	router.HandleFunc(
		"/auth/register",
		handler.Register,
	).Methods("POST")


	router.HandleFunc(
		"/auth/login",
		handler.Login,
	).Methods("POST")


	router.HandleFunc(
		"/auth/logout",
		handler.Logout,
	).Methods("POST")

}