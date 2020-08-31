package strngs

import (
	"fmt"
	"os"
	"strings"
)

// FindAnagramOrNot find given two strings are anagrams or not
func FindAnagramOrNot(str1, str2 string) {

	if len(strings.Replace(str1, " ", "", -1)) != len(strings.Replace(str2, " ", "", -1)) {
		fmt.Println("Not an anagram")
	}

	for _, first := range str1 {
		flag := false
		for _, second := range str2 {
			if strings.ToLower(string(first)) == strings.ToLower(string(second)) {
				flag = true
				break
			}
		}
		if !flag {
			fmt.Println("strings are not anagrams")
			os.Exit(0)
		}
	}

	fmt.Println("Given strings are Anagrams")
}
