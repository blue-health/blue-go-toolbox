package datetime_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/blue-health/blue-go-toolbox/datetime"
	"github.com/stretchr/testify/require"
)

func TestNewPeriod(t *testing.T) {
	var (
		now      = time.Now()
		dayBegin = datetime.StartOfDay(now)
		dayEnd   = datetime.EndOfDay(now)
	)

	testCases := []struct {
		begin, end time.Time
		result     datetime.Period
		err        bool
	}{
		{},
		{
			begin: now,
			result: datetime.Period{
				Begin: dayBegin,
			},
		},
		{
			end: now,
			result: datetime.Period{
				End: dayEnd,
			},
		},
		{
			begin: now,
			end:   now,
			result: datetime.Period{
				Begin: dayBegin,
				End:   dayEnd,
			},
		},
		{
			begin: now.Add(time.Hour),
			end:   now,
			err:   true,
		},
		{
			begin: now,
			end:   now.Add(-time.Hour),
			err:   true,
		},
	}

	for _, c := range testCases {
		t.Run(fmt.Sprintf("%s %s", c.begin.String(), c.end.String()), func(t *testing.T) {
			p, err := datetime.NewPeriod(c.begin, c.end)
			if c.err {
				require.NotNil(t, err)
			} else {
				require.Nil(t, err)
				require.Equal(t, c.result, p)
			}
		})
	}
}

func TestMonthOf(t *testing.T) {
	testCases := []struct {
		in  time.Time
		out datetime.Period
	}{
		{in: time.Date(2022, time.July, 1, 0, 0, 0, 0, time.UTC), out: datetime.Period{
			Begin: time.Date(2022, time.July, 1, 0, 0, 0, 0, time.UTC),
			End:   time.Date(2022, time.July, 31, 23, 59, 59, 0, time.UTC),
		}},
		{in: time.Date(2022, time.August, 1, 0, 0, 0, 0, time.UTC), out: datetime.Period{
			Begin: time.Date(2022, time.August, 1, 0, 0, 0, 0, time.UTC),
			End:   time.Date(2022, time.August, 31, 23, 59, 59, 0, time.UTC),
		}},
		{in: time.Date(2022, time.September, 3, 0, 0, 0, 0, time.UTC), out: datetime.Period{
			Begin: time.Date(2022, time.September, 1, 0, 0, 0, 0, time.UTC),
			End:   time.Date(2022, time.September, 30, 23, 59, 59, 0, time.UTC),
		}},
		{in: time.Date(2022, time.September, 15, 0, 0, 0, 0, time.UTC), out: datetime.Period{
			Begin: time.Date(2022, time.September, 1, 0, 0, 0, 0, time.UTC),
			End:   time.Date(2022, time.September, 30, 23, 59, 59, 0, time.UTC),
		}},
		{in: time.Date(2022, time.September, 30, 0, 0, 0, 0, time.UTC), out: datetime.Period{
			Begin: time.Date(2022, time.September, 1, 0, 0, 0, 0, time.UTC),
			End:   time.Date(2022, time.September, 30, 23, 59, 59, 0, time.UTC),
		}},
		{in: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC), out: datetime.Period{
			Begin: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
			End:   time.Date(2023, time.January, 31, 23, 59, 59, 0, time.UTC),
		}},
	}

	for _, c := range testCases {
		t.Run(c.in.Format(time.RFC1123), func(t *testing.T) {
			require.Equal(t, c.out, datetime.MonthOf(c.in))
		})
	}
}

func TestPreviousMonthOf(t *testing.T) {
	testCases := []struct {
		in  time.Time
		out datetime.Period
	}{
		{in: time.Date(2022, time.July, 1, 0, 0, 0, 0, time.UTC), out: datetime.Period{
			Begin: time.Date(2022, time.June, 1, 0, 0, 0, 0, time.UTC),
			End:   time.Date(2022, time.June, 30, 23, 59, 59, 0, time.UTC),
		}},
		{in: time.Date(2022, time.August, 1, 0, 0, 0, 0, time.UTC), out: datetime.Period{
			Begin: time.Date(2022, time.July, 1, 0, 0, 0, 0, time.UTC),
			End:   time.Date(2022, time.July, 31, 23, 59, 59, 0, time.UTC),
		}},
		{in: time.Date(2022, time.September, 3, 0, 0, 0, 0, time.UTC), out: datetime.Period{
			Begin: time.Date(2022, time.August, 1, 0, 0, 0, 0, time.UTC),
			End:   time.Date(2022, time.August, 31, 23, 59, 59, 0, time.UTC),
		}},
		{in: time.Date(2022, time.September, 15, 0, 0, 0, 0, time.UTC), out: datetime.Period{
			Begin: time.Date(2022, time.August, 1, 0, 0, 0, 0, time.UTC),
			End:   time.Date(2022, time.August, 31, 23, 59, 59, 0, time.UTC),
		}},
		{in: time.Date(2022, time.September, 30, 0, 0, 0, 0, time.UTC), out: datetime.Period{
			Begin: time.Date(2022, time.August, 1, 0, 0, 0, 0, time.UTC),
			End:   time.Date(2022, time.August, 31, 23, 59, 59, 0, time.UTC),
		}},
		{in: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC), out: datetime.Period{
			Begin: time.Date(2022, time.December, 1, 0, 0, 0, 0, time.UTC),
			End:   time.Date(2022, time.December, 31, 23, 59, 59, 0, time.UTC),
		}},
	}

	for _, c := range testCases {
		t.Run(c.in.Format(time.RFC1123), func(t *testing.T) {
			require.Equal(t, c.out, datetime.PreviousMonthOf(c.in))
		})
	}
}

