package testutil

import "acc/types"

// LogEntrySlice provides a slice of logEntries
func LogEntrySlice() []types.LogEntry {
	slice := make([]types.LogEntry, 0)

	amount, date, category := LogEntryProperties(-5, "05.01.2018", "Ice-Cream")
	slice = append(slice, types.LogEntry{Amount: amount, Date: date, Category: category})

	amount, date, category = LogEntryProperties(-500, "01.02.2018", "Holiday")
	slice = append(slice, types.LogEntry{Amount: amount, Date: date, Category: category})

	amount, date, category = LogEntryProperties(-10, "01.02.2018", "Pizza")
	slice = append(slice, types.LogEntry{Amount: amount, Date: date, Category: category})

	amount, date, category = LogEntryProperties(1500, "06.02.2018", "Salary")
	slice = append(slice, types.LogEntry{Amount: amount, Date: date, Category: category})

	amount, date, category = LogEntryProperties(1500, "06.03.2018", "Salary")
	slice = append(slice, types.LogEntry{Amount: amount, Date: date, Category: category})

	amount, date, category = LogEntryProperties(-10, "19.03.2018", "Pizza")
	slice = append(slice, types.LogEntry{Amount: amount, Date: date, Category: category})

	amount, date, category = LogEntryProperties(-500, "26.06.2019", "Pizza")
	slice = append(slice, types.LogEntry{Amount: amount, Date: date, Category: category})

	amount, date, category = LogEntryProperties(-650, "05.08.2019", "Holiday")
	slice = append(slice, types.LogEntry{Amount: amount, Date: date, Category: category})

	return slice
}

// ContainsEntries checks wether the logEntries passed as second argument are contained in the logEntry slice passes as first argument
func ContainsEntries(entries []types.LogEntry, wantedEntries []types.LogEntry) bool {
	entriesFound := 0

	for _, wantedEntry := range wantedEntries {
		for _, entry := range entries {
			if wantedEntry.Amount == entry.Amount &&
				wantedEntry.Date == entry.Date &&
				wantedEntry.Category == entry.Category {
				entriesFound++
			}
		}
	}

	if len(wantedEntries) == entriesFound {
		return true
	}
	return false
}

// LogEntryProperties turns the passed arguments into objects which can be used to create a types.LogEntry
func LogEntryProperties(amountI int, dateS string, categoryS string) (float64, types.Date, types.Category) {
	amount := float64(amountI)
	date := types.Date{}
	date.Set(dateS)
	category := types.Category{}
	category.Set(categoryS)

	return amount, date, category
}
