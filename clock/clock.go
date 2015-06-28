// Package clock implements a clock that represents a point in time.
// It allows addition (and subtraction) to calulcate time.
package clock

import (
	"fmt"
)

// TestVersion is used by the testing framework.
const TestVersion = 2

// 1440 = 24 * 60 = minutes in a day
const MinsInDay = 1440

const MinInHour = 60

// Clock represents a point in time.
type Clock struct {
	total int
}

// Time takes an hour and minute and returns a Clock object
func Time(hour, minute int) Clock {
	total := hour*MinInHour + minute

	for total >= MinsInDay || total < 0 {
		if total >= MinsInDay {
			total -= MinsInDay
		} else if total < 0 {
			total += MinsInDay
		}
	}

	return Clock{total}
}

// String returns the string representation of the clock's time
func (c Clock) String() string {
	hour := c.total / 60
	minute := c.total % 60
	return fmt.Sprintf("%02d:%02d", hour, minute)
}

// Add takes minutes (positive or negative) and adds them to the
// current time.  If it is negative, the time is decreased.
func (c Clock) Add(minutes int) Clock {
	total := c.total + minutes
	return Time(0, total)
}
