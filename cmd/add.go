/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new entry",
	Long:  `Add new entry to the database`,
	Run: func(cmd *cobra.Command, args []string) {
		firstName, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Please enter your First Name").Show()
		lastName, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Please enter your Last Name").Show()
		fmt.Println("First Name:", firstName)
		fmt.Println("Last Name:", lastName)

		db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

		if err != nil {
			fmt.Println(err.Error())
			pterm.Error.Println("Error connecting to database")
		}
		db.AutoMigrate(&User{})

		db.Create(&User{FirstName: firstName, LastName: lastName})

		pterm.Success.Println("Entry added to database")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
