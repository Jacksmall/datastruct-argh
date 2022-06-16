package main

import "fmt"

// Stack LIFO 后进先出
type Stack struct {
	slice []int
}

// push
func (s *Stack) Push(value int) {
	s.slice = append(s.slice, value)
}

// pop
func (s *Stack) Pop() int {
	var toPop int
	l := len(s.slice) - 1
	toPop = s.slice[l]
	s.slice = s.slice[:l]
	return toPop
}

func main() {
	fmt.Println("Stack")

	s := &Stack{}
	pushed := []int{100, 200, 300}
	for _, v := range pushed {
		s.Push(v)
	}
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
}
