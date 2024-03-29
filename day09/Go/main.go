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

	halt := make(chan bool)
	in, out := make(chan int64, 1), make(chan int64)
	go RunProgram(program, in, out, halt)
	in <- 2
	go func() {
		fmt.Printf("Test output: ")
		for o := range out {
			fmt.Printf("%v ", o)
		}
	}()
	fmt.Println(<-out)
	for i := 0; i < 1; i++ {
		<-halt
	}
}
