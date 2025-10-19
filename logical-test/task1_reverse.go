package logicaltest

import "strings"

// type Student struct {
// 	name string
// 	age  int
// }

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

// func Haha(){
// 	// var students []Student
// 	students := []Student{
// 		{"hao",21},

// 	}
// }
