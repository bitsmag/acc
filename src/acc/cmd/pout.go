package cmd

import (
	"acc/db"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// Pay Out command
var poutCmd = &cobra.Command{
	Use:   "pout [amount to pay out from the account]",
	Short: "Pay out a certain amount from the account         FLAGS: -d --date [(DD).(MM).YYYY]; -c  category <category>",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		poutAmount, err := getPoutAmountFromArgs(args)
		if err != nil {
			return fmt.Errorf("couldn't get amount from arguments: %v", err)
		}

		if err := db.ProcessTransaction(poutAmount, poutDate, poutCategory); err != nil {
			return fmt.Errorf("couldn't process transaction: %v", err)
		}

		return balanceCmdRunner()
	},
}

func getPoutAmountFromArgs(args []string) (float64, error) {
	var amount float64
	var err error
	if amount, err = strconv.ParseFloat(args[0], 64); err != nil {
		return 0, fmt.Errorf("couldn't parse amount, please check provided argument: %v", err)
	}
	amount = -amount
	return amount, nil
}
