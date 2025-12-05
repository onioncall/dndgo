package shared

// takes a string and a length. If the string exceeds that length, we truncate and replace the last three letters with an elipses
func TruncateString(s string, length int) string {
	runes := []rune(s)

	if len(runes) < length {
		// no truncation needed
		return s
	}

	truncatedRunes := runes[:length-3]

	return string(truncatedRunes) + "..."
}
