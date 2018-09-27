package hamming

import "errors"

const testVersion = 5

// Distance calculates the Hemming distance between DNA strands a & b.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("invalid strands. Must be of equal length")
	}

	var numDifferences int

	for i := range a {
		if a[i] != b[i] {
			numDifferences++
		}
	}

	return numDifferences, nil
}
