package logicaltest

import "unicode"

func NumCount(arr []string) int {
	var count int

	for _, v := range arr {
		runes := []rune(v)
		isDigit := unicode.IsDigit(runes[0])
		if isDigit {
			count++
		}
	}
	return count
}
