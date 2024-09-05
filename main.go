/*
Copyright © 2024 JK <jay.luke.kim@gmail.com>
*/
package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jaylkim/ups/cmd"
	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func main() {
	defer database.Close()
	cmd.Execute()
}

func init() {
	var err error
	database, err = sql.Open("sqlite3", "./db/example.db")
	if err != nil {
		log.Fatal(err)
	}
	// Create a table
	statement, err := database.Prepare(`CREATE TABLE IF NOT EXISTS patient (
		pid INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		age INTEGER NOT NULL,

	)`)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	fmt.Println("Table patient successfully created!")
}
