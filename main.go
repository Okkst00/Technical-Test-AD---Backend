package main

import (
	"fmt"
	"log"
	"net/http"

	"backend-api-commerce/config"
	"backend-api-commerce/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)


func main(){

	if err := godotenv.Load(); err != nil {
		log.Fatal("gagal load .env")
	}

	config.ConnectDB()

	router:=mux.NewRouter()
	routes.ProductRoutes(router)
	routes.AuthRoutes(router)
	routes.OrderRoutes(router)
	routes.PaymentRoutes(router)
	routes.CheckoutRoutes(router)

	handler := cors.New(cors.Options{

		AllowedOrigins: []string{
			"http://localhost:3000",
		},

		AllowedMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},

		AllowedHeaders: []string{
			"*",
		},

	}).Handler(router)

	fmt.Println("server running :8080")

	http.ListenAndServe(
		":8080",
		handler,
	)


}