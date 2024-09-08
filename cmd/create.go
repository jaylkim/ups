/*
Copyright © 2024 JK <jay.luke.kim@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var flagPatient bool

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
			fmt.Println("Creating a patient")
		} else {
			fmt.Println("Creating other things")
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().BoolVarP(&flagPatient, "patient", "p", false, "Create a patient")
}
