package safemap

import "sync"

type SafeMap[T0 comparable, T1 any] struct {
	mutex sync.RWMutex
	val   map[T0]T1
}

func New[T0 comparable, T1 comparable]() *SafeMap[T0, T1] {
	return &SafeMap[T0, T1]{
		val: map[T0]T1{},
	}
}

func (s *SafeMap[T0, T1]) Put(key T0, value T1) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.val[key] = value
}

func (s *SafeMap[T0, T1]) Get(key T0) (T1, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	v, ok := s.val[key]
	return v, ok
}

func (s *SafeMap[T0, T1]) Delete(key T0) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.val, key)
}

func (s *SafeMap[T0, T1]) Len() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return len(s.val)
}
