package problem

import "sync"

type set[V comparable] struct {
	mu sync.Mutex
	m  map[V]bool
}

func (s *set[V]) Contains(v V) bool {
	if s == nil || s.m == nil {
		return false
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.m[v]
	return ok
}

func (s *set[V]) Len() int {
	if s == nil || s.m == nil {
		return 0
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	if s.m == nil {
		return 0
	}
	return len(s.m)
}

func (s *set[V]) Add(v V) {
	if s == nil {
		s = &set[V]{}
	}
	if s.m == nil {
		s.m = map[V]bool{}
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	if s.m == nil {
		s.m = make(map[V]bool)
	}

	s.m[v] = true
}

func (s *set[V]) Del(v V) {
	if s == nil || s.m == nil {
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	if s.m == nil {
		return
	}

	delete(s.m, v)
}

func (s *set[V]) List() []V {
	if s == nil || s.m == nil {
		return []V{}
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	if s.m == nil {
		return []V{}
	}

	arr := make([]V, 0, len(s.m))
	for v := range s.m {
		arr = append(arr, v)
	}
	return arr
}

func (s *set[V]) Range(f func(V) bool) {
	if s == nil || s.m == nil {
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	if s.m == nil {
		return
	}

	for v := range s.m {
		if !f(v) {
			return
		}
	}
}
