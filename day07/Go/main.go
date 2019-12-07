package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	// input parsing
	fi := OpenFile("./input.txt")
	defer fi.Close()

	b, err := ioutil.ReadAll(fi)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	text := strings.TrimSpace(string(b))
	programStr := strings.Split(text, ",")

	program := make([]int, len(programStr))

	for i, s := range programStr {
		code, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Failed to convert string to int: %v", err)
		}
		program[i] = code
	}

	permuts := AllPermutations([]int{0, 1, 2, 3, 4})
	signal := 0
	for _, phase := range permuts {
		sign := runAmps(program, phase)
		if sign > signal {
			signal = sign
		}
	}
	fmt.Printf("Part 1: %d\n", signal)

	permuts = AllPermutations([]int{5, 6, 7, 8, 9})
	signal = 0
	for _, phase := range permuts {
		sign := runAmps(program, phase)
		if sign > signal {
			signal = sign
		}
	}
	fmt.Printf("Part 2: %d\n", signal)
}

func runAmps(program []int, phase []int) int {
	halt := make(chan bool)
	inAmp1, inAmp2, inAmp3, inAmp4, inAmp5 := make(chan int, 1), make(chan int), make(chan int),
		make(chan int), make(chan int)

	go RunProgram(program, inAmp1, inAmp2, halt)
	go RunProgram(program, inAmp2, inAmp3, halt)
	go RunProgram(program, inAmp3, inAmp4, halt)
	go RunProgram(program, inAmp4, inAmp5, halt)
	go RunProgram(program, inAmp5, inAmp1, halt)

	// set phase settings
	inAmp1 <- phase[0]
	inAmp2 <- phase[1]
	inAmp3 <- phase[2]
	inAmp4 <- phase[3]
	inAmp5 <- phase[4]

	// input signal
	inAmp1 <- 0

	// wait for completion
	for i := 0; i < 5; i++ {
		<-halt
	}

	return <-inAmp1
}
