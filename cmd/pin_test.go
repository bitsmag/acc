package cmd

import "testing"

func TestGetPinAmountFromArgs01Success(t *testing.T) {
	args := []string{"123.123", "b", "c", "d"}
	amount, err := getPinAmountFromArgs(args)

	if amount != 123.123 {
		t.Errorf("cmd.pin.getPinAmountFromArgs(args), got: %v, want: %v.", amount, 123.123)
	}
	if err != nil {
		t.Errorf("Error occured while attempting to read amount from args: %v.", err.Error())
	}
}

func TestGetPinAmountFromArgs02Error(t *testing.T) {
	args := []string{"a", "b", "c", "d"}
	_, err := getPinAmountFromArgs(args)

	if err == nil {
		t.Errorf("Expected error for invalid argument but no error was returned.")
	}
}
