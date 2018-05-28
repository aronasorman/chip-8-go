package chip8

import (
	"errors"
	"sync"
)

// Stack implementation, with Push and Pop methods
type Stack struct {
	lock sync.Mutex
	s    []uint16
}

// NewStack returns a newly created stack
func NewStack() *Stack {
	return &Stack{s: make([]uint16, 24)}
}

// Push the given value on to the stack
func (s *Stack) Push(val uint16) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, val)
}

// Pop returns the top of the stack, or returns 0
// and an error if the stack is empty
func (s *Stack) Pop() (uint16, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	if l == 0 {
		return 0, errors.New("Empty stack")
	}

	result := s.s[l-1]
	s.s = s.s[:l-1]
	return result, nil
}
