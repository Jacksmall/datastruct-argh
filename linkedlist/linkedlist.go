package main

import "fmt"

/*
 * 链表结构体
 * 包含指向head node的指针 和 长度 int
 */
type linkedList struct {
	head   *node
	length int
}

/*
 * 链表中的node
 * data 保存的数据
 * next 指向下一个node的指针
 */
type node struct {
	data int
	next *node
}

// prepend FI 前置入队列 lpush node
func (l *linkedList) prepend(node ...*node) {
	for _, m := range node {
		second := l.head
		l.head = m
		l.head.next = second
		l.length++
	}
}

// search
func (l linkedList) search(value int) bool {
	toSearch := l.head.data
	if toSearch == value {
		return true
	}

	for l.length > 0 {
		if l.head.data == value {
			return true
		}
		l.head = l.head.next
		l.length--
	}
	return false
}

// PrintMyList 打印队列
func (l linkedList) PrintMyList() {
	toPrint := l.head
	for l.length > 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

// delete 出队列 lpop
func (l *linkedList) delete(value int) {
	if l.length == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	previousToDelete := l.head
	for l.length > 0 {
		if previousToDelete.next.data != value {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

// Len 队列长度 时间复杂度 O(1)
func (l *linkedList) Len() int { return l.length }

func main() {
	mylist := linkedList{}
	node1 := &node{data: 89}
	node2 := &node{data: 814}
	node3 := &node{data: 82}
	node4 := &node{data: 9}
	node5 := &node{data: 15}
	node6 := &node{data: 22}
	node7 := &node{data: 33}
	mylist.prepend(node1, node2, node3, node4, node5, node6, node7)
	fmt.Println(mylist.search(82))
	mylist.PrintMyList()
	mylist.delete(33)
	fmt.Println(mylist.search(33))
	mylist.PrintMyList()
}
