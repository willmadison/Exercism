package scrabble

import "unicode"

// Score calculates the raw Scrabble points scored for a given word
// without consideration for special tile locations (i.e. Double/Triple Word,
// Double/Triple Letter, etc.).
func Score(word string) int {
	var score int

	for _, c := range word {
		score += pointValue(c)
	}

	return score
}

func pointValue(c rune) int {
	switch unicode.ToLower(c) {
	case 'd', 'g':
		return 2
	case 'b', 'c', 'm', 'p':
		return 3
	case 'f', 'h', 'v', 'w', 'y':
		return 4
	case 'k':
		return 5
	case 'j', 'x':
		return 8
	case 'q', 'z':
		return 10
	default:
		return 1
	}
}
