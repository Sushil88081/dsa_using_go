package main

import (
	"fmt"
)

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return 0
	}
	top := len(s.items) - 1
	item := s.items[top]
	s.items = s.items[:top]
	return item
}
func main() {
	var item int
	Stack := Stack{}
	fmt.Println("Enter item which you insert in stack:")
	for i := 0; i < 5; i++ {
		fmt.Scan(&item)
		Stack.Push(item)

	}

	fmt.Println(" element of stack is:", Stack.items)
	fmt.Println("Popped element is:", Stack.Pop())
	fmt.Println("element of stack is:", Stack.items)

}
