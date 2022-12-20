package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
    "log"
)

var (
    DB *gorm.DB
    err error
)

func Connect() {
    connString := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
    DB, err = gorm.Open(postgres.Open(connString))
    if err != nil {
        log.Panic("Error connecting to the database")
    }
}
