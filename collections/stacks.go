package collections

import "errors"

// ErrStackEmpty is an error indicating that the stack is empty
var ErrStackEmpty = errors.New("stack is empty")

// Stack is a generic stack implementation using slices and generics
type Stack[T any] struct {
	items []T
}

// Push adds an item to the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		return *new(T), ErrStackEmpty
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

// Peek returns the top item from the stack without removing it
func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		return *new(T), ErrStackEmpty
	}

	return s.items[len(s.items)-1], nil
}

// IsEmpty returns true if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}
