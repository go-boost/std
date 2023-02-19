package collections

type Set[T comparable] map[T]struct{}

func (set Set[T]) Add(value T) {
	set[value] = struct{}{}
}

func (set Set[T]) Remove(value T) {
	delete(set, value)
}

func (set Set[T]) Has(value T) bool {
	_, exist := set[value]
	return exist
}
