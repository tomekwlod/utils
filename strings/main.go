package strings

import (
	"strings"
	"unicode"
)

func FirstChar(str string) (c string) {
	if len(str) == 0 {
		return ""
	}
	// value := "ü:ü to eo"
	// Convert string to rune slice before taking substrings.
	// ... This will handle Unicode characters correctly.
	//     Not needed for ASCII strings.
	runes := []rune(str)
	// fmt.Println("First 1:", string(runes[0]))
	// fmt.Println("Last 2:", string(runes[1:]))

	c = string(runes[0])

	return
}

// IsASCII detects if given string contains some special characters
func IsASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

// LastWord return the last word from the string
func LastWord(str string) string {
	exp := SplitString(str)
	last := exp[len(exp)-1]

	return strings.TrimSpace(last)
}

// SplitString splits given string by the spaces and dashes and returns a slice of the words
func SplitString(str string) []string {
	str = strings.TrimSpace(str)

	str = strings.Replace(str, "-", " ", -1)

	return strings.Fields(str)
}
