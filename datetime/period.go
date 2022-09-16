package datetime

import "time"

type Period struct {
	Begin time.Time `json:"begin" yaml:"begin"`
	End   time.Time `json:"end" yaml:"end"`
}

func NewPeriod(a, b time.Time) Period {
	return Period{Begin: StartOfDay(a), End: EndOfDay(b)}
}

func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
}

func MonthOf(t time.Time) Period {
	var (
		l     = t.Location()
		begin = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, l)
		d     = begin.AddDate(0, 1, -1)
		end   = time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, 0, l)
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
