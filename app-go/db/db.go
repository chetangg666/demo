package db

import (
	"database/sql"
	"fmt"
	"log"

	//	"os"

	_ "github.com/go-sql-driver/mysql"
)

type person struct {
	Id        int
	LastName  string
	FirstName string
}

func DataBase() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/")
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err.Error())
	}
	fmt.Println("Connection opened successfully!")

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	//// Create Database if not exist
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS testdb222")
	if err != nil {
		log.Fatal(err)
	}

	//// Use our database
	_, err = db.Exec("USE testdb222")
	if err != nil {
		log.Fatal(err)
	}

	//// Create Table if not exists
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Persons (PersonID int NOT NULL AUTO_INCREMENT, LastName varchar(255), FirstName varchar(255), PRIMARY KEY (PersonID));")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
