package isbn

import (
	"strings"
	"unicode"
    "unicode/utf8"
)

func IsValidISBN(isbn string) bool {
	cleanIsbn := strings.ReplaceAll(isbn, "-", "") // to be sure that have no dashes

	runeCount := utf8.RuneCountInString(cleanIsbn)
	if runeCount != 10 { // ISBN must have 10 digits or 9 + 'X' char
		return false
	}

	result := 0
	for idx, val := range cleanIsbn {
		// Check if 'X' appears in any position other than the last
		if val == 'X' && idx != 9 {
			return false
		}

		if val == 'X' {
			result += (runeCount - idx) * 10
		} else if !unicode.IsDigit(val) {
			return false
		} else {
			result += (runeCount - idx) * int(val-'0')
		}
	}

	return result%11 == 0
}