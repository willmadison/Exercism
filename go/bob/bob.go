package bob

import "strings"

const testVersion = 2

// Hey returns Bob's response to the given phrase.
func Hey(phrase string) string {
	phrase = strings.TrimSpace(phrase)

	var response string

	switch {
	case strings.ContainsAny(phrase, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") &&
		phrase == strings.ToUpper(phrase):
		response = "Whoa, chill out!"
	case strings.HasSuffix(phrase, "?"):
		response = "Sure."
	case phrase == "":
		response = "Fine. Be that way!"
	default:
		response = "Whatever."
	}

	return response
}
