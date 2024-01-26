package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_Push(t *testing.T) {
	stack := Stack[int]{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	assert.Equal(t, []int{1, 2, 3}, stack.items)
}

func TestStack_Pop(t *testing.T) {
	stack := Stack[int]{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	item, err := stack.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 3, item)
	assert.Equal(t, []int{1, 2}, stack.items)
}

func TestStack_Pop_Error(t *testing.T) {
	// Test pop from an empty stack
	emptyStack := Stack[int]{}
	_, err := emptyStack.Pop()
	assert.EqualError(t, err, ErrStackEmpty.Error())
}

func TestStack_Peek(t *testing.T) {
	stack := Stack[int]{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	item, err := stack.Peek()
	assert.NoError(t, err)
	assert.Equal(t, 3, item)
	assert.Equal(t, []int{1, 2, 3}, stack.items)
}

func TestStack_Peek_Error(t *testing.T) {
	// Test peek from an empty stack
	emptyStack := Stack[int]{}
	_, err := emptyStack.Peek()
	assert.EqualError(t, err, ErrStackEmpty.Error())
}

func TestStack_IsEmpty(t *testing.T) {
	// Test IsEmpty on a non-empty stack
	stack := Stack[int]{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	assert.False(t, stack.IsEmpty())
}

func TestStack_IsEmpty_EmptyStack(t *testing.T) {
	// Test IsEmpty on an empty stack
	emptyStack := Stack[int]{}
	assert.True(t, emptyStack.IsEmpty())
}
