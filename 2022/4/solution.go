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

func (r *Range) overlaps(other *Range) bool {
	return r.contains(other) || other.contains(r) || r.overlap_lower(other) || r.overlap_upper(other)
}

func (r *Range) overlap_lower(other *Range) bool {
	return r.start <= other.start && r.end >= other.start && r.end <= other.end
}

func (r *Range) overlap_upper(other *Range) bool {
	return r.start >= other.start && r.start <= other.end && r.end >= other.end
}

func trueCounter() func(bool) int {
	total := 0
	return func(p bool) int {
		if p {
			total++
		}
		return total
	}
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

	fullyContainedCounter := trueCounter()
	fullyContainedTotal := 0
	overlappingCounter := trueCounter()
	overlappingTotal := 0
	for scanner.Scan() {
		r1, r2 := parseLine(scanner.Text())
		fullyContainedTotal = fullyContainedCounter(r1.contains(r2) || r2.contains(r1))
		overlappingTotal = overlappingCounter(r1.overlaps(r2))
	}
	fmt.Println("Total fully contained ranges: ", fullyContainedTotal)
	fmt.Println("Total overlapping ranges: ", overlappingTotal)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
