package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	len := len(s.items)
	if len == 0 {
		return 0 // or handle underflow
	}
	item := s.items[len-1]
	s.items = s.items[:len-1]
	return item
}

func main() {
	mystack := Stack{}
	mystack.Push(1)
	mystack.Push(2)
	mystack.Push(3)

	fmt.Println(mystack)

	fmt.Println("Popped item:", mystack.Pop())
	fmt.Println("Popped item:", mystack.Pop())
	fmt.Println(mystack)
}
