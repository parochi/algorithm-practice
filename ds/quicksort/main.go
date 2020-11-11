package main

import "fmt"

func main() {
	elems := []int{3, 22, 6, 14, 7, 8, 5}
	quickSort(elems, 0, len(elems)-1)
	fmt.Println(elems)
}

func quickSort(elems []int, low, high int) {
	if low < high {
		pindex := parition(elems, low, high)

		quickSort(elems, low, pindex-1)
		quickSort(elems, pindex+1, high)
	}
}

func parition(elems []int, l, h int) int {
	pivot := elems[h]
	i := l - 1
	for j := l; j <= h-1; j++ {
		if elems[j] < pivot {
			i++
			elems[i], elems[j] = elems[j], elems[i]
		}
	}
	elems[i+1], elems[h] = elems[h], elems[i+1]
	return i + 1
}
