package main

import "fmt"

func main() {
	elems := []int{64, 34, 25, 12, 22, 11, 90}
	for i := 0; i < len(elems); i++ {
		swap := false
		for j := 0; j < len(elems)-i-1; j++ {
			if elems[j] > elems[j+1] {
				elems[j], elems[j+1] = elems[j+1], elems[j]
				swap = true
			}
		}
		if !swap {
			break
		}
	}

	fmt.Println(elems)
}
