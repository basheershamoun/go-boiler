package repository

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func ConnectDB() error {
	//dataSourceName := os.Getenv("DB_SOURCE")

	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	DB = db
	fmt.Println("😁 Connected to database")
	return nil
}

//func ConnectDB() error {
//	dataSourceName := os.Getenv("DB_SOURCE")
//
//	db, err := sql.Open("mysql", dataSourceName+"?charset=utf8mb4&parseTime=True")
//	if err != nil {
//		return err
//	}
//	err = db.Ping()
//	if err != nil {
//		return err
//	}
//	DB = db
//	fmt.Println("😁 Connected to database")
//	return nil
//}
