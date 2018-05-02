package cmd

import (
	"acc/testutil"
	"acc/types"
	"testing"
)

func TestFilterEntries01DateDayAccuracySuccess(t *testing.T) {
	entries := testutil.LogEntrySlice()

	filterDate := types.Date{}
	filterDate.Set("01.02.2018")
	filterCategory := types.Category{}
	filterCategory.Set("")

	filterResults := filterEntries(entries, filterDate, filterCategory)

	expectedResults := make([]types.LogEntry, 0)
	amount, date, category := testutil.LogEntryProperties(-500, "01.02.2018", "Holiday")
	expectedResults = append(expectedResults, types.LogEntry{Amount: amount, Date: date, Category: category})
	amount, date, category = testutil.LogEntryProperties(-10, "01.02.2018", "Pizza")
	expectedResults = append(expectedResults, types.LogEntry{Amount: amount, Date: date, Category: category})

	if len(filterResults) != 2 {
		t.Errorf("cmd.log.filterEntries(entries, filterDate, filterCategory) - unexpected result length, got: %v, want: %v.", len(filterResults), 2)
	}
	if !testutil.ContainsEntries(filterResults, expectedResults) {
		t.Errorf("cmd.log.filterEntries(entries, filterDate, filterCategory) - result didn't contain the expected entries")
	}
}

func TestFilterEntries02DateMonthAccuracySuccess(t *testing.T) {
	entries := testutil.LogEntrySlice()

	filterDate := types.Date{}
	filterDate.Set("03.2018")
	filterCategory := types.Category{}
	filterCategory.Set("")

	filterResults := filterEntries(entries, filterDate, filterCategory)

	expectedResults := make([]types.LogEntry, 0)
	amount, date, category := testutil.LogEntryProperties(1500, "06.03.2018", "Salary")
	expectedResults = append(expectedResults, types.LogEntry{Amount: amount, Date: date, Category: category})
	amount, date, category = testutil.LogEntryProperties(-10, "19.03.2018", "Pizza")
	expectedResults = append(expectedResults, types.LogEntry{Amount: amount, Date: date, Category: category})

	if len(filterResults) != 2 {
		t.Errorf("cmd.log.filterEntries(entries, filterDate, filterCategory) - unexpected result length, got: %v, want: %v.", len(filterResults), 2)
	}
	if !testutil.ContainsEntries(filterResults, expectedResults) {
		t.Errorf("cmd.log.filterEntries(entries, filterDate, filterCategory) - result didn't contain the expected entries")
	}
}

func TestFilterEntries03DateYearAccuracySuccess(t *testing.T) {
	entries := testutil.LogEntrySlice()

	filterDate := types.Date{}
	filterDate.Set("2018")
	filterCategory := types.Category{}
	filterCategory.Set("")

	filterResults := filterEntries(entries, filterDate, filterCategory)

	expectedResults := make([]types.LogEntry, 0)
	amount, date, category := testutil.LogEntryProperties(-5, "05.01.2018", "Ice-Cream")
	expectedResults = append(expectedResults, types.LogEntry{Amount: amount, Date: date, Category: category})
	amount, date, category = testutil.LogEntryProperties(-500, "01.02.2018", "Holiday")
	expectedResults = append(expectedResults, types.LogEntry{Amount: amount, Date: date, Category: category})
	amount, date, category = testutil.LogEntryProperties(-10, "01.02.2018", "Pizza")
	expectedResults = append(expectedResults, types.LogEntry{Amount: amount, Date: date, Category: category})
	amount, date, category = testutil.LogEntryProperties(1500, "06.02.2018", "Salary")
	expectedResults = append(expectedResults, types.LogEntry{Amount: amount, Date: date, Category: category})
	amount, date, category = testutil.LogEntryProperties(1500, "06.03.2018", "Salary")
	expectedResults = append(expectedResults, types.LogEntry{Amount: amount, Date: date, Category: category})
	amount, date, category = testutil.LogEntryProperties(-10, "19.03.2018", "Pizza")
	expectedResults = append(expectedResults, types.LogEntry{Amount: amount, Date: date, Category: category})

	if len(filterResults) != 6 {
		t.Errorf("cmd.log.filterEntries(entries, filterDate, filterCategory) - unexpected result length, got: %v, want: %v.", len(filterResults), 6)
	}
	if !testutil.ContainsEntries(filterResults, expectedResults) {
		t.Errorf("cmd.log.filterEntries(entries, filterDate, filterCategory) - result didn't contain the expected entries")
	}
}

