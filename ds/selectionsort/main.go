package main

import "fmt"

func main() {
	elems := []int{23, 78, 45, 8, 32, 46}
	for i := 0; i < len(elems); i++ {
		min := i
		for j := i + 1; j < len(elems); j++ {
			if elems[j] < elems[min] {
				min = j
			}
		}
		elems[i], elems[min] = elems[min], elems[i]
	}

	fmt.Println(elems)
}
