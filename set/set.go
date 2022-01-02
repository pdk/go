package set

type Set[T comparable] struct {
	items map[T]struct{}
}

func New[T comparable]() Set[T] {
	return Set[T]{
		items: map[T]struct{}{},
	}
}

func (s *Set[T]) Add(item T) {
	if s.items == nil {
		s.items = make(map[T]struct{})
	}
	s.items[item] = struct{}{}
}

func (s *Set[T]) Remove(item T) {
	if len(s.items) == 0 {
		return
	}

	delete(s.items, item)
}

func (s *Set[T]) Contains(item T) bool {
	if len(s.items) == 0 {
		return false
	}

	_, present := s.items[item]
	return present
}

func (s *Set[T]) Size() int {
	return len(s.items)
}