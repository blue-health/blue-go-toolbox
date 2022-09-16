package arrays

func Contains[T comparable](s []T, e ...T) bool {
	for _, a := range s {
		for _, b := range e {
			if a == b {
				return true
			}
		}
	}

	return false
}

func Subset[T comparable](s []T, e ...T) bool {
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

func Intersects[T comparable](s []T, e ...T) bool {
	for _, a := range e {
		if Contains(s, a) {
			return true
		}
	}

	return false
}

func Remove[T comparable](s []T, k T) []T {
	for i, v := range s {
		if v == k {
			s = append(s[:i], s[i+1:]...)
			break
		}
	}

	return s
}
