// Copyright (C) 2019 José Martínez Ruiz <jmmrcp@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package texttowidth

import (
	"unicode"
)

//Format -- Formats a string of text to a specified width
func Format(text string, width int) (result string) {
	long := len(text)
	if long < width {
		return text
	}

	index := width
	previousIndex := 0
	for index < long {
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
