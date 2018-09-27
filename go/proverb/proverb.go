// Package proverb provides mechanisms for generating catchy proverbs.
package proverb

import "fmt"

// Proverb generates a proper proverb for the given items.
func Proverb(items []string) []string {
	var stanzas []string

	if len(items) > 0 {
		for i := 0; i < len(items)-1; i++ {
			stanzas = append(stanzas, fmt.Sprintf("For want of a %v the %v was lost.", items[i], items[i+1]))
		}

		stanzas = append(stanzas, fmt.Sprintf("And all for the want of a %v.", items[0]))
	}

	return stanzas
}
