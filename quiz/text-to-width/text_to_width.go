package texttowidth

import (
	"unicode"
)

//Format -- Formats a string of text to a specified width
func Format(text string, width int) (result string) {
	if len(text) < width {
		return text
	}

	index := width
	previousIndex := 0
	for index < len(text) {
		spaceExist := false
		for i := index; i > index-width; i-- {
			if unicode.IsSpace(rune(text[i])) {
				index = i
				spaceExist = true
				break
			}
		}
		result += text[previousIndex:index] + "\n"

		//If space, add 1 to the index to skip over the blank space
		if spaceExist {
			previousIndex = index + 1
		} else {
			previousIndex = index
		}
		index += width
	}

	//Adds the rest of the text to result
	result += text[previousIndex:]

	return result
}
