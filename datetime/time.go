package types

import "time"

const DateFormat = "02.01.2006"

type Period struct {
	Begin time.Time `json:"begin" yaml:"begin"`
	End   time.Time `json:"end" yaml:"end"`
}

func StartOfDay(t time.Time) time.Time {
	nt := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	return nt
}

func EndOfDay(t time.Time) time.Time {
	nt := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, time.UTC)
	return nt
}
