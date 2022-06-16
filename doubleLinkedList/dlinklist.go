package main

import "fmt"

type dlinkedList struct {
	length int
	head   *node
}

type node struct {
	data int
	next *node
	prev *node
}

func (d *dlinkedList) prepend(node ...*node) {
	for k, m := range node {
		second := d.head
		d.head = m
		d.head.next = second
		if k == 0 {
			d.head.prev = nil
		} else {
			d.head.next.prev = second
		}

		d.length++
	}
}

func (l dlinkedList) PrintMyList() {
	toPrint := l.head
	for l.length > 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func main() {
	fmt.Println("Welcome to double linked list")
	mydlinkedList := dlinkedList{}
	node1 := &node{data: 100}
	node2 := &node{data: 80}
	node3 := &node{data: 70}
	node4 := &node{data: 50}
	node5 := &node{data: 10}
	mydlinkedList.prepend(node5, node4, node3, node2, node1)
	mydlinkedList.PrintMyList()
}
