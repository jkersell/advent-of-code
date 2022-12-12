package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Stack []rune

func (s Stack) push(r rune) Stack {
	return append(s, r)
}

func (s Stack) pop() (Stack, rune, error) {
	l := len(s)
	if l == 0 {
		return s, 0, errors.New("Can not pop from empty stack")
	}
	return s[:l-1], s[l-1], nil
}

func (s Stack) peak() (rune, error) {
	l := len(s)
	if l == 0 {
		return 0, errors.New("Can not peak on empty stack")
	}
	return s[l-1], nil
}

func runSimulation(scanner *bufio.Scanner) string {
	// [N]         [C]     [Z]
	// [Q] [G]     [V]     [S]         [V]
	// [L] [C]     [M]     [T]     [W] [L]
	// [S] [H]     [L]     [C] [D] [H] [S]
	// [C] [V] [F] [D]     [D] [B] [Q] [F]
	// [Z] [T] [Z] [T] [C] [J] [G] [S] [Q]
	// [P] [P] [C] [W] [W] [F] [W] [J] [C]
	// [T] [L] [D] [G] [P] [P] [V] [N] [R]
	// 1   2   3   4   5   6   7   8   9
	var stacks = []Stack{
		Stack{'T', 'P', 'Z', 'C', 'S', 'L', 'Q', 'N'},
		Stack{'L', 'P', 'T', 'V', 'H', 'C', 'G'},
		Stack{'D', 'C', 'Z', 'F'},
		Stack{'G', 'W', 'T', 'D', 'L', 'M', 'V', 'C'},
		Stack{'P', 'W', 'C'},
		Stack{'P', 'F', 'J', 'D', 'C', 'T', 'S', 'Z'},
		Stack{'V', 'W', 'G', 'B', 'D'},
		Stack{'N', 'J', 'S', 'Q', 'H', 'W'},
		Stack{'R', 'C', 'Q', 'F', 'S', 'L', 'V'},
	}

	for scanner.Scan() {
		count, from, to := parseInstruction(scanner.Text())
		execute(stacks, count, from, to)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return readResult(stacks)
}

func parseInstruction(l string) (int, int, int) {
	pattern, err := regexp.Compile("^move (\\d+) from (\\d+) to (\\d)+$")
	if err != nil {
		log.Fatal(err)
	}

	matches := pattern.FindStringSubmatch(l)[1:]
	if len(matches) != 3 {
		log.Fatal("Failed to parse an instruction")
	}

	var result = make([]int, 3)
	for i := range matches {
		result[i], err = strconv.Atoi(matches[i])
		if err != nil {
			log.Fatal(err)
		}
	}
	return result[0], result[1], result[2]
}

func execute(stacks []Stack, count, from, to int) {
	fromStack := stacks[from-1]
	toStack := stacks[to-1]

	for i := 0; i < count; i++ {
		s, item, err := fromStack.pop()
		fromStack = s
		if err != nil {
			break
		}
		toStack = toStack.push(item)
	}
	stacks[from-1] = fromStack
	stacks[to-1] = toStack
}

func readResult(stacks []Stack) string {
	var result string
	for _, s := range stacks {
		item, err := s.peak()
		if err != nil {
			continue
		}
		result = result + string(item)
	}
	return result
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := runSimulation(scanner)

	fmt.Println(result)
}
