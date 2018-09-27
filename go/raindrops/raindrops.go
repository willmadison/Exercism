package raindrops

import (
	"strconv"
	"strings"
)

const testVersion = 2

// Convert convers the given integer to a
func Convert(i int) string {
	var words []string

	if i%3 == 0 {
		words = append(words, "Pling")
	}

	if i%5 == 0 {
		words = append(words, "Plang")
	}

	if i%7 == 0 {
		words = append(words, "Plong")
	}

	var output string

	if len(words) == 0 {
		output = strconv.Itoa(i)
	} else {
		output = strings.Join(words, "")
	}

	return output
}
