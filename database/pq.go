package database

import (
	"fmt"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "docker"
	dbname   = "tsc-1"
)

func ConnectDB() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		log.Error(fmt.Sprintf("Error while connecting to the database: %v", err))
		return nil
	}
	fmt.Println("Successfully connected! DB")

	return db

}
