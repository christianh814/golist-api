package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/christianh814/golist-api/pkg/app"
	"github.com/christianh814/golist-api/pkg/config"
	"github.com/christianh814/golist-api/pkg/db"
)

func main() {
	// Load Application configurations
	if err := config.LoadConfig(); err != nil {
		log.Fatal(err)
	}

	// Connect to the database
	db.Connect(config.AppConfig.Conn)

	// Create the tables if needed
	if err := db.InitDB(); err != nil {
		log.Fatal(err)
	}

	// Start the application
	app.Start()
}
