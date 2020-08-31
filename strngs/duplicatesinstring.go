package strngs

import "fmt"

// FindDuplicatesInGivenString find the duplicate chars in given string
func FindDuplicatesInGivenString(str string) {

	set := make(map[string]bool)

	for _, c := range str {
		counter := 0
		for _, ic := range str {
			if counter == 2 {
				if !set[string(c)] {
					set[string(c)] = true
				}
				break
			}
			if c == ic {
				counter = counter + 1
			}
		}
	}

	fmt.Printf("%v \n", set)
}
