package datetime

import (
	"errors"
	"time"
)

type Period struct {
	Begin time.Time `json:"begin" yaml:"begin"`
	End   time.Time `json:"end" yaml:"end"`
}

var ErrBeforeAfterEnd = errors.New("before_cannot_be_after_end")

func NewPeriod(begin, end time.Time) (Period, error) {
	switch {
	case begin.IsZero() && !end.IsZero():
		return Period{End: EndOfDay(end)}, nil
	case !begin.IsZero() && end.IsZero():
		return Period{Begin: StartOfDay(begin)}, nil
	case begin.IsZero() && end.IsZero():
		return Period{}, nil
	case begin.After(end):
		return Period{}, ErrBeforeAfterEnd
	}

	begin = StartOfDay(begin)
	end = EndOfDay(end)

	return Period{Begin: begin, End: end}, nil
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

func (p Period) Contains(t time.Time) bool {
	return (p.Begin.Equal(t) || p.Begin.Before(t)) &&
		(p.End.Equal(t) || p.End.After(t))
}

func MustPeriod(p Period, err error) Period {
	if err != nil {
		panic(err)
	}

	return p
}
