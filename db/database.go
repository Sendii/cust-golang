package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func PConenctDB()(*sql.DB, error){
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_belajar_golang")
	if err != nil {
		return nil, err
	}
	return db, nil
}