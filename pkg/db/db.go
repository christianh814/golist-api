package db

import (
	"errors"

	"github.com/christianh814/golist-api/pkg/entities"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// define an instance of a DB and the error if it occurs
var Instance *gorm.DB
var err error

// Connect connects to the database
func Connect(connectionString string) error {
	// connect to the database
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return errors.New("Error connecting to database: " + err.Error())
	}

	// if we're here, we're good
	return nil
}

// InitDB initializes the database with the proper tables
func InitDB() error {
	// create the tables
	log.Info("Creating tables if needed...")
	return Instance.AutoMigrate(&entities.Product{})
}
