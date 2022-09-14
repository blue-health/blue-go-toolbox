package datetime

import "time"

type Period struct {
	Begin time.Time `json:"start" yaml:"start"`
	End   time.Time `json:"end" yaml:"end"`
}

const maxNsec = 999999999

func NewPeriod(a, b time.Time) Period {
	return Period{Begin: StartOfDay(a), End: EndOfDay(b)}
}

func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}

func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, maxNsec, time.UTC)
}

func MonthOf(t time.Time) Period {
	var (
		begin    = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)
		finalDay = time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
		end      = time.Date(t.Year(), t.Month(), finalDay, 23, 59, 59, maxNsec, time.UTC)
	)

	return Period{Begin: begin, End: end}
}

func PreviousMonthOf(t time.Time) Period {
	return MonthOf(t.AddDate(0, -1, 0))
}

func (p Period) Within(o Period) bool {
	return (p.Begin.Equal(o.Begin) || p.Begin.After(o.Begin)) &&
		(p.End.Equal(o.End) || p.End.Before(o.End))
}
