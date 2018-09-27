package acronym

import "strings"

const testVersion = 1

func abbreviate(word string) string {
	word = strings.Replace(word, "-", " ", -1)

	words := strings.Split(word, " ")

	var firstCharacters []string

	for _, w := range words {
		firstCharacters = append(firstCharacters, toAcronym(w))
	}

	return strings.ToUpper(strings.Join(firstCharacters, ""))
}

func toAcronym(word string) string {
	word = strings.Trim(word, ",:")

	words := strings.FieldsFunc(word, func(char rune) bool {
		return strings.ContainsRune("abcdefghijklmnopqrstuvwxyz", char)
	})

	var acronym string

	switch {
	case len(words) > 0:
		var characters []string

		for _, w := range words {
			characters = append(characters, string(w[0]))
		}

		acronym = strings.Join(characters, "")
	default:
		acronym = string(word[0])
	}

	return acronym
}
