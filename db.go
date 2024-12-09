package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func getDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", "gotemp_db_u:root#root@/gotemp_db")
	if err != nil {
		fmt.Println(err)
		return db, err
	}
	return db, nil
}
