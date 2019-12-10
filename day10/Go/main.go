package main

import (
	"bufio"
	"fmt"
	"math"
	"sort"
	"strings"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	input = kingpin.Arg("input file", "file to read").Default("input.txt").String()
)

type Point struct {
	x int
	y int
}

func (p Point) angle(t Point) (angle float64) {
	angle = math.Atan2(float64(t.x-p.x), float64(t.y-p.y)) * 180 / math.Pi
	if angle < 0 {
		angle += 360
	}
	return
}

func (p Point) distance(t Point) int {
	return Abs(p.x-t.x) + Abs(p.y-t.y)
}

func main() {
	kingpin.Version("0.1.0")
	kingpin.Parse()

	astMap := genAstMap(*input)

	part1(astMap)
	station := Point{}
	mostAst := -1
	for k, v := range astMap {
		if mostAst < v {
			mostAst = v
			station = k
		}
	}
	fmt.Printf("The station is located at %v.\n", station)
	fmt.Printf("From there it see %d asteroids.\n", mostAst)

	bet := 200
	hitAst := part2(astMap, station, bet)
	fmt.Printf("The %d asteroid to be hit: %d\n", bet, hitAst)
}

func genAstMap(fiDsc string) map[Point]int {
	fi := OpenFile(fiDsc)
	defer fi.Close()

	astMap := make(map[Point]int)

	scanner := bufio.NewScanner(fi)
	y := 0
	for scanner.Scan() {
		sc := scanner.Text()
		line := strings.TrimSpace(sc)
		for x, r := range line {
			if r == '#' {
				astMap[Point{x, y}] = 0
			}
		}
		y++
	}
	return astMap
}

func part1(astMap map[Point]int) {
	for curr, _ := range astMap {
		angles := make(map[float64]bool)
		var count int
		for target, _ := range astMap {
			if curr == target {
				continue
			}
			angle := curr.angle(target)
			if _, ok := angles[angle]; !ok {
				angles[angle] = true
				count++
			}
		}
		astMap[curr] = count
	}
}

func part2(astMap map[Point]int, station Point, bet int) int {
	targets := make(map[float64]map[int]Point)
	var angles []float64
	for target, _ := range astMap {
		if station == target {
			fmt.Println("Skipping station location")
			continue
		}
		angle := target.angle(station)
		if angle == 0 {
			angle += 360
		}
		distance := target.distance(station)

		if _, ok := targets[angle]; !ok {
			targets[angle] = make(map[int]Point)
			angles = append(angles, angle)
		}
		targets[angle][distance] = target
	}

	// sort.Float64s(angles)
	sort.Sort(sort.Reverse(sort.Float64Slice(angles)))
	hit := 0

	for len(angles) > 0 {
		for _, angle := range angles {
			closest := -1

			for dist, _ := range targets[angle] {
				if dist < closest || closest == -1 {
					closest = dist
				}
			}
			hit++
			if hit == bet {
				b := targets[angle][closest]
				return b.x*100 + b.y
			}
			delete(targets[angle], closest)
		}

		var clean []float64
		for _, angle := range angles {
			if len(targets[angle]) > 0 {
				clean = append(clean, angle)
			}
		}
		angles = clean
	}
	return -1
}
