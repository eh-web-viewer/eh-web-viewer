package tools

import "fmt"

type Slice[T comparable] []T

func NewSlice[T comparable](e ...T) Slice[T] {
	return Slice[T](e)
}

func (s Slice[T]) Get(index int) (T, error) {
	if index < 0 || index >= len(s) {
		var defaultValue T
		return defaultValue, fmt.Errorf("out of range")
	}
	return s[index], nil
}

func (s Slice[T]) GetOrDefault(index int, defaultValue T) T {
	if index < 0 || index >= len(s) {
		return defaultValue
	}
	return s[index]
}

func (s Slice[T]) Last() T {
	return s[len(s)-1]
}

func (s Slice[T]) Push(e T) Slice[T] {
	s = append(s, e)
	return s
}

func (s Slice[T]) Pop() T {
	e := s.Last()
	s = s[:len(s)-1]
	return e
}

func (s Slice[T]) FirstNonDefaultValue(defaultValue T) T {
	for _, e := range s {
		if e != defaultValue {
			return e
		}
	}
	return defaultValue
}

// func (s Slice[T]) Map(index i[RT any]nt, defaultValue T) RT {
// 	if index < 0 || len(s) >= index {
// 		return defaultValue
// 	}
// 	return s[index]
// }
