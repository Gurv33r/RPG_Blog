package database

import (
	"github.com/Gurv33r/go-env"
	"github.com/go-pg/pg/v10"
)

func NewConn() *pg.DB {
	ENV := env.EnvAsMap([]string{"DBUSER", "PASSWORD", "HOST", "DBPORT", "DBNAME"})
	db := pg.Connect(&pg.Options{
		Addr:     ENV["HOST"] + ":" + ENV["DBPORT"],
		User:     ENV["DBUSER"],
		Password: ENV["PASSWORD"],
		Database: ENV["DBNAME"],
	})
	return db
}
