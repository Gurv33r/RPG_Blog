package database

import (
	"fmt"
	"log"

	"github.com/Gurv33r/go-env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConn() *gorm.DB {
	env.LoadFrom("./env/db.env")
	ENV := env.EnvAsMap([]string{"USR", "PASSWORD", "HOST", "PORT", "DBNAME", "DIALECT"})
	dbconn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", ENV["HOST"], ENV["USR"], ENV["DBNAME"], ENV["PASSWORD"], ENV["PORT"])
	db, err := gorm.Open(postgres.Open(dbconn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
