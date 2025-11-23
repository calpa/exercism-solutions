package logs

import (
    "strings"
    "unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(line string) string {
	for _, c := range line {
		switch c {
		case '\u2757':
			return "recommendation"
		case '\U0001F50D':
			return "search"
		case '\u2600':
			return "weather"
		}
	}
    
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
