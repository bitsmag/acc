package types

import (
	"testing"
	"time"
)

func TestSet01DayAccuracySuccess(t *testing.T) {
	var date Date
	err := date.Set("31.03.1984")

	expectedTimeValue := time.Date(1984, time.Month(3), 31, int(0), int(0), int(0), int(0), time.UTC)

	if !date.Time.Equal(expectedTimeValue) {
		t.Errorf("types.Date.Set(), is: %v, want: %v.", date.Time, expectedTimeValue)
	}
	if err != nil {
		t.Errorf("Error occured while attempting to set the date: %v.", err.Error())
	}
}

func TestSet02MonthAccuracySuccess(t *testing.T) {
	var date Date
	err := date.Set("28.2.9")

	expectedTimeValue := time.Date(9, time.Month(3), 0, int(0), int(0), int(0), int(0), time.UTC)

	if !date.Time.Equal(expectedTimeValue) {
		t.Errorf("types.Date.Set(), is: %v, want: %v.", date.Time, expectedTimeValue)
	}
	if err != nil {
		t.Errorf("Error occured while attempting to set the date: %v.", err.Error())
	}
}

func TestSet03AutoIncreasingDateSuccess(t *testing.T) {
	var date Date
	err := date.Set("12.15.2012")

	expectedTimeValue := time.Date(2013, time.Month(3), 12, int(0), int(0), int(0), int(0), time.UTC)

	if !date.Time.Equal(expectedTimeValue) {
		t.Errorf("types.Date.Set(), is: %v, want: %v.", date.Time, expectedTimeValue)
	}
	if err != nil {
		t.Errorf("Error occured while attempting to set the date: %v.", err.Error())
	}
}

func TestSet04NegativeDayError(t *testing.T) {
	var date Date
	err := date.Set("-2.10.2009")

	if err == nil {
		t.Errorf("Expected error for invalid argument but no error was returned.")
	}
}

func TestSet05ZeroMonthError(t *testing.T) {
	var date Date
	err := date.Set("2.0.2009")

	if err == nil {
		t.Errorf("Expected error for invalid argument but no error was returned.")
	}
}

func TestSet06WrongFormatError(t *testing.T) {
	var date Date
	err := date.Set("2.0..2009")

	if err == nil {
		t.Errorf("Expected error for invalid argument but no error was returned.")
	}
}

func TestString01DayAccuracySuccess(t *testing.T) {
	var date Date
	timeValue := time.Date(1984, time.Month(3), 31, int(0), int(0), int(0), int(0), time.UTC)
	date.Time = timeValue
	date.Accuracy = "day"

	expectedDateString := "31.03.1984"
	dateString := date.String()

	if dateString != expectedDateString {
		t.Errorf("type.Date.String(), is: %v, want: %v.", dateString, expectedDateString)
	}
}

func TestString02YearAccuracySuccess(t *testing.T) {
	var date Date
	timeValue := time.Date(3, time.Month(1), 1, int(0), int(0), int(0), int(0), time.UTC)
	date.Time = timeValue
	date.Accuracy = "year"

	expectedDateString := "......0003"
	dateString := date.String()

	if dateString != expectedDateString {
		t.Errorf("type.Date.String(), is: %v, want: %v.", dateString, expectedDateString)
	}
}

func TestFrom01YearAccuracySuccess(t *testing.T) {
	var date Date
	date.Set("2013")

	expectedFrom := time.Date(2013, time.Month(1), 1, int(0), int(0), int(0), int(0), time.UTC)
	from := date.From()

	if from != expectedFrom {
		t.Errorf("type.Date.From(), is: %v, want: %v.", from, expectedFrom)
	}
}

func TestTo01DayAccuracySuccess(t *testing.T) {
	var date Date
	date.Set("5.4.2018")

	expectedTo := time.Date(2018, time.Month(4), 5, int(23), int(59), int(59), int(999999999), time.UTC)
	from := date.To()

	if from != expectedTo {
		t.Errorf("type.Date.To(), is: %v, want: %v.", from, expectedTo)
	}
}

func TestTo02MonthAccuracySuccess(t *testing.T) {
	var date Date
	date.Set("4.2018")

	expectedTo := time.Date(2018, time.Month(4), 30, int(23), int(59), int(59), int(999999999), time.UTC)
	to := date.To()

	if to != expectedTo {
		t.Errorf("type.Date.To(), is: %v, want: %v.", to, expectedTo)
	}
}

func TestTo03MonthAccuracySuccess(t *testing.T) {
	var date Date
	date.Set("2.2016")

	expectedTo := time.Date(2016, time.Month(2), 29, int(23), int(59), int(59), int(999999999), time.UTC)
	to := date.To()

	if to != expectedTo {
		t.Errorf("type.Date.To(), is: %v, want: %v.", to, expectedTo)
	}
}

func TestTo04YearAccuracySuccess(t *testing.T) {
	var date Date
	date.Set("2016")

	expectedTo := time.Date(2016, time.Month(12), 31, int(23), int(59), int(59), int(999999999), time.UTC)
	to := date.To()

	if to != expectedTo {
		t.Errorf("type.Date.To(), is: %v, want: %v.", to, expectedTo)
	}
}
