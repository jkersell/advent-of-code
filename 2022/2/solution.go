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
	Draw = 0
	Win  = 1
	Lose = 2
)

const (
	Rock     = 0
	Paper    = 1
	Scissors = 2
)

var p1Cipher = map[string]Play{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
}

var outcomeCipher = map[string]Outcome{
	"X": Lose,
	"Y": Draw,
	"Z": Win,
}

var playScores = map[Play]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

var outcomeScores = map[Outcome]int{
	Lose: 0,
	Draw: 3,
	Win:  6,
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

func myPlay(p1 Play, o Outcome) Play {
	if o == Win {
		return Play(mod(int(p1+1), 3))
	} else if o == Lose {
		return Play(mod(int(p1-1), 3))
	}
	return p1
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

func decipherRound(line string) (Play, Outcome) {
	plays := strings.Split(line, " ")
	p1 := p1Cipher[plays[0]]
	o := outcomeCipher[plays[1]]
	return p1, o
}

func main() {
	lines := readInput()

	totalScore := 0
	for l := range lines {
		p1, outcome := decipherRound(l)
		p2 := myPlay(p1, outcome)
		if beats(p1, p2) != outcome {
			fmt.Println("Wrong result")
		}
		roundScore := playScores[p2] + outcomeScores[outcome]
		totalScore += roundScore
	}
	fmt.Println("Total score: ", totalScore)
}
