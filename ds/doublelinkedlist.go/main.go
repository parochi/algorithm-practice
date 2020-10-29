package main

import (
	"fmt"
)

type node struct {
	next *node
	prev *node
	data int
}
type linkedList struct {
	head   *node
	tail   *node
	length int
}

func (l *linkedList) addFirst(n *node) {
	if l.length == 0 {
		l.head = n
		l.tail = n
	} else {
		l.head.prev = n
		l.head = l.head.prev
	}
	l.length++
}

func (l linkedList) printData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d,", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}

	fmt.Println()
}

func main() {
	localList := linkedList{}
	node1 := &node{data: 11}
	localList.addFirst(node1)

	node2 := &node{data: 12}
	localList.addFirst(node2)

	node3 := &node{data: 13}
	localList.addFirst(node3)
	/*
		node4 := &node{data: 14}
		localList.addFirst(node4)
		node5 := &node{data: 15}
		localList.addFirst(node5)
	*/
	localList.printData()

}
