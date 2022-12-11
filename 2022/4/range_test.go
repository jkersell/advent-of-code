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

func TestRangeOverlaps(t *testing.T) {
	r1 := newRange(0, 10)
	var cases = []struct {
		name     string
		r1, r2   Range
		expected bool
	}{
		{"Fully contained", r1, newRange(1, 9), true},
		{"Fully contained reversed", newRange(1, 9), r1, true},
		{"Equal", r1, newRange(r1.start, r1.end), true},
		{"Overlapping lower bound", r1, newRange(-5, 5), true},
		{"Overlapping lower bound reversed", newRange(-5, 5), r1, true},
		{"R2 ends where R1 starts", r1, newRange(-5, r1.start), true},
		{"Overlapping upper bound", r1, newRange(5, 15), true},
		{"R2 starts where R1 ends", r1, newRange(r1.start, 15), true},
		{"No intersection", r1, newRange(15, 20), false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := c.r1.overlaps(&c.r2)

			assert.Equal(t, c.expected, actual)
		})
	}
}
