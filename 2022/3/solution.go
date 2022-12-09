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

func main() {
	total := 0
	for l := range readInput() {
		fmt.Println(l)
		firstPocket := make(map[rune]bool)
		mid := utf8.RuneCountInString(l) / 2
		for i, r := range l {
			if i < mid {
				firstPocket[r] = true
			} else {
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
		}
		fmt.Println("")
	}
	fmt.Println(total)
}
