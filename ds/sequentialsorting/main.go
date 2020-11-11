package main

import "fmt"

var (
	firstArray  []int
	secondArray []int
)

func main() {
	firstArray = []int{3, 22, 6, 14, 7, 8, 5}
	secondArray = []int{9, 6, 64, 5, 82, 17, 8}

	mergeresult := append(firstArray, secondArray...)

	fmt.Println(sortIt(mergeresult))
}

// Order of n^2
func sortIt(input []int) []int {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input)-1; j++ {
			if input[j] < input[j+1] {
				temp := input[j]
				input[j] = input[j+1]
				input[j+1] = temp
			}
		}
	}
	return input
}
