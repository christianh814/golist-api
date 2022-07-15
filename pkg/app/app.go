package app

import (
	"net/http"

	"github.com/christianh814/golist-api/pkg/config"
	gh "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
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
	cors := gh.AllowedOrigins([]string{"*"})

	// try to start the app and log output
	log.Info("Starting server on port " + config.AppConfig.Port)
	log.Fatal(http.ListenAndServe(":"+config.AppConfig.Port, gh.CORS(cors)(router)))
}
