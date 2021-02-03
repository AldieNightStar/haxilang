package haxilang

import (
	"container/list"
)

// Stack - just a LIFO stack
type Stack struct {
	s *list.List
}

// NewStack - creates new Stack
func NewStack() *Stack {
	return &Stack{
		s: list.New(),
	}
}

// Push - adds elems to the Stack
func (s *Stack) Push(o interface{}) {
	s.s.PushFront(o)
}

// Pop - removes last element from the Stack and returns
func (s *Stack) Pop() (val interface{}) {
	if s.s.Len() > 0 {
		elem := s.s.Front()
		val = elem.Value
		s.s.Remove(elem)
	}
	return
}
