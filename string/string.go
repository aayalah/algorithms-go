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

func longestPalindrome(s string) string {
	runeArr := []rune(s)
	table := make([][]bool, len(runeArr))

	for j := 0; j < len(runeArr); j += 1 {
		table[j] = make([]bool, len(runeArr))
	}

	ans := ""
	length := 0

	for palLen := 0; palLen < len(runeArr); palLen += 1 {
		for i := 0; i <= len(runeArr)-palLen-1; i += 1 {
			if palLen == 0 {
				table[i][i+palLen] = true
			} else if palLen == 1 {
				table[i][i+palLen] = runeArr[i] == runeArr[i-palLen]
			} else {
				table[i][i+palLen] = table[i+1][i+palLen-1] && runeArr[i] == runeArr[i+palLen]
			}

			if table[i][i+palLen] && (palLen+1) > length {
				ans = string(runeArr[i : i+palLen+1])
				length = palLen + 1
			}
		}
	}

	return ans
}
