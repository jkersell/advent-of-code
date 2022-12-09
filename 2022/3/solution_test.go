package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriority(t *testing.T) {
	cases := []struct {
		in   rune
		want int
	}{
		{'a', 1},
		{'A', 27},
		{'z', 26},
		{'Z', 52},
	}

	for _, c := range cases {
		t.Run(string(c.in), func(t *testing.T) {
			actual, err := priority(c.in)
			if err != nil {
				t.Fatal("Unexpected error")
			}

			assert.Equal(t, c.want, actual)
		})
	}
}
