package main

import (
	"fmt"

	"github.com/algorithm-practice/ds/arrs"
	"github.com/algorithm-practice/strngs"
)

func main() {

	fmt.Println("Display Triangle")
	strngs.DisplayTriangle()
	fmt.Println("Reverse String")
	strngs.ReverseString()
	fmt.Println("Duplicate Characters")
	strngs.FindDuplicatesInGivenString("martin thomas")
	fmt.Println("String Anagrams")
	strngs.FindAnagramOrNot("School master", "The classroom")
	fmt.Println("Get the string permutations")
	strngs.GetAllPermutationsOfGivenString("abcd")

	fmt.Println("Print the missing number")
	arrs.FindMissingNumber([]int{1, 2, 3, 5}, 5)
	arrs.FindMoreThanOneMissingNumber([]int{1, 2, 5, 8, 9}, 10)
}
