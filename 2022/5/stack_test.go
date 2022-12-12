package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackPush_AddsAnElementToTheStack(t *testing.T) {
	var cases = []struct {
		name     string
		start    Stack
		r        rune
		expected Stack
	}{
		{"Empty stack", Stack{}, 'a', Stack{'a'}},
		{"Small stack", Stack{'a', 'b', 'c'}, 'a', Stack{'a', 'b', 'c', 'a'}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := c.start.push(c.r)

			assert.Equal(t, c.expected, actual)
		})
	}
}

func TestStackPop_ReturnsAnErrorWhenEmpty(t *testing.T) {
	s := Stack{}

	_, _, err := s.pop()

	assert.EqualError(t, err, "Can not pop from empty stack")
}

func TestStackPop_RemovesTheLastElementFromTheStack(t *testing.T) {
	var cases = []struct {
		name          string
		start         Stack
		expectedStack Stack
		expectedRune  rune
	}{
		{"Single element stack", Stack{'a'}, Stack{}, 'a'},
		{"Small stack", Stack{'a', 'b', 'c'}, Stack{'a', 'b'}, 'c'},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actualStack, actualRune, err := c.start.pop()

			assert.NoError(t, err)
			assert.Equal(t, c.expectedRune, actualRune)
			assert.Equal(t, c.expectedStack, actualStack)
		})
	}
}
