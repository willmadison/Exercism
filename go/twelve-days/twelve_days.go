package twelve

import (
	"bytes"
	"fmt"
	"strings"
)

const testVersion = 1

//Song outputs the entire 12 Days of Christmas carol.
func Song() string {
	var buffer bytes.Buffer

	for i := 1; i <= 12; i++ {
		buffer.WriteString(Verse(i))
		buffer.WriteRune('\n')
	}

	return buffer.String()
}

var days = []string{
	"",
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

var numbers = []string{
	"",
	"a",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
}

var gifts = []string{
	"",
	"Partridge in a Pear Tree",
	"Turtle Doves",
	"French Hens",
	"Calling Birds",
	"Gold Rings",
	"Geese-a-Laying",
	"Swans-a-Swimming",
	"Maids-a-Milking",
	"Ladies Dancing",
	"Lords-a-Leaping",
	"Pipers Piping",
	"Drummers Drumming",
}

//Verse outputs a given verse for the integer day of the 12 Days of Christmas.
func Verse(day int) string {
	parts := []string{fmt.Sprintf("On the %s day of Christmas my true love gave to me", days[day])}

	for day > 0 {
		stanza := fmt.Sprintf("%s %s", numbers[day], gifts[day])

		if day == 1 && len(parts) > 1 {
			stanza = "and " + stanza
		}

		parts = append(parts, stanza)
		day--
	}

	return strings.Join(parts, ", ") + "."
}
