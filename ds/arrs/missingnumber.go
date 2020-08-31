package arrs

import (
	"fmt"
	"os"
)

/*
FindMissingNumber The formula for total sum of the given numbers is
for example sum of 5 numbers
	1 + 2 + 3 + 4 + 5 = 15 = n (n + 1) / 2
*/
func FindMissingNumber(numbers []int, totalNumbers int) {
	if len(numbers) == totalNumbers {
		fmt.Println("None of the number is missing")
		os.Exit(0)
	}

	totalCount := 0
	for _, number := range numbers {
		totalCount = totalCount + number
	}
	result := (totalNumbers*(totalNumbers+1))/2 - totalCount
	fmt.Printf("missing number %d \n", result)
}

// FindMoreThanOneMissingNumber find the more than one missing number
func FindMoreThanOneMissingNumber(numbers []int, totalNumbers int) {

	for i := 1; i <= totalNumbers; i++ {
		found := false
		for _, n := range numbers {
			if n == i {
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("Missing Number %d\n", i)
		}
	}
}
