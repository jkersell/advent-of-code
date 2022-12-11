package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeContains(t *testing.T) {
	r1 := newRange(0, 10)
	var cases = []struct {
		name     string
		r1, r2   Range
		expected bool
	}{
		{"Fully contained", r1, newRange(1, 9), true},
		{"Equal", r1, newRange(r1.start, r1.end), true},
		{"Overlapping lower bound", r1, newRange(-5, 5), false},
		{"Overlapping upper bound", r1, newRange(5, 15), false},
		{"No intersection", r1, newRange(15, 20), false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := c.r1.contains(&c.r2)

			assert.Equal(t, c.expected, actual)
		})
	}
}
