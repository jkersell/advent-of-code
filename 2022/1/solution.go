package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

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

func main() {
	lines := make(chan string)

	go readInput(lines)

	elfTotal := 0
	largestSoFar := 0
	for l := range lines {
		if l == "" {
			if elfTotal > largestSoFar {
				largestSoFar = elfTotal
			}
			elfTotal = 0
			continue
		}

		calories, err := strconv.Atoi(l)
		elfTotal = elfTotal + calories
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(largestSoFar)
}
