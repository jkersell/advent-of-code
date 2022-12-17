package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func uniqueCharsFinder(length int) func(r rune) (int, bool) {
	window := make([]rune, 4)
	cache := make(map[rune]bool)
	i := 0
	return func(r rune) (int, bool) {
		idx := i % length
		to_evict := window[idx]
		window[idx] = r
		delete(cache, to_evict)
		cache[r] = true
		if len(cache) == length {
			return i, true
		}
		i++
		return 0, false
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	packetFinder := uniqueCharsFinder(4)

	for scanner.Scan() {
		r := []rune(scanner.Text())[0]
		if i, ok := packetFinder(r); ok {
			fmt.Println(i)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
