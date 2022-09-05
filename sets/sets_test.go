package sets_test

import (
	"testing"

	"github.com/blue-health/blue-go-toolbox/sets"
	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	testCases := []struct {
		name string
		a    sets.Set[int]
		b    sets.Set[int]
		c    sets.Set[int]
	}{
		{
			name: "equal sets no change",
			a:    sets.Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			b:    sets.Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			c:    sets.Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		},
		{
			name: "different sets",
			a:    sets.Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			b:    sets.Set[int]{4: struct{}{}, 5: struct{}{}, 6: struct{}{}},
			c: sets.Set[int]{
				1: struct{}{}, 2: struct{}{}, 3: struct{}{},
				4: struct{}{}, 5: struct{}{}, 6: struct{}{},
			},
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			d := sets.Sum(c.a, c.b)

			require.Equal(t, c.c, d)
		})
	}
}

func TestEqual(t *testing.T) {
	testCases := []struct {
		name  string
		a     sets.Set[int]
		b     sets.Set[int]
		equal bool
	}{
		{
			name:  "equal sets",
			a:     sets.Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			b:     sets.Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			equal: true,
		},
		{
			name:  "different sets",
			a:     sets.Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			b:     sets.Set[int]{4: struct{}{}, 5: struct{}{}, 6: struct{}{}},
			equal: false,
		},
		{
			name:  "different order sets",
			a:     sets.Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			b:     sets.Set[int]{3: struct{}{}, 1: struct{}{}, 2: struct{}{}},
			equal: true,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			d := sets.Equal(c.a, c.b)

			require.Equal(t, c.equal, d)
		})
	}
}
