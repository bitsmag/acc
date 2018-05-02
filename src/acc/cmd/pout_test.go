package cmd

import "testing"

func TestGetPoutAmountFromArgs01Success(t *testing.T) {
	args := []string{"123.123", "b", "c", "d"}
	amount, err := getPoutAmountFromArgs(args)

	if amount != -123.123 {
		t.Errorf("cmd.pout.getPoutAmountFromArgs(args), got: %v, want: %v.", amount, -123.123)
	}
	if err != nil {
		t.Errorf("Error occured while attempting to read amount from args: %v.", err.Error())
	}
}

func TestGetPoutAmountFromArgs02Error(t *testing.T) {
	args := []string{"a"}
	_, err := getPoutAmountFromArgs(args)

	if err == nil {
		t.Errorf("Expected error for invalid argument but no error was returned.")
	}
}
