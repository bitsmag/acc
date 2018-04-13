package useroutput

import (
	"acc/testutil"
	"acc/types"
	"reflect"
	"testing"
)

func TestCreateBalance01Success(t *testing.T) {
	balanceOutput := CreateBalance(1.23123)
	expectedBalanceOutput := types.Line{Line: "Balance is: 1.23", LineColor: "green"}

	if balanceOutput != expectedBalanceOutput {
		t.Errorf("useroutput.CreateBalance(amount), is: %v, want: %v.", balanceOutput, expectedBalanceOutput)
	}
}

func TestSeparateEntries01SingleGroupSuccess(t *testing.T) {
	entries := testutil.LogEntrySlice()
	groupedLog := separateEntries(entries, func(entry types.LogEntry) string {
		return "group"
	})

	keys := make([]string, len(groupedLog))
	i := 0
	for k := range groupedLog {
		keys[i] = k
		i++
	}

	if len(keys) != 1 {
		t.Errorf("useroutput.separateEntries(entries, getSeparator(types.LogEntry)) - wrong number of groups, is: %v, want: %v.", len(groupedLog), 1)
	}
	if keys[0] != "group" {
		t.Errorf("useroutput.separateEntries(entries, getSeparator(types.LogEntry)) - wrong group key, is: %v, want: %v.", keys[0], "group")
	}
}

func TestSortKeys01Success(t *testing.T) {
	keys := [5]string{"2017", "2003", "N/A", "321", "abc"}

	logEntrySlice := make([]types.LogEntry, 0)
	amount, date, category := testutil.LogEntryProperties(-10, "05.01.2018", "Ice-Cream")
	logEntrySlice = append(logEntrySlice, types.LogEntry{Amount: amount, Date: date, Category: category})

	var groupedLog = make(map[string][]types.LogEntry)
	for i := range keys {
		groupedLog[keys[i]] = logEntrySlice
	}

	expectedSortedKey := []string{"N/A", "abc", "321", "2003", "2017"}
	sortedKeys := sortKeys(groupedLog)

	if !reflect.DeepEqual(sortedKeys, expectedSortedKey) {
		t.Errorf("useroutput.sortKeys(groupedLog), is: %v, want: %v.", sortedKeys, expectedSortedKey)
	}
}

func TestCreteLog01Success(t *testing.T) {
	entries := testutil.LogEntrySlice()
	var order types.Order
	order.Set("Year")

	expectedPrintableLog := make([]types.Printable, 14)
	line := types.Line{Line: "======================================2018", LineColor: "white"}
	expectedPrintableLog = append(expectedPrintableLog, line)
	line = types.Line{Line: "05.01.2018    [pay-out]          -5.00", LineColor: "red"}
	expectedPrintableLog = append(expectedPrintableLog, line)
	line = types.Line{Line: "01.02.2018    [pay-out]        -500.00", LineColor: "red"}
	expectedPrintableLog = append(expectedPrintableLog, line)
	line = types.Line{Line: "01.02.2018    [pay-out]         -10.00", LineColor: "red"}
	expectedPrintableLog = append(expectedPrintableLog, line)
	line = types.Line{Line: "06.02.2018    [pay-in]         1500.00", LineColor: "green"}
	expectedPrintableLog = append(expectedPrintableLog, line)
	line = types.Line{Line: "06.03.2018    [pay-in]         1500.00", LineColor: "green"}
	expectedPrintableLog = append(expectedPrintableLog, line)
	line = types.Line{Line: "19.03.2018    [pay-out]         -10.00", LineColor: "red"}
	expectedPrintableLog = append(expectedPrintableLog, line)
	line = types.Line{Line: "======================================2019", LineColor: "white"}
	expectedPrintableLog = append(expectedPrintableLog, line)
	line = types.Line{Line: "26.06.2019    [pay-out]        -500.00", LineColor: "red"}
	expectedPrintableLog = append(expectedPrintableLog, line)
	line = types.Line{Line: "05.08.2019    [pay-out]        -650.00", LineColor: "red"}
	expectedPrintableLog = append(expectedPrintableLog, line)

	printableLog := CreateLog(entries, order)

	if reflect.DeepEqual(printableLog, expectedPrintableLog) {
		t.Errorf("useroutput.CreateLog(entries, order), is: %v, want: %v.", printableLog, expectedPrintableLog)
	}
}
