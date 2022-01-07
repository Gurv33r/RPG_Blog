package database

import (
	"github.com/Gurv33r/go-env"
	"github.com/go-pg/pg/v10"
)

func NewConn() *pg.DB {
	ENV := env.EnvAsMap([]string{"USR", "PASSWORD", "HOST", "PORT", "DBNAME"})
	//dbconn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", ENV["HOST"], ENV["USR"], ENV["DBNAME"], ENV["PASSWORD"], ENV["PORT"])
	db := pg.Connect(&pg.Options{
		Addr:     ENV["HOST"] + ":" + ENV["PORT"],
		User:     ENV["USR"],
		Password: ENV["PASSWORD"],
		Database: ENV["DBNAME"],
	})
	return db
}
