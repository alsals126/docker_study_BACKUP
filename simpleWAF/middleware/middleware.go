package middleware

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func DbConnection() *sql.DB {
	var dbinfo = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		"jeongmin", "1234asdf#$", "jeongmindb")
	var db *sql.DB

	var err error
	db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		log.Printf("DB ERRROR(Connection) : %v\n", err)
		return nil
	}
	return db
}
