package logicaltest

import "strings"

func Reverse(text string) string {
	var result string
	splitted := strings.Split(text, " ")
	for _, v := range splitted {
		for i := len(v) - 1; i >= 0; i-- {
			result += string(v[i])
		}
		result += " "

	}
	return strings.TrimSpace(result)
}
