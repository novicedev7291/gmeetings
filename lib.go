package gmeetings

import "time"

type Type struct {
	_t string
}

var (
	DAILY   = Type{_t: "daily"}
	WEEKLY  = Type{_t: "weekly"}
	MONTHLY = Type{_t: "monthly"}
	YEARLY  = Type{_t: "yearly"}
)

type RangeType struct {
	_t string
}

var (
	NOEND    = RangeType{_t: "noEnd"}
	ENDDATE  = RangeType{_t: "endDate"}
	NUMBERED = RangeType{_t: "numbered"}
)

type Range struct {
	rangeType   RangeType
	startdate   time.Time
	enddate     time.Time
	occurrences int // i.e. numbered 10 occurrences
}
type Recurrence struct {
	patternType Type
	interval    int // i.e. in every two weeks, months, day(s)
	daysOfWeek  []string
	day         int    // i.e. day of month & year
	dayOfMonth  string // i.e. which day of month i.e. Tue, Mon, Weekday, Weekend
	dayOfYear   string
	index       int // i.e. First, Second, Third, Last day of month
}

type Meeting struct {
	title      string
	recurrence Recurrence
	ranges     Range
}

func PossibleNextDatesFor(meeting Meeting) []time.Time {
	var result []time.Time
	if meeting.recurrence.patternType == DAILY {
		end := meeting.ranges.enddate
		for start := meeting.ranges.startdate; !start.After(end); {
			result = append(result, start)
			start = start.AddDate(0, 0, 1)
		}
	}
	return result
}
