package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"unicode/utf8"
)

func readInput() <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		file, err := os.Open("input.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			out <- scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()
	return out
}

func priority(r rune) (int, error) {
	asInt := int(r)
	if asInt >= 97 && asInt <= 123 {
		return asInt - 96, nil
	} else if asInt >= 65 && asInt <= 91 {
		return asInt - 38, nil
	} else {
		return 0, errors.New("Rune out of range")
	}
}

func buildCache(s string) map[rune]bool {
	cache := make(map[rune]bool)
	for _, r := range s {
		cache[r] = true
	}
	return cache
}

func main() {
	total := 0
	for l := range readInput() {
		fmt.Println(l)
		mid := utf8.RuneCountInString(l) / 2
		firstPocket := buildCache(l[:mid])
		for _, r := range l[mid:] {
			_, ok := firstPocket[r]
			if ok {
				fmt.Println("Seeing an item again: ", string(r))
				pri, err := priority(r)
				if err != nil {
					log.Fatal(err)
				}
				total += pri
				break
			}
		}
		fmt.Println("")
	}
	fmt.Println(total)
}
