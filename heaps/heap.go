package main

import (
	"fmt"
)

// MaxHeap 最大堆
type MaxHeap struct {
	array []int
}

func (m *MaxHeap) Insert(key int) {
	m.array = append(m.array, key)
	m.MaxHeapUp(len(m.array) - 1)
}

func (m *MaxHeap) MaxHeapUp(index int) {
	for m.array[parent(index)] < m.array[index] {
		m.swap(parent(index), index)
		index = parent(index)
	}
}

func (m *MaxHeap) Extract() int {
	extracted := m.array[0]
	l := len(m.array) - 1

	if len(m.array) == 0 {
		return -1
	}

	m.array[0] = m.array[l]
	m.array = m.array[:l]
	m.MaxHeapDown(0)

	return extracted
}

func (m *MaxHeap) MaxHeapDown(n int) {
	lastIndex := len(m.array) - 1
	l, r := left(n), right(n)
	childToCompare := 0

	// loop while index has at least one child
	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if m.array[l] > m.array[lastIndex] { // 左 大于 lastIndex
			childToCompare = l
		} else {
			childToCompare = r
		}

		if m.array[n] < m.array[childToCompare] {
			m.swap(n, childToCompare)
			n = childToCompare
			l, r = left(n), right(n)
		} else {
			return
		}
	}
}

func (m *MaxHeap) swap(a, b int) {
	m.array[a], m.array[b] = m.array[b], m.array[a]
}

func parent(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return index*2 + 1
}

func right(index int) int {
	return index*2 + 2
}

func main() {
	fmt.Println("hello,ck")

	heap := &MaxHeap{}
	arr := []int{30, 10, 20, 5, 7, 9, 13, 17, 21}

	for _, v := range arr {
		heap.Insert(v)
		fmt.Println(heap)
	}

	for i := 0; i < 5; i++ {
		heap.Extract()
		fmt.Println(heap)
	}
}
