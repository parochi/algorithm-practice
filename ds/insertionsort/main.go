package main

import "fmt"

func main() {
	elems := []int{3, 22, 6, 14, 7, 8, 5}
	for i := 1; i < len(elems); i++ {
		key := elems[i]
		j := i - 1

		for j >= 0 && elems[j] > key {
			elems[j+1] = elems[j]
			j = j - 1
		}

		elems[j+1] = key
	}

	fmt.Println(elems)
}
