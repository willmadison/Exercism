// Package twofer provides a mechanism for communicating altruistic sharing.
// https://golang.org/doc/effective_go.html#commentary
package twofer

// ShareWith properly communicates ones altruism.
func ShareWith(person string) string {
	if person == "" {
		person = "you"
	}

	return "One for " + person + ", one for me."
}
