package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	DB_DRIVER := "mysql"
	DB_USER := "root"
	DB_PASS := ""
	DB_NAME := "go_crud"

	db, err := sql.Open(DB_DRIVER, DB_USER+":"+DB_PASS+"@/"+DB_NAME)
	if err != nil {
		panic("failed to connect database")
	}
	return db, nil
}
