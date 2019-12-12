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
	input                        = kingpin.Arg("input file", "file to read").Default("input.txt").String()
	steps                        = kingpin.Arg("time steps", "amount of time steps to take").Default("1000").Int()
	directionMapX map[string]int = map[string]int{"U": 0, "D": 0, "L": 1, "R": -1}
	directionMapY map[string]int = map[string]int{"U": 1, "D": -1, "L": 0, "R": -0}
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
	program := strings.Split(text, "\n")
	wire1 := strings.Split(program[0], ",")
	wire2 := strings.Split(program[1], ",")
	wirePoint1 := GetPoints(wire1)
	wirePoint2 := GetPoints(wire2)
	overlap := Overlap(wirePoint1, wirePoint2)
	minDist := 999999999

	fewestSteps := 999999999
	for k := range overlap {
		dist := ManhattanDistance([]int{0, 0}, []int{k.x, k.y})
		if dist < minDist {
			minDist = dist
		}
		totalSteps := wirePoint1[k] + wirePoint2[k]
		if totalSteps < fewestSteps {
			fewestSteps = totalSteps
		}
	}
	fmt.Printf("Manhattan distance of the closest point (part1): %d\n", minDist)
	fmt.Printf("Fewest steps (part2): %d\n", fewestSteps)
}

type point struct {
	x, y int
}

func GetPoints(moves []string) (ans map[point]int) {
	x, y := 0, 0
	steps := 0
	ans = make(map[point]int)
	for _, m := range moves {
		direction := m[:1]
		dist, err := strconv.Atoi(m[1:])
		if err != nil {
			log.Fatalf("Failed to convert string to int: %s\n", m[1:])
		}
		for i := 0; i < dist; i++ {
			x += directionMapX[direction]
			y += directionMapY[direction]
			steps += 1
			if _, ok := ans[point{x, y}]; !ok {
				ans[point{x, y}] = steps
			}
		}
	}
	return
}

func Overlap(a, b map[point]int) (ans map[point]int) {
	ans = make(map[point]int)
	for k, v := range a {
		if _, ok := b[k]; ok {
			ans[k] = v
		}
	}
	return
}

func part1() {
}

func part2() {
}
