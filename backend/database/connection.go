package database

import (
	"github.com/Gurv33r/go-env"
	"github.com/go-pg/pg/v10"
)

func NewConn() *pg.DB {
	ENV := env.EnvAsMap([]string{"DBUSER", "DBPASS", "DBHOST", "DBPORT", "DBNAME"})
	db := pg.Connect(&pg.Options{
		Addr:     ENV["DBHOST"] + ":" + ENV["DBPORT"],
		User:     ENV["DBUSER"],
		Password: ENV["DBPASS"],
		Database: ENV["DBNAME"],
	})
	return db
}
