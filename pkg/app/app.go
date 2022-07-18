package app

import (
	"net/http"

	"github.com/christianh814/golist-api/pkg/config"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

// Start starts the application
func Start() {
	// create and register the router
	router := mux.NewRouter().StrictSlash(true)

	// register the routes
	router.HandleFunc("/api/products", GetProducts).Methods("GET")
	router.HandleFunc("/api/products/{id}", GetProductById).Methods("GET")
	router.HandleFunc("/api/products", CreateProduct).Methods("POST")
	router.HandleFunc("/api/products/{id}", UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/products/{id}", DeleteProduct).Methods("DELETE")

	// Set cors setting
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		//Debug:            true,
	})

	// set router handler
	handler := cors.Handler(router)

	// try to start the app and log output
	log.Info("Starting server on port " + config.AppConfig.Port)
	log.Fatal(http.ListenAndServe(":"+config.AppConfig.Port, handler))
}
