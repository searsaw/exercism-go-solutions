package gigasecond

import "time"

// TestVersion is used by the testing framework.
const TestVersion = 2

// Gigasecond in a Duration type
var Gigasecond = 1e9 * time.Second

// My Birthday
var Birthday = time.Date(1989, time.May, 22, 0, 0, 0, 0, time.Local)

// AddGigasecond adds a Gigasecond (10**9 seconds) to the given date
func AddGigasecond(date time.Time) time.Time {
	return date.Add(Gigasecond)
}
