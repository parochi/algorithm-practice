package main

import (
	"fmt"

	"github.com/algorithm-practice/problems"
)

func main() {
	fmt.Println("Display Triangle")
	problems.DisplayTriangle()
	fmt.Println("Reverse String")
	problems.ReverseString()
	fmt.Println("Duplicate Characters")
	problems.FindDuplicatesInGivenString("martin thomas")
	fmt.Println("String Anagrams")
	problems.FindAnagramOrNot("School master", "The classroom")
}
