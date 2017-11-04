package string

func lengthOfLongestSubstring(s string) int {
	seenLetters := make(map[rune]bool)
	runeArr := []rune(s)

	startIndex := 0
	endIndex := 0
	length := 0

	for ; endIndex < len(runeArr); endIndex += 1 {
		currChar := runeArr[endIndex]
		if ind, ok := seenLetters[currChar]; ok && ind == startIndex {
			if currLength > length {
				length = currLength
			}
			startIndex = ind + 1
			seenLetters[currChar] = endIndex
		} else {
			seenLetters[currChar] = endIndex
		}
	}

	currLength := endIndex - startIndex
	if currLength > length {
		length = currLength
	}
	return length
}
