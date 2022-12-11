package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start, end int
}

func newRange(start, end int) (*Range, error) {
	if end < start {
		return nil, errors.New("End must be greater than start")
	}
	return &Range{
		start: start,
		end:   end,
	}, nil
}

func (r *Range) contains(other *Range) bool {
	return r.start <= other.start && r.end >= other.end
}

func countFullyContained(s *bufio.Scanner) int {
	total := 0
	for s.Scan() {
		r1, r2 := parseLine(s.Text())
		if r1.contains(r2) || r2.contains(r1) {
			total++
		}
	}
	return total
}

func parseLine(l string) (*Range, *Range) {
	range_strings := strings.Split(l, ",")

	var result []*Range
	for _, r := range range_strings {
		bounds := strings.Split(r, "-")
		start, err := strconv.Atoi(bounds[0])
		if err != nil {
			log.Fatal(err)
		}
		end, err := strconv.Atoi(bounds[1])
		if err != nil {
			log.Fatal(err)
		}
		r2, err := newRange(start, end)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, r2)
	}
	return result[0], result[1]
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := countFullyContained(scanner)
	fmt.Println("Total: ", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
