package leap

const testVersion = 2

// IsLeapYear returns true if the given year is a leap year, false otherwise.
func IsLeapYear(year int) bool {
	var isLeapYear bool

	if year%4 == 0 {
		if year%100 == 0 {
			isLeapYear = year%400 == 0
		} else {
			isLeapYear = true
		}
	}

	return isLeapYear
}
