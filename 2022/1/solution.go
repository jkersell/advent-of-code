package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type topNCache struct {
	n     int
	cache []int
}

func newTopNCache(n int) *topNCache {
	return &topNCache{
		n:     n,
		cache: make([]int, 0, n),
	}
}

func (c *topNCache) push(value int) {
	if len(c.cache) < c.n {
		c.cache = append(c.cache, value)
	} else if value > c.cache[0] {
		c.cache[0] = value
	}
	sort.Ints(c.cache)
}

func readInput(lines chan string) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines <- scanner.Text()
	}

	close(lines)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func sumInts(ints []int) int {
	var result int
	for _, val := range ints {
		result += val
	}
	return result
}

func main() {
	lines := make(chan string)

	go readInput(lines)

	elfTotal := 0
	largestSoFar := newTopNCache(3)
	for l := range lines {
		if l == "" {
			largestSoFar.push(elfTotal)
			elfTotal = 0
			continue
		}

		calories, err := strconv.Atoi(l)
		elfTotal = elfTotal + calories
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(sumInts(largestSoFar.cache))
}
