package main

import "fmt"

type stack struct {
	item []int
}

func (s *stack) push(item int) {
	s.item = append(s.item, item)
}

func (s *stack) pop() int {
	if len(s.item) == 0 {
		return 0
	}
	top := s.item[len(s.item)-1]
	s.item = s.item[:len(s.item)-1]
	return top
}

func main() {
	stack := stack{}
	stack.push(20)
	stack.push(30)
	stack.push(40)
	fmt.Println(stack.item)
	// fmt.Println(stack.pop())

}
