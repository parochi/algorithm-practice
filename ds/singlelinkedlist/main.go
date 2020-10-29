package main

import (
	"fmt"
)

type node struct {
	next *node
	data int
}
type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) append(n *node) {
	if l.head == nil {
		l.head = n
	} else {
		tmp := l.head
		for tmp.next != nil {
			tmp = tmp.next
		}
		tmp.next = n
	}
	l.length++
}

func (l *linkedList) remove(val int) {
	if l.length == 0 {
		return
	}

	if l.head.data == val {
		l.head = l.head.next
		l.length--
		return
	}

	tmp := l.head
	flag := false
	for tmp.next.data == val {
		tmp = tmp.next
		flag = true
	}
	if flag {
		tmp.next = tmp.next.next
		l.length--
	}
	return
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
	localList.append(node1)
	node2 := &node{data: 12}
	localList.append(node2)
	node3 := &node{data: 13}
	localList.append(node3)
	node4 := &node{data: 14}
	localList.append(node4)
	node5 := &node{data: 15}
	localList.append(node5)

	localList.printData()

	localList.remove(24)
	localList.printData()
}
