package shared

import (
	"unicode/utf8"
)

// takes a string and a length. If the string exceeds that length, we truncate and replace the last three letters with an elipses
func TruncateString(s string, length int) string {
	runes := []rune(s)
	// logger.Infof("'%s': String Length: %d, VP Length: %d", s, len(s), length)

	if len(runes) < length {
		// no truncation needed
		return s
	}

	truncatedRunes := runes[:length-3]

	return string(truncatedRunes) + "..."
}

// Returns utf8 string length. The tui needs it all the time, and it uses too many characters
func StringLen(s string) int {
	return utf8.RuneCountInString(s)
}
