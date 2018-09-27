// Clock stub file

package clock

import (
	"fmt"
	"time"
)

const testVersion = 4

// Clock represents a particular moment in time.
type Clock struct {
	currentTime time.Time
}

// New creates a new Clock
func New(hour, minute int) Clock {
	hour, minute = normalizeTime(hour, minute)
	t, _ := time.Parse("15:04", fmt.Sprintf("%02d:%02d", hour, minute))
	return Clock{currentTime: t}
}

func normalizeTime(hour, minute int) (int, int) {
	switch {
	case hour > 23:
		hour = hour % 24
	case hour < 0:
		hour = 24 - (-1*hour)%24
	}

	switch {
	case minute > 59:
		additionalHours := minute / 60
		hour += additionalHours % 24
		minute = minute % 60
	case minute < 0:
		hoursAgo := (-1*minute)/60 + 1
		hour -= hoursAgo

		if hour < 0 {
			hour = 24 - (-1*hour)%24
		}

		minute = 60 - (-1*minute)%60
	}

	return hour, minute
}

func (c Clock) String() string {
	return c.currentTime.Format("15:04")
}

// Add increments the current time of the clock by the given number of minutes.
func (c Clock) Add(minutes int) Clock {
	return Clock{currentTime: c.currentTime.Add(time.Duration(minutes) * time.Minute)}
}
