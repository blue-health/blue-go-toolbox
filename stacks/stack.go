package stacks

type Stack[T comparable] []T

func (s *Stack[T]) Push(val T) {
	*s = append(*s, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(*s) == 0 {
		var zero T
		return zero, false
	}

	top := (*s)[len(*s)-1]

	(*s) = (*s)[:len(*s)-1]

	return top, true
}

func (s Stack[T]) Contains(val T) bool {
	for _, v := range s {
		if v == val {
			return true
		}
	}

	return false
}

func (s Stack[T]) Iterate(f func(T)) {
	for _, v := range s {
		f(v)
	}
}
