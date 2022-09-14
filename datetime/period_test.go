package datetime_test

import (
	"testing"
	"time"

	"github.com/blue-health/blue-go-toolbox/datetime"
	"github.com/stretchr/testify/require"
)

func TestPreviousMonthOf(t *testing.T) {
	testCases := []struct {
		in  time.Time
		out datetime.Period
	}{
		{in: time.Date(2022, time.July, 1, 0, 0, 0, 0, time.UTC), out: datetime.Period{
			Begin: time.Date(2022, time.June, 1, 0, 0, 0, 0, time.UTC),
			End:   time.Date(2022, time.June, 30, 23, 59, 59, 999999999, time.UTC),
		}},
		{in: time.Date(2022, time.August, 1, 0, 0, 0, 0, time.UTC), out: datetime.Period{
			Begin: time.Date(2022, time.July, 1, 0, 0, 0, 0, time.UTC),
			End:   time.Date(2022, time.July, 31, 23, 59, 59, 999999999, time.UTC),
		}},
		{in: time.Date(2022, time.September, 3, 0, 0, 0, 0, time.UTC), out: datetime.Period{
			Begin: time.Date(2022, time.August, 1, 0, 0, 0, 0, time.UTC),
			End:   time.Date(2022, time.August, 31, 23, 59, 59, 999999999, time.UTC),
		}},
		{in: time.Date(2022, time.September, 15, 0, 0, 0, 0, time.UTC), out: datetime.Period{
			Begin: time.Date(2022, time.August, 1, 0, 0, 0, 0, time.UTC),
			End:   time.Date(2022, time.August, 31, 23, 59, 59, 999999999, time.UTC),
		}},
		{in: time.Date(2022, time.September, 30, 0, 0, 0, 0, time.UTC), out: datetime.Period{
			Begin: time.Date(2022, time.August, 1, 0, 0, 0, 0, time.UTC),
			End:   time.Date(2022, time.August, 31, 23, 59, 59, 999999999, time.UTC),
		}},
		{in: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC), out: datetime.Period{
			Begin: time.Date(2022, time.December, 1, 0, 0, 0, 0, time.UTC),
			End:   time.Date(2022, time.December, 31, 23, 59, 59, 999999999, time.UTC),
		}},
	}

	for _, c := range testCases {
		t.Run(c.in.Format(time.RFC1123), func(t *testing.T) {
			require.Equal(t, c.out, datetime.PreviousMonthOf(c.in))
		})
	}
}
