/*
Copyright © 2024 JK <jay.luke.kim@gmail.com>
*/
package cmd

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var flagPatient bool

var database *sql.DB

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if flagPatient {
			createPatient(database, args[0], args[1], args[2])
			fmt.Println("A patient is created.")
		} else {
			fmt.Println("Creating other things")
		}
	},
}

func createPatient(db *sql.DB, name string, age string, sex string) (int64, error) {
	statement, err := db.Prepare("INSERT INTO patient (name, age, sex) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	_, err = statement.Exec(
		name,
		age,
		sex,
	)
	if err != nil {
		return 0, err
	}

	return 1, nil
}

func init() {
	var err error
	database, err = sql.Open("sqlite3", "./db/example.db")
	if err != nil {
		log.Fatal(err)
	}
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().BoolVarP(&flagPatient, "patient", "p", false, "Create a patient")
}
