package cmd

import (
	"acc/types"

	"github.com/spf13/cobra"
)

// AccCmd is the root command.
// Every other command attached to accCmd is a child command to it
var AccCmd = &cobra.Command{
	Use:   "acc",
	Short: "acc helps you manage your savings",
	Long:  `acc provides an account to manage your incomes and expenses.`}

// Execute initializes the commands and executes the user input
func Execute() error {
	initCommands()
	if err := AccCmd.Execute(); err != nil {
		return err
	}
	return nil
}

// flag variables hold the corresponding userinput and are processed in the command-handlers
var pinDate types.Date
var pinCategory types.Category
var poutDate types.Date
var poutCategory types.Category
var logOrder types.Order
var logDate types.Date
var logCategory types.Category

// initCommands adds commands to the root command and sets flags appropriately
func initCommands() {
	AccCmd.AddCommand(pinCmd)
	AccCmd.AddCommand(poutCmd)
	AccCmd.AddCommand(balanceCmd)
	AccCmd.AddCommand(logCmd)
	pinCmd.Flags().VarP(&pinDate, "date", "d", "The date on which the transaction took place")
	pinCmd.Flags().VarP(&pinCategory, "category", "c", "A category which gets attached to the transaction")
	poutCmd.Flags().VarP(&poutDate, "date", "d", "The date on which the transaction took place")
	poutCmd.Flags().VarP(&poutCategory, "category", "c", "A category which gets attached to the transaction")
	logCmd.Flags().VarP(&logOrder, "order", "o", "The order by which the log is printed to the terminal ['year' | 'category']")
	logCmd.Flags().VarP(&logDate, "date", "d", "Prints only logs for transactions which took place on the provided date")
	logCmd.Flags().VarP(&logCategory, "category", "c", "Prints only logs for transactions which hold the provided category")
}
