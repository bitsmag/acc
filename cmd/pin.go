package cmd

import (
	"acc/db"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// Pay In command
var pinCmd = &cobra.Command{
	Use:   "pin [amount to pay into the account]",
	Short: "Pay in a certain amount to the account            FLAGS: -d --date [(DD).(MM).YYYY]; -c  category <category>",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pinAmount, err := getPinAmountFromArgs(args)
		if err != nil {
			return fmt.Errorf("couldn't get amount from arguments: %v", err)
		}

		if err := db.ProcessTransaction(pinAmount, pinDate, pinCategory); err != nil {
			return fmt.Errorf("couldn't process transaction: %v", err)
		}

		return balanceCmdRunner()
	},
}

func getPinAmountFromArgs(args []string) (float64, error) {
	var amount float64
	var err error
	if amount, err = strconv.ParseFloat(args[0], 64); err != nil {
		return 0, fmt.Errorf("couldn't parse amount, please check provided argument: %v", err)
	}
	return amount, nil
}
