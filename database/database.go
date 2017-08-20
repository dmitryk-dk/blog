package database

import (
	"database/sql"
	"github.com/dmitryk-dk/blog/config"
)

var dbInstance *sql.DB

func Connect (config *config.Config) (*sql.DB, error) {
	if dbInstance == nil {
		db, err := sql.Open(config.DbDriverName, config.User+":"+config.Password+"@"+config.Host+"/"+config.DbName)
		if err != nil {
			return nil, err
		}
		dbInstance = db
	}
	return dbInstance, nil
}


