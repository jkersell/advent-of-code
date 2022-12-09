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

func part1() {
	total := 0
	for l := range readInput() {
		fmt.Println(l)
		pri, err := checkRuck(l)
		if err != nil {
			log.Fatal(err)
		}
		total += pri
		fmt.Println("")
	}
	fmt.Println(total)
}

func part2() {
	total := 0
	i := 0
	ruck := make(map[rune]bool)
	for l := range readInput() {
		fmt.Println(l)
		i++
		if i == 1 {
			ruck = buildCache(l)
		}
		ruck = intersect(ruck, l)
		fmt.Println(ruck)
		if i == 3 {
			if len(ruck) > 1 {
				log.Fatal("Expected one duplicate item but there were more")
			} else if len(ruck) < 1 {
				log.Fatal("Did not find a duplicated element")
			}
			for k, _ := range ruck {
				pri, err := priority(k)
				if err != nil {
					log.Fatal(err)
				}
				total += pri
			}
			i = 0
		}
	}
	fmt.Println(total)
}

func checkRuck(ruck string) (int, error) {
	mid := utf8.RuneCountInString(ruck) / 2
	firstPocket := buildCache(ruck[:mid])
	intersection := intersect(firstPocket, ruck[mid:])
	if len(intersection) > 1 {
		return 0, errors.New("Expected one duplicate item but there were more")
	} else if len(intersection) < 1 {
		return 0, errors.New("Did not find a duplicated element")
	}
	for k, _ := range intersection {
		pri, err := priority(k)
		if err != nil {
			log.Fatal(err)
		}
		return pri, nil
	}
	return 0, errors.New("Unexpected error while checking a ruck")
}

func intersect(cache map[rune]bool, ruck string) map[rune]bool {
	intersection := make(map[rune]bool)
	for _, r := range ruck {
		_, ok := cache[r]
		if ok {
			intersection[r] = true
		}
	}
	return intersection
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
	part2()
}
