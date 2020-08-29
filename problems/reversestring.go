package problems

import (
	"fmt"
)

// ReverseString reverse given string
func ReverseString() {
	str := "abcdef"
	fmt.Println(reverseIt(str))
}

func reverseIt(s string) string {

	if len(s) == 0 {
		return s
	}

	return string(s[len(s)-1]) + reverseIt(s[0:len(s)-1])
}
