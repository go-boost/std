package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet_Add(t *testing.T) {
	s := make(Set[int])
	s.Add(1)

	assert.True(t, s.Contains(1), "Expected set to contain 1")
	assert.False(t, s.Contains(2), "Expected set not to contain 2")
}

func TestSet_Remove(t *testing.T) {
	s := make(Set[int])
	s.Add(1)
	s.Add(2)

	s.Remove(1)

	assert.False(t, s.Contains(1), "Expected set not to contain 1 after removal")
	assert.True(t, s.Contains(2), "Expected set to contain 2")
}

func TestSet_Contains(t *testing.T) {
	s := make(Set[string])
	s.Add("apple")

	assert.True(t, s.Contains("apple"), "Expected set to contain 'apple'")
	assert.False(t, s.Contains("orange"), "Expected set not to contain 'orange'")
}

func TestSet_Size(t *testing.T) {
	s := make(Set[int])
	s.Add(1)
	s.Add(2)

	assert.Equal(t, 2, s.Size(), "Expected size of set to be 2")
}

func TestSet_ToSlice(t *testing.T) {
	s := make(Set[string])
	s.Add("apple")
	s.Add("orange")

	slice := s.ToSlice()
	expectedSlice := []string{"apple", "orange"}

	assert.ElementsMatch(t, expectedSlice, slice, "Expected set elements %v, got %v", expectedSlice, slice)
}
