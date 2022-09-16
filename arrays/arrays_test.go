package arrays_test

import (
	"strings"
	"testing"

	"github.com/blue-health/blue-go-toolbox/arrays"
)

func TestContains(t *testing.T) {
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
		if arrays.Contains(c.hay, c.needle) != c.contains {
			t.Fatalf("Contains returned false bool. Haystack: %v, Needle: %s, Expected: %t", c.hay, c.needle, c.contains)
		}
	}
}

func TestSubset(t *testing.T) {
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
			if arrays.Subset(c.super, c.sub...) != c.result {
				t.Fatalf("Subset returned false bool. Superset: %v, Subset: %s, Expected: %t", c.super, c.sub, c.result)
			}
		})
	}
}

func TestIntersects(t *testing.T) {
	cases := []struct {
		one        []string
		two        []string
		intersects bool
	}{
		{
			one:        []string{},
			two:        []string{},
			intersects: false,
		},
		{
			one:        []string{},
			two:        []string{""},
			intersects: false,
		},
		{
			one:        []string{""},
			two:        []string{""},
			intersects: true,
		},
		{
			one:        []string{"hello", "world"},
			two:        []string{"wo"},
			intersects: false,
		},
		{
			one:        []string{"hello", "world"},
			two:        []string{"world"},
			intersects: true,
		},
		{
			one:        []string{"hello", "world"},
			two:        []string{"world", "hi"},
			intersects: true,
		},
		{
			one:        []string{"hello", "world"},
			two:        []string{"hello", "world"},
			intersects: true,
		},
		{
			one:        []string{"hello", "world", "there"},
			two:        []string{"hello", "there"},
			intersects: true,
		},
		{
			one:        []string{"hello", "world", "there"},
			two:        []string{"hello", "world", "there", "earth"},
			intersects: true,
		},
	}

	for _, c := range cases {
		if arrays.Intersects(c.one, c.two...) != c.intersects {
			t.Fatalf("Intersects returned false bool. First: %v, Second: %s, Expected: %t", c.one, c.two, c.intersects)
		}
	}
}
