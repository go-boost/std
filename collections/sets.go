package collections

// Set is a generic set implementation using a map with empty struct as a placeholder value
type Set[T comparable] map[T]struct{}

// Add adds an element to the set
func (s Set[T]) Add(item T) {
	s[item] = struct{}{}
}

// Remove removes an element from the set
func (s Set[T]) Remove(item T) {
	delete(s, item)
}

// Contains checks if the set contains a specific element
func (s Set[T]) Contains(item T) bool {
	_, exists := s[item]
	return exists
}

// Has checks if the set contains a specific element
//
// Deprecated: Use Contains instead. Has will be removed in the next release.
func (s Set[T]) Has(item T) bool {
	return s.Contains(item)
}

// Size returns the number of elements in the set
func (s Set[T]) Size() int {
	return len(s)
}

// ToSlice returns a slice of all elements in the set
func (s Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(s))
	for item := range s {
		slice = append(slice, item)
	}
	return slice
}
