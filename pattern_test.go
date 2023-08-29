package gmeetings

import (
	"testing"
	"time"
)

func utcTime(year int, month int, day int, hh int, mm int, ss int) (time.Time, error) {
	zone, err := time.LoadLocation("UTC")
	if err != nil {
		return time.Time{}, err
	}
	return time.Date(year, time.Month(month), day, hh, mm, ss, 0, zone), nil
}

func utcDate(year int, month int, day int) (time.Time, error) {
	return utcTime(year, month, day, 0, 0, 0)
}

func repeatDailyScenario(startdate time.Time, enddate time.Time) (Meeting, error) {

	return Meeting{
		title: "Repeat Daily till end date",
		recurrence: Recurrence{
			patternType: DAILY,
		},
		ranges: Range{
			startdate: startdate,
			enddate:   enddate,
			rangeType: ENDDATE,
		},
	}, nil
}

func TestRecurrenceEachDay(t *testing.T) {
	startdate, err := utcTime(2023, 8, 1, 0, 0, 0)
	if err != nil {
		t.Errorf("Failed to create startdate with err %q", err)
		return
	}
	enddate, err := utcTime(2023, 8, 3, 0, 0, 0)
	if err != nil {
		t.Errorf("Failed to create enddate with err %q", err)
		return
	}

	meeting, err := repeatDailyScenario(startdate, enddate)
	if err != nil {
		t.Errorf("Failed to create meeting with %q, %q", startdate, enddate)
		return
	}

	allPossibleDates := PossibleNextDatesFor(meeting)

	var expected [3]time.Time

	for i := 1; i <= 3; i++ {
		date, err := utcDate(2023, 8, i)
		if err != nil {
			t.Errorf("Failed to create expected array with error %q", err)
			return
		}
		expected[i-1] = date
	}

	for _, expect := range expected {
		matches := false
		for _, actual := range allPossibleDates {
			if actual == expect {
				matches = true
				break
			}
		}

		if !matches {
			t.Errorf("Failed to find %q into %q", expect, allPossibleDates)
			t.FailNow()
			return
		}
	}
}
