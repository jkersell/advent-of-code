package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueCharsFinder_FindsPacketMarker(t *testing.T) {
	var cases = []struct {
		stream   string
		expected int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			messageFinder := uniqueCharsFinder(4)
			for _, r := range c.stream {
				if actual, ok := messageFinder(r); ok {
					assert.Equal(t, c.expected, actual)
					break
				}
			}
		})
	}
}

func TestUniqueCharsFinder_FindsMessageMarker(t *testing.T) {
	var cases = []struct {
		stream   string
		expected int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}

	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			messageFinder := uniqueCharsFinder(14)
			for _, r := range c.stream {
				if actual, ok := messageFinder(r); ok {
					assert.Equal(t, c.expected, actual)
					break
				}
			}
		})
	}
}
