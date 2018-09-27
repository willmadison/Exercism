package pangram

import "strings"

const testVersion = 1

//IsPangram determines whether or not the given sentence is a pangram.
func IsPangram(sentence string) bool {
	occcurencesByCharacter := map[rune]int{}

	for _, x := range `.1234567890abcdefghijklmnopqrstuvwxyz"` {
		occcurencesByCharacter[x] = 0
	}

	sentence = strings.ToLower(sentence)

	for _, character := range sentence {
		if character == ' ' {
			continue
		}

		if _, present := occcurencesByCharacter[character]; !present {
			continue
		}

		occcurencesByCharacter[character]++
	}

	for _, x := range "abcdefghijklmnopqrstuvwxyz" {
		if occcurencesByCharacter[x] == 0 {
			return false
		}
	}

	return true
}
