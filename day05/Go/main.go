package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	fi := OpenFile("./input.txt")
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
	runProgram(part1)
	// fmt.Printf("%d\n", part1)
	// fmt.Println("Done")
}

func setParam(noun, verb int, program []int) []int {
	program[1] = noun
	program[2] = verb
	return program
}

func runProgram(program []int) []int {
	start := 0
	run, jump := runInstruction(&program, start)
	for run {
		start += jump
		run, jump = runInstruction(&program, start)
	}
	return program
}

func runInstruction(program *[]int, index int) (run bool, jump int) {
	opcode := (*program)[index] % 100
	modesNumber := (*program)[index] / 100
	modes := make([]int, 3)
	i := 0
	for modesNumber > 0 {
		modes[i] = modesNumber % 10
		modesNumber /= 10
		i += 1
	}
	switch opcode {
	case 1:
		params := getParam((*program)[index:], 3)
		opA := params[0]
		opB := params[1]
		dest := params[2]
		if modes[0] == 0 {
			opA = (*program)[opA]
		}
		if modes[1] == 0 {
			opB = (*program)[opB]
		}
		(*program)[dest] = opA + opB
		jump = 4
	case 2:
		params := getParam((*program)[index:], 3)
		opA := params[0]
		opB := params[1]
		dest := params[2]
		if modes[0] == 0 {
			opA = (*program)[opA]
		}
		if modes[1] == 0 {
			opB = (*program)[opB]
		}
		(*program)[dest] = opA * opB
		jump = 4
	case 3:
		params := getParam((*program)[index:], 1)
		dest := params[0]
		fmt.Print("Enter ID to run diagnostics: ")
		var id int
		_, err := fmt.Scanf("%d", &id)
		if err != nil {
			log.Fatalf("Could not convert string to id: %v\n", err)
		}
		(*program)[dest] = id
		jump = 2
	case 4:
		params := getParam((*program)[index:], 1)
		src := params[0]
		if modes[0] == 0 {
			src = (*program)[src]
		}
		fmt.Printf("Current test output : %d\n", src)
		jump = 2
	case 5:
		params := getParam((*program)[index:], 2)
		opA := params[0]
		opB := params[1]
		if modes[0] == 0 {
			opA = (*program)[opA]
		}
		if modes[1] == 0 {
			opB = (*program)[opB]
		}
		if opA != 0 {
			jump = opB - index
		} else {
			jump = 3
		}
	case 6:
		params := getParam((*program)[index:], 2)
		opA := params[0]
		opB := params[1]
		if modes[0] == 0 {
			opA = (*program)[opA]
		}
		if modes[1] == 0 {
			opB = (*program)[opB]
		}
		if opA == 0 {
			jump = opB - index
		} else {
			jump = 3
		}
	case 7:
		params := getParam((*program)[index:], 3)
		opA := params[0]
		opB := params[1]
		dest := params[2]
		if modes[0] == 0 {
			opA = (*program)[opA]
		}
		if modes[1] == 0 {
			opB = (*program)[opB]
		}
		if opA < opB {
			(*program)[dest] = 1
		} else {
			(*program)[dest] = 0
		}
		jump = 4
	case 8:
		params := getParam((*program)[index:], 3)
		opA := params[0]
		opB := params[1]
		dest := params[2]
		if modes[0] == 0 {
			opA = (*program)[opA]
		}
		if modes[1] == 0 {
			opB = (*program)[opB]
		}
		if opA == opB {
			(*program)[dest] = 1
		} else {
			(*program)[dest] = 0
		}
		jump = 4
	case 99:
		return false, 1
	default:
		return false, 0
	}
	return true, jump
}

func getParam(program []int, param int) []int {
	params := make([]int, param)
	for i := 0; i < param; i++ {
		params[i] = program[i+1]
	}
	return params
}
