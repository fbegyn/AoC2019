package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	input = kingpin.Arg("input file", "file to read").Default("input.txt").String()
	steps = kingpin.Arg("time steps", "amount of time steps to take").Default("1000").Int()
)

func main() {
	kingpin.Version("0.1.0")
	kingpin.Parse()

	fi := OpenFile(*input)
	defer fi.Close()

	b, err := ioutil.ReadAll(fi)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	text := strings.TrimSpace(string(b))
	program := strings.Split(text, ",")

	opcodes := make([]int, len(program))

	for i, s := range program {
		code, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Failed to convert string to int: %v", err)
		}
		opcodes[i] = code
	}

	part1 := make([]int, len(opcodes)+4)
	copy(part1, opcodes)
	setParam(12, 2, part1)
	runProgram(part1)
	fmt.Printf("Solution part 1: %d\n", part1[0])

	// part 2
	for i := 1; i < 100; i++ {
		for j := 41; i < 100; i++ {
			part2 := make([]int, len(opcodes)+4)
			copy(part2, opcodes)
			setParam(i, j, part2)
			runProgram(part2)
			if part2[0] == 19690720 {
				fmt.Printf("Noun: %d\nVerb:%d\nSolution part 2: %d\n", i, j, 100*i+j)
				break
			}
		}
	}
	fmt.Println("Done")
}

func setParam(noun, verb int, program []int) []int {
	program[1] = noun
	program[2] = verb
	return program
}

func runProgram(program []int) []int {
	start := 0
	for runInstruction(&program, start) {
		start += 4
	}
	return program
}

func runInstruction(program *[]int, index int) bool {
	code := (*program)[index]
	switch code {
	case 1:
		opA := (*program)[index+1]
		opB := (*program)[index+2]
		dest := (*program)[index+3]
		(*program)[dest] = (*program)[opA] + (*program)[opB]
	case 2:
		opA := (*program)[index+1]
		opB := (*program)[index+2]
		dest := (*program)[index+3]
		(*program)[dest] = (*program)[opA] * (*program)[opB]
	case 99:
		return false
	}
	return true
}
