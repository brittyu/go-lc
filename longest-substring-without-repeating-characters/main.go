package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	charMap := make(map[rune]int)
	maxLength := 0
	startIndex := 0
	for i, char := range s {
		if index, ok := charMap[char]; ok && index >= startIndex {
			startIndex = index + 1
		}

		charMap[char] = i
		currentLength := i - startIndex + 1

		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

func main() {
	example := "abcabcbb"
	result := lengthOfLongestSubstring(example)

	fmt.Printf("example result: %v", result)
}
