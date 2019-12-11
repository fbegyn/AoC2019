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

	program := make([]int64, len(programStr)+1000)

	for i, s := range programStr {
		code, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			log.Fatalf("Failed to convert string to int: %v", err)
		}
		program[i] = code
	}

	grid := make(map[Point]int64)
	start := Point{0, 0}
	grid[start] = 1
	halt := make(chan bool)
	in, out := make(chan int64, 1), make(chan int64)
	go RunRobot(grid, start, out, in)
	go RunProgram(program, in, out, halt)
	for i := 0; i < 1; i++ {
		<-halt
	}
	fmt.Printf("The robot will paint %d tiles.\n", len(grid))
	fmt.Println("Rendering image:")
	image := RenderGrid(grid)
	PrintImage(image)
}
