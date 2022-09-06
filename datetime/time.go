package types

import "time"

type Period struct {
	Begin, End time.Time
}

func StartOfDay(t time.Time) time.Time {
	nt := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	return nt
}

func EndOfDay(t time.Time) time.Time {
	nt := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, time.UTC)
	return nt
}

func MonthOf(t time.Time) Period {
	var (
		begin    = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)
		finalDay = time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
		end      = time.Date(t.Year(), t.Month(), finalDay, 23, 59, 59, 0, time.UTC)
	)

	return Period{Begin: begin, End: end}
}
