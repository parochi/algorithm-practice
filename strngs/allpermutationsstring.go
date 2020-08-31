package strngs

import "fmt"

// GetAllPermutationsOfGivenString - get all the permutations of given string
func GetAllPermutationsOfGivenString(str string) {
	generate("", str)
}

func generate(perm, word string) {
	if len(word) == 0 {
		fmt.Printf("%s%s\n", perm, word)
	}

	for i, c := range word {
		l := len(word)
		generate(fmt.Sprintf("%s%s", perm, string(c)), fmt.Sprintf("%s%s", word[0:i], word[i+1:l]))
	}
}
