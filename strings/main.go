package strings

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
