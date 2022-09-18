package arrays_test

import (
	"strings"
	"testing"

	"github.com/blue-health/blue-go-toolbox/arrays"
	"github.com/stretchr/testify/require"
)

func TestContainsAny(t *testing.T) {
	cases := []struct {
		hay      []string
		needle   string
		contains bool
	}{
		{
			hay:      []string{},
			needle:   "",
			contains: false,
		},
		{
			hay:      []string{""},
			needle:   "",
			contains: true,
		},
		{
			hay:      []string{"hello", "world"},
			needle:   "he",
			contains: false,
		},
		{
			hay:      []string{"hello", "world"},
			needle:   "world",
			contains: true,
		},
	}

	for _, c := range cases {
		t.Run(strings.Join(c.hay, ",")+":"+c.needle, func(t *testing.T) {
			if arrays.ContainsAny(c.hay, c.needle) != c.contains {
				t.Fatalf("Contains returned false bool. Haystack: %v, Needle: %s, Expected: %t", c.hay, c.needle, c.contains)
			}
		})
	}
}

func TestContainsAll(t *testing.T) {
	cases := []struct {
		super  []string
		sub    []string
		result bool
	}{
		{
			super:  []string{},
			sub:    []string{},
			result: true,
		},
		{
			super:  []string{},
			sub:    []string{""},
			result: false,
		},
		{
			super:  []string{""},
			sub:    []string{""},
			result: true,
		},
		{
			super:  []string{"hello", "world"},
			sub:    []string{"wo"},
			result: false,
		},
		{
			super:  []string{"hello", "world"},
			sub:    []string{"world"},
			result: true,
		},
		{
			super:  []string{"hello", "world"},
			sub:    []string{"world", "hi"},
			result: false,
		},
		{
			super:  []string{"hello", "world"},
			sub:    []string{"hello", "world"},
			result: true,
		},
		{
			super:  []string{"hello", "world", "there"},
			sub:    []string{"hello", "there"},
			result: true,
		},
		{
			super:  []string{"hello", "world", "there"},
			sub:    []string{"hello", "world", "there", "earth"},
			result: false,
		},
	}

	for _, c := range cases {
		t.Run(strings.Join(c.super, ",")+":"+strings.Join(c.sub, ","), func(t *testing.T) {
			if arrays.ContainsAll(c.super, c.sub...) != c.result {
				t.Fatalf("Subset returned false bool. Superset: %v, Subset: %s, Expected: %t", c.super, c.sub, c.result)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	cases := []struct {
		in  []string
		re  []string
		out []string
	}{
		{
			in:  []string{},
			re:  []string{},
			out: []string{},
		},
		{
			in:  []string{},
			re:  []string{""},
			out: []string{},
		},
		{
			in:  []string{""},
			re:  []string{""},
			out: []string{},
		},
		{
			in:  []string{""},
			re:  []string{},
			out: []string{""},
		},
		{
			in:  []string{"hello", "world"},
			re:  []string{"wo"},
			out: []string{"hello", "world"},
		},
		{
			in:  []string{"hello", "world"},
			re:  []string{"world"},
			out: []string{"hello"},
		},
		{
			in:  []string{"hello", "world"},
			re:  []string{"hello"},
			out: []string{"world"},
		},
		{
			in:  []string{"hello", "world"},
			re:  []string{"world", "hello"},
			out: []string{},
		},
	}

	for _, c := range cases {
		t.Run(strings.Join(c.in, ",")+":"+strings.Join(c.out, ","), func(t *testing.T) {
			require.ElementsMatch(t, c.out, arrays.Remove(c.in, c.re...))
		})
	}
}
