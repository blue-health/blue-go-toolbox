package arrays

func ContainsAny[T comparable](s []T, e ...T) bool {
	for _, a := range s {
		for _, b := range e {
			if a == b {
				return true
			}
		}
	}

	return false
}

func ContainsAll[T comparable](s []T, e ...T) bool {
	for _, a := range e {
		var c bool

		for _, b := range s {
			if a == b {
				c = true
				break
			}
		}

		if !c {
			return false
		}
	}

	return true
}

func Remove[T comparable](s []T, k ...T) []T {
	c := make([]T, 0, len(s))

outer:
	for _, v := range s {
		for _, e := range k {
			if v == e {
				continue outer
			}
		}

		c = append(c, v)
	}

	return c
}