func TestFilterEntries03CategorySuccess(t *testing.T) {
	entries := testutil.LogEntrySlice()

	filterDate := types.Date{}
	filterDate.Set("")
	filterCategory := types.Category{}
	filterCategory.Set("Pizza")

	filterResults := filterEntries(entries, filterDate, filterCategory)

	expectedResults := make([]types.LogEntry, 0)
	amount, date, category := testutil.LogEntryProperties(-10, "01.02.2018", "Pizza")
	expectedResults = append(expectedResults, types.LogEntry{Amount: amount, Date: date, Category: category})
	amount, date, category = testutil.LogEntryProperties(-10, "19.03.2018", "Pizza")
	expectedResults = append(expectedResults, types.LogEntry{Amount: amount, Date: date, Category: category})
	amount, date, category = testutil.LogEntryProperties(-500, "26.06.2019", "Pizza")
	expectedResults = append(expectedResults, types.LogEntry{Amount: amount, Date: date, Category: category})

	if len(filterResults) != 3 {
		t.Errorf("cmd.log.filterEntries(entries, filterDate, filterCategory) - unexpected result length, got: %v, want: %v.", len(filterResults), 3)
	}
	if !testutil.ContainsEntries(filterResults, expectedResults) {
		t.Errorf("cmd.log.filterEntries(entries, filterDate, filterCategory) - result didn't contain the expected entries")
	}
}

func TestFilterEntries04CategoryAndDateYearAccuracySuccess(t *testing.T) {
	entries := testutil.LogEntrySlice()

	filterDate := types.Date{}
	filterDate.Set("2019")
	filterCategory := types.Category{}
	filterCategory.Set("Pizza")

	filterResults := filterEntries(entries, filterDate, filterCategory)

	expectedResults := make([]types.LogEntry, 0)
	amount, date, category := testutil.LogEntryProperties(-500, "26.06.2019", "Pizza")
	expectedResults = append(expectedResults, types.LogEntry{Amount: amount, Date: date, Category: category})

	if len(filterResults) != 1 {
		t.Errorf("cmd.log.filterEntries(entries, filterDate, filterCategory) - unexpected result length, got: %v, want: %v.", len(filterResults), 1)
	}
	if !testutil.ContainsEntries(filterResults, expectedResults) {
		t.Errorf("cmd.log.filterEntries(entries, filterDate, filterCategory) - result didn't contain the expected entries")
	}
}

func TestFilterEntries06NonExistingCategorySuccess(t *testing.T) {
	entries := testutil.LogEntrySlice()

	filterDate := types.Date{}
	filterDate.Set("")
	filterCategory := types.Category{}
	filterCategory.Set("NonExistingCategory")

	filterResults := filterEntries(entries, filterDate, filterCategory)

	if len(filterResults) != 0 {
		t.Errorf("cmd.log.filterEntries(entries, filterDate, filterCategory) - unexpected result length, got: %v, want: %v.", len(filterResults), 2)
	}
}

func TestFilterEntries07WithoutFilterArgumentsSuccess(t *testing.T) {
	entries := testutil.LogEntrySlice()

	filterDate := types.Date{}
	filterDate.Set("")
	filterCategory := types.Category{}
	filterCategory.Set("")

	filterResults := filterEntries(entries, filterDate, filterCategory)

	expectedResults := entries

	if len(filterResults) != len(entries) {
		t.Errorf("cmd.log.filterEntries(entries, filterDate, filterCategory) - unexpected result length, got: %v, want: %v.", len(filterResults), len(entries))
	}
	if !testutil.ContainsEntries(filterResults, expectedResults) {
		t.Errorf("cmd.log.filterEntries(entries, filterDate, filterCategory) - result didn't contain the expected entries")
	}
}
