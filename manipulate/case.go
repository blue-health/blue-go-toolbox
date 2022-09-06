package manipulate

import "strings"

func SnakeCase(s string) string {
	return delimiterCase(s, '_', false)
}

func UpperSnakeCase(s string) string {
	return delimiterCase(s, '_', true)
}

func delimiterCase(s string, delimiter rune, upperCase bool) string {
	s = strings.TrimSpace(s)
	buffer := make([]rune, 0, len(s)+3)

	adjustCase := toLower
	if upperCase {
		adjustCase = toUpper
	}

	var (
		prev rune
		curr rune
	)

	for _, next := range s {
		switch {
		case isDelimiter(curr):
			if !isDelimiter(prev) {
				buffer = append(buffer, delimiter)
			}
		case isUpper(curr):
			if isLower(prev) || (isUpper(prev) && isLower(next)) {
				buffer = append(buffer, delimiter)
			}

			buffer = append(buffer, adjustCase(curr))
		case curr != 0:
			buffer = append(buffer, adjustCase(curr))
		}

		prev = curr
		curr = next
	}

	if len(s) > 0 {
		if isUpper(curr) && isLower(prev) && prev != 0 {
			buffer = append(buffer, delimiter)
		}

		buffer = append(buffer, adjustCase(curr))
	}

	return string(buffer)
}

// isLower checks if a character is lower case. More precisely it evaluates if it is
// in the range of ASCII character 'a' to 'z'.
func isLower(ch rune) bool {
	return ch >= 'a' && ch <= 'z'
}

// toLower converts a character in the range of ASCII characters 'A' to 'Z' to its lower
// case counterpart. Other characters remain the same.
func toLower(ch rune) rune {
	if ch >= 'A' && ch <= 'Z' {
		return ch + 32
	}

	return ch
}

// isLower checks if a character is upper case. More precisely it evaluates if it is
// in the range of ASCII characters 'A' to 'Z'.
func isUpper(ch rune) bool {
	return ch >= 'A' && ch <= 'Z'
}

// toLower converts a character in the range of ASCII characters 'a' to 'z' to its lower
// case counterpart. Other characters remain the same.
func toUpper(ch rune) rune {
	if ch >= 'a' && ch <= 'z' {
		return ch - 32
	}

	return ch
}

// isSpace checks if a character is some kind of whitespace.
func isSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

// isDelimiter checks if a character is some kind of whitespace or '_' or '-'.
func isDelimiter(ch rune) bool {
	return ch == '-' || ch == '_' || isSpace(ch)
}