func TestWithin(t *testing.T) {
	testCases := []struct {
		name       string
		sub, super datetime.Period
		ret        bool
	}{
		{
			name: "subset is zero to infinity",
			sub: datetime.Period{
				Begin: time.Time{},
				End:   time.Now().Add(10e15),
			},
			super: datetime.Period{
				Begin: time.Now(),
				End:   time.Now(),
			},
			ret: false,
		},
		{
			name: "subset is zero to zero",
			sub: datetime.Period{
				Begin: time.Time{},
				End:   time.Time{},
			},
			super: datetime.Period{
				Begin: time.Now(),
				End:   time.Now(),
			},
			ret: false,
		},
		{
			name: "subset is infinity to infinity",
			sub: datetime.Period{
				Begin: time.Now().Add(10e15),
				End:   time.Now().Add(10e15),
			},
			super: datetime.Period{
				Begin: time.Now(),
				End:   time.Now(),
			},
			ret: false,
		},
		{
			name: "both zero",
			sub: datetime.Period{
				Begin: time.Time{},
				End:   time.Time{},
			},
			super: datetime.Period{
				Begin: time.Time{},
				End:   time.Time{},
			},
			ret: true,
		},
		{
			name: "both infinity",
			sub: datetime.Period{
				Begin: time.Time{}.Add(10e15),
				End:   time.Time{}.Add(10e15),
			},
			super: datetime.Period{
				Begin: time.Time{}.Add(10e15),
				End:   time.Time{}.Add(10e15),
			},
			ret: true,
		},
		{
			name: "superset is zero to infinity",
			sub: datetime.Period{
				Begin: time.Now(),
				End:   time.Now(),
			},
			super: datetime.Period{
				Begin: time.Time{},
				End:   time.Now().Add(10e17),
			},
			ret: true,
		},
		{
			name: "superset is zero to zero",
			sub: datetime.Period{
				Begin: time.Now(),
				End:   time.Now(),
			},
			super: datetime.Period{
				Begin: time.Time{},
				End:   time.Time{},
			},
			ret: false,
		},
		{
			name: "superset is infinity to infinity",
			sub: datetime.Period{
				Begin: time.Now(),
				End:   time.Now(),
			},
			super: datetime.Period{
				Begin: time.Now().Add(10e15),
				End:   time.Now().Add(10e15),
			},
			ret: false,
		},
		{
			name: "subset is backwards infinity to zero",
			sub: datetime.Period{
				Begin: time.Now().Add(10e15),
				End:   time.Time{},
			},
			super: datetime.Period{
				Begin: time.Now(),
				End:   time.Now(),
			},
			ret: true,
		},
		{
			name: "subset is normal subset",
			sub: datetime.Period{
				Begin: time.Now(),
				End:   time.Now(),
			},
			super: datetime.Period{
				Begin: time.Now().Add(-time.Hour),
				End:   time.Now().Add(time.Hour),
			},
			ret: true,
		},
		{
			name: "subset is normal subset",
			sub: datetime.Period{
				Begin: time.Now(),
				End:   time.Now(),
			},
			super: datetime.Period{
				Begin: time.Now().Add(-time.Hour),
				End:   time.Now().Add(time.Hour),
			},
			ret: true,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.ret, c.sub.Within(c.super))
		})
	}
}

func TestContainsDate(t *testing.T) {
	testCases := []struct {
		name string
		p    datetime.Period
		t    time.Time
		ret  bool
	}{
		{
			name: "zero period and zero time",
			p:    datetime.Period{},
			ret:  true,
		},
		{
			name: "now period and zero time",
			p: datetime.Period{
				Begin: time.Now(),
				End:   time.Now(),
			},
			ret: false,
		},
		{
			name: "zero period and now time",
			p:    datetime.Period{},
			t:    time.Now(),
			ret:  false,
		},
		{
			name: "zero to infinite period and now time",
			p: datetime.Period{
				End: time.Now().Add(10000 * time.Minute),
			},
			t:   time.Now(),
			ret: true,
		},
		{
			name: "zero to infinite period and zero time",
			p: datetime.Period{
				End: time.Now().Add(10000 * time.Minute),
			},
			ret: true,
		},
		{
			name: "zero to infinite period and now time",
			p: datetime.Period{
				End: time.Now().Add(10000 * time.Minute),
			},
			t:   time.Now(),
			ret: true,
		},
		{
			name: "1 hour period and infinite time",
			p: datetime.Period{
				Begin: time.Now(),
				End:   time.Now().Add(time.Hour),
			},
			t:   time.Now().Add(10000 * time.Minute),
			ret: false,
		},
		{
			name: "1 hour period and now time",
			p: datetime.Period{
				Begin: time.Now(),
				End:   time.Now().Add(time.Hour),
			},
			t:   time.Now(),
			ret: true,
		},
		{
			name: "2 hour period and infinite time",
			p: datetime.Period{
				Begin: time.Now().Add(-time.Hour),
				End:   time.Now().Add(time.Hour),
			},
			t:   time.Now().Add(10000 * time.Minute),
			ret: false,
		},
		{
			name: "2 hour period and now time",
			p: datetime.Period{
				Begin: time.Now().Add(-time.Hour),
				End:   time.Now().Add(time.Hour),
			},
			t:   time.Now(),
			ret: true,
		},
		{
			name: "begin in 1 hour period and now time",
			p: datetime.Period{
				Begin: time.Now().Add(time.Hour),
				End:   time.Now().Add(time.Hour),
			},
			t:   time.Now(),
			ret: false,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.ret, c.p.Contains(c.t))
		})
	}
}
