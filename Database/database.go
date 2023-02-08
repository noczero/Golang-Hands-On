package Database

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	sourceName := "root:zero_pwd@tcp(localhost:3306)/zero_db?parseTime=true" // make connection to localhost and use flag parse time to automatically parse date type to go time
	db, err := sql.Open("mysql", sourceName)

	if err != nil {
		panic(err)
	}

	// database pooling configuration
	db.SetMaxIdleConns(10)                  // maximal idle connection pool
	db.SetMaxOpenConns(100)                 // maximal open connection pool
	db.SetConnMaxIdleTime(5 * time.Minute)  // maximal time of connection pool idle, set to 5 minuets
	db.SetConnMaxLifetime(60 * time.Minute) // maximal time of connection pool both active or idle, set to 60 minutes

	return db
}
