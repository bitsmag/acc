package cmd

import (
	"acc/db"
	"acc/useroutput"
	"fmt"

	"github.com/spf13/cobra"
)

// Balance command
var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Prints out the current balance of the account",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return balanceCmdRunner()
	},
}

func balanceCmdRunner() error {
	var balance float64
	var err error
	if balance, err = db.ReadBalance(); err != nil {
		return fmt.Errorf("couldn't read balance from database: %v", err)
	}

	balancePrintabnle := useroutput.CreateBalance(balance)
	useroutput.PrintPrintablelns(balancePrintabnle)
	return nil
}
