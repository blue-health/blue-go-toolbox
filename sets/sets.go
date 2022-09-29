package sets

type Set[T comparable] map[T]struct{}

func Sum[T comparable](i ...Set[T]) Set[T] {
	c := make(map[T]struct{}, max(i...))

	for _, e := range i {
		for k := range e {
			c[k] = struct{}{}
		}
	}

	return c
}

func Subtract[T comparable](a Set[T], i ...Set[T]) Set[T] {
	c := make(map[T]struct{}, max(i...))

	sum := Sum(i...)

	for k := range a {
		if _, ok := sum[k]; !ok {
			c[k] = struct{}{}
		}
	}

	return c
}

func Union[T comparable](i ...Set[T]) Set[T] {
	var (
		c = make(map[T]int, max(i...))
		s = make(map[T]struct{}, min(i...))
	)

	for _, e := range i {
		for k := range e {
			c[k]++
		}
	}

	for k, v := range c {
		if v == len(i) {
			s[k] = struct{}{}
		}
	}

	return s
}

func Equal[T comparable](i ...Set[T]) bool {
	for _, a := range i {
		for k := range a {
			for _, b := range i {
				if _, ok := b[k]; !ok {
					return false
				}
			}
		}
	}

	return true
}

func From[T comparable](v ...T) Set[T] {
	r := make(Set[T], len(v))

	for i := range v {
		r[v[i]] = struct{}{}
	}

	return r
}

func FromSlice[T comparable](s []T) Set[T] {
	r := make(Set[T], len(s))

	for x := range s {
		r[s[x]] = struct{}{}
	}

	return r
}

func (s Set[T]) ToSlice() []T {
	r := make([]T, 0, len(s))

	for x := range s {
		r = append(r, x)
	}

	return r
}

func max[T comparable](i ...Set[T]) int {
	var m int
	for _, a := range i {
		if len(a) > m {
			m = len(a)
		}
	}

	return m
}

func min[T comparable](i ...Set[T]) int {
	var m int
	for _, a := range i {
		if len(a) < m {
			m = len(a)
		}
	}

	return m
}
