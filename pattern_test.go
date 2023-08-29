package gmeetings

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type Date struct {
	y int
	m int
	d int
}

func newDate(y int, m int, d int) Date {
	return Date{y, m, d}
}

func utcTime(year int, month int, day int, hh int, mm int, ss int) time.Time {
	zone, err := time.LoadLocation("UTC")
	if err != nil {
		fmt.Printf("Failed to create datetime with (%d, %d, %d, %d, %d, %d)", year, month, day, hh, mm, ss)
		return time.Time{}
	}
	return time.Date(year, time.Month(month), day, hh, mm, ss, 0, zone)
}

func utcDate(year int, month int, day int) time.Time {
	return utcTime(year, month, day, 0, 0, 0)
}

func repeatDailyScenario(startdate Date, enddate Date) Meeting {
	start := utcDate(startdate.y, startdate.m, startdate.d)
	end := utcDate(enddate.y, enddate.m, enddate.d)
	return Meeting{
		title: "Repeat Daily till end date",
		recurrence: Recurrence{
			patternType: DAILY,
		},
		ranges: Range{
			startdate: start,
			enddate:   end,
			rangeType: ENDDATE,
		},
	}
}

func repeatEveryWeekSundayScenario(startdate Date, enddate Date) Meeting {
	start := utcDate(startdate.y, startdate.m, startdate.d)
	end := utcDate(enddate.y, enddate.m, enddate.d)
	return Meeting{
		title: "Repeat every week on sunday till end date",
		recurrence: Recurrence{
			patternType: WEEKLY,
			interval:    1,
			daysOfWeek:  []string{"Sunday"},
		},
		ranges: Range{
			startdate: start,
			enddate:   end,
			rangeType: ENDDATE,
		},
	}

}

func TestRecurrenceEachDay(t *testing.T) {
	meeting := repeatDailyScenario(newDate(2023, 8, 1), newDate(2023, 8, 3))
	allPossibleDates := NextDatesFor(meeting)
	expected := []time.Time{
		utcDate(2023, 8, 1),
		utcDate(2023, 8, 2),
		utcDate(2023, 8, 3),
	}

	if !reflect.DeepEqual(expected, allPossibleDates) {
		t.Errorf("Failed to compare %q == %q", expected, allPossibleDates)
	}
}

func TestRecurrenceWeeklyOnSundy(t *testing.T) {
	meeting := repeatEveryWeekSundayScenario(newDate(2023, 8, 1), newDate(2023, 8, 31))
	allPossibleDates := NextDatesFor(meeting)
	expected := []time.Time{
		utcDate(2023, 8, 6),
		utcDate(2023, 8, 13),
		utcDate(2023, 8, 20),
		utcDate(2023, 8, 27),
	}

	if !reflect.DeepEqual(expected, allPossibleDates) {
		t.Errorf("Failed to compare %q == %q", expected, allPossibleDates)
	}
}
