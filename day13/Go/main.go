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
)

func main() {
	kingpin.Version("0.1.0")
	kingpin.Parse()

	// input parsing
	fi := OpenFile(*input)
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

	// part 1
	part1(program)
	part2(program)
}

func part1(prog []int64) {
	grid := make(map[Point]int64)
	halt := make(chan bool)
	end := make(chan bool)
	in, out := make(chan int64, 1), make(chan int64)
	go RunProgram(prog, in, out, halt)
	go func() {
		for {
			select {
			case <-halt:
				end <- true
				return
			case x := <-out:
				y := <-out
				v := <-out
				p := Point{x, y}
				grid[p] = v
			default:
			}
		}
	}()
	<-end
	blocks := 0
	for _, v := range grid {
		if v == 2 {
			blocks++
		}
	}
	fmt.Printf("There are %d blocks on the screen.\n", blocks)
}

func part2(prog []int64) {
	// set to free play mode
	prog[0] = 2

	// initialize program channels
	grid := make(map[Point]int64)
	score := int64(0)
	halt := make(chan bool)
	end := make(chan bool)
	in, out := make(chan int64, 1), make(chan int64)
	// channel to indictatie player turn
	turn := make(chan bool)

	// initialize ai vars
	ballLocation := Point{0, 0}
	paddleLocation := Point{0, 0}

	go RunProgram(prog, in, out, halt)
	go func() {
		for {
			select {
			case <-halt:
				end <- true
				return
			case x := <-out:
				y := <-out
				v := <-out
				if x == -1 && y == 0 {
					score = v
				} else {
					p := Point{x, y}
					grid[p] = v
					if v == 3 {
						paddleLocation.x = x
						paddleLocation.y = y
					} else if v == 4 {
						ballLocation.x = x
						ballLocation.y = y
						turn <- true
					}
				}
			default:
				continue
			}
		}
	}()
	go func() {
		for {
			select {
			case <-turn:
				if ballLocation.x < paddleLocation.x {
					in <- -1
				} else if paddleLocation.x < ballLocation.x {
					in <- 1
				} else {
					in <- 0
				}
			default:
				continue
			}
		}
	}()
	<-end
	fmt.Printf("End score: %d\n", score)
}
