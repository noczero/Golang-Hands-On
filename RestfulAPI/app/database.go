package app

import (
	"database/sql"
	"github.com/noczero/Golang-Hands-On/RestfulAPI/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:zero_pwd@tcp(localhost:3306)/restful_api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
