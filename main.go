package main

import (
	"database/sql"
	"fmt"
)

func main() {

	var connString string
	var db *sql.DB

	for true {
		fmt.Print("Input a connection string: ")
		_, err := fmt.Scanln(&connString)

		if err == nil {
			break
		}

		db, err := sql.Open("postgres", connString)

	}
	fmt.Print("Connected to db")

	for true {
		//TODO write a while that takes lines of sql code
	}

	defer db.Close()
}
