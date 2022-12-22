package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func uniqueCharsFinder(length int) func(r rune) (int, bool) {
	window := make([]rune, length)
	cache := make(map[rune]int)
	i := 0
	return func(r rune) (int, bool) {
		idx := i % length
		to_evict := window[idx]
		window[idx] = r

		if count, ok := cache[to_evict]; ok {
			if count > 0 {
				cache[to_evict]--
			} else if count == 0 {
				delete(cache, to_evict)
			} else {
				log.Fatal("Cache count should not be negative")
			}
		}

		if _, ok := cache[r]; ok {
			cache[r]++
		} else {
			cache[r] = 0
		}

		if len(cache) == length {
			return i + 1, true
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
	messageFinder := uniqueCharsFinder(14)

	packetFound := false
	packetIdx := 0
	for scanner.Scan() {
		r := []rune(scanner.Text())[0]

		if !packetFound {
			var ok bool
			packetIdx, ok = packetFinder(r)
			packetFound = packetFound || ok
		}

		messageIdx, messageFound := messageFinder(r)

		if packetFound && messageFound {
			fmt.Printf(
				"Packet found at: %d\nMessage found at: %d\n",
				packetIdx,
				messageIdx,
			)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
