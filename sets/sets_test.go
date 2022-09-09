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
			a:    sets.From(1, 2, 3),
			b:    sets.From(1, 2, 3),
			c:    sets.From(1, 2, 3),
		},
		{
			name: "different sets",
			a:    sets.From(1, 2, 3),
			b:    sets.From(4, 5, 6),
			c: sets.From(
				1, 2, 3,
				4, 5, 6,
			),
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			d := sets.Sum(c.a, c.b)

			require.Equal(t, c.c, d)
		})
	}
}

func TestUnion(t *testing.T) {
	testCases := []struct {
		name string
		a    sets.Set[int]
		b    sets.Set[int]
		c    sets.Set[int]
	}{
		{
			name: "equal sets no change",
			a:    sets.From(1, 2, 3),
			b:    sets.From(1, 2, 3),
			c:    sets.From(1, 2, 3),
		},
		{
			name: "different sets only one common",
			a:    sets.From(1, 2, 3),
			b:    sets.From(1, 5, 6),
			c:    sets.From(1),
		},
		{
			name: "different sets two common",
			a:    sets.From(1, 2, 3),
			b:    sets.From(1, 2, 6),
			c:    sets.From(1, 2),
		},
		{
			name: "different sets nothing common",
			a:    sets.From(1, 2, 3),
			b:    sets.From(4, 5, 6),
			c:    sets.Set[int]{},
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			d := sets.Union(c.a, c.b)

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
			a:     sets.From(1, 2, 3),
			b:     sets.From(1, 2, 3),
			equal: true,
		},
		{
			name:  "different sets",
			a:     sets.From(1, 2, 3),
			b:     sets.From(4, 5, 6),
			equal: false,
		},
		{
			name:  "different order sets",
			a:     sets.From(1, 2, 3),
			b:     sets.From(3, 2, 1),
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
