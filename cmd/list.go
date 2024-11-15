/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all entries",
	Long:  `List alla entries in database.`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		var users []User
		db.Find(&users)

		tableData := pterm.TableData{{"First Name", "Last Name"}}
		for _, user := range users {
			tableData = append(tableData, []string{user.FirstName, user.LastName})
		}
		pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
