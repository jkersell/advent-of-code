package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Play int
type Outcome int

const (
	Lose = 0
	Draw = 3
	Win  = 6
)

const (
	Rock     = 1
	Paper    = 2
	Scissors = 3
)

var p1Cipher = map[string]Play{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
}

var p2Cipher = map[string]Play{
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

func beats(p1, p2 Play) Outcome {
	if p1 == p2 {
		return Draw
	} else if mod(int(p2-p1), 3) == 1 {
		return Win
	} else {
		return Lose
	}
}

func mod(a, b int) int {
	r := a % b
	if r < 0 && a < 0 {
		r += b
	}
	return r
}

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

func decipherRound(line string) (Play, Play) {
	plays := strings.Split(line, " ")
	p1 := p1Cipher[plays[0]]
	p2 := p2Cipher[plays[1]]
	return p1, p2
}

func main() {
	lines := readInput()

	totalScore := 0
	for l := range lines {
		p1, p2 := decipherRound(l)
		outcome := beats(p1, p2)
		roundScore := int(p2) + int(outcome)
		totalScore += roundScore
	}
	fmt.Println("Total score: ", totalScore)
}
