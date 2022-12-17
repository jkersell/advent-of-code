package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func packetStartIndex(scanner *bufio.Scanner) int {
	capacity := 4
	window := make([]rune, 4)
	cache := make(map[rune]bool)
	i := 0
	for scanner.Scan() {
		r := []rune(scanner.Text())[0]
		idx := i % capacity
		to_evict := window[idx]
		window[idx] = r
		delete(cache, to_evict)
		cache[r] = true
		if len(cache) == capacity {
			return i
		}
		i++
	}
	return -1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	fmt.Println(packetStartIndex(scanner))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
