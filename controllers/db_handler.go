package controllers

import (
	"database/sql"
	"net/http"
)

func connect(w http.ResponseWriter) *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_latihan_pbp?parseTime=true&loc=Asia%2FJakarta")
	if err != nil {
		sendErrorResponse(w, 500, "Internal Server Error! Database Connection Error")
	}
	return db
}
