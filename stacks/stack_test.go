package stacks_test

import (
	"testing"

	"github.com/blue-health/blue-go-toolbox/stacks"
	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	testCases := []struct {
		name       string
		in         []int
		out        []int
		overextend bool
	}{
		{
			name: "push 5, pop 5",
			in:   []int{1, 2, 3, 4, 5},
			out:  []int{5, 4, 3, 2, 1},
		},
		{
			name: "push 5, pop 3",
			in:   []int{1, 2, 3, 4, 5},
			out:  []int{5, 4, 3},
		},
		{
			name:       "push 3, pop 5",
			in:         []int{1, 2, 3},
			out:        []int{3, 2, 1, 0, 0},
			overextend: true,
		},
		{
			name: "push 0, pop 0",
			in:   []int{},
			out:  []int{},
		},
		{
			name: "push 10, pop 10",
			in:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			out:  []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			s := stacks.Stack[int]{}

			for _, i := range c.in {
				s.Push(i)
			}

			for _, i := range c.in {
				require.True(t, s.Contains(i))
			}

			s.Iterate(func(i int) {
				require.Contains(t, c.in, i)
			})

			for _, i := range c.out {
				o, ok := s.Pop()
				if !c.overextend {
					require.True(t, ok)
				}
				require.Equal(t, i, o)
			}
		})
	}
}
