package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

func OpenFile(f string) (file *os.File) {
	file, err := os.Open(f)
	if err != nil {
		log.Fatalf("failed to read file into scanner: %v", err)
	}
	return
}

func SumOfFloat64Array(test []float64) (result float64) {
	for _, v := range test {
		result += v
	}
	return
}

func SumOfIntArray(test []int) (result int) {
	for _, v := range test {
		result += v
	}
	return
}

func ManhattanDistance(x, y []int) int {
	deltax := Abs(x[0] - y[0])
	deltay := Abs(x[1] - y[1])
	return deltax + deltay
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(m map[string]int) (ind string) {
	min := 320000000
	for k, v := range m {
		if v == min {
			ind = "D"
		}
		if v < min {
			min = v
			ind = k
		}
	}
	return
}

func Max(m map[string]int) (ind string) {
	max := 0
	for k, v := range m {
		if v > max {
			max = v
			ind = k
		}
	}
	return
}

func AllPermutations(values []int) (result [][]int) {
	if len(values) == 1 {
		result = append(result, values)
		return
	}
	for i, current := range values {
		others := make([]int, 0, len(values)-1)
		others = append(others, values[:i]...)
		others = append(others, values[i+1:]...)
		for _, route := range AllPermutations(others) {
			result = append(result, append(route, current))
		}
	}
	return
}

func RunProgram(program []int64, input <-chan int64, output chan<- int64, halt chan<- bool) {
	mem := make([]int64, len(program))
	copy(mem, program)
	pc, relBase := int64(0), int64(0)

	for {
		opcode := mem[pc] % 100
		modesNumber := mem[pc] / 100
		modes := make([]int64, 3)
		i := 0
		for modesNumber > 0 {
			modes[i] = modesNumber % 10
			modesNumber /= 10
			i += 1
		}
		switch opcode {
		case 1:
			params := getParam(mem[pc:], 3)
			opA := params[0]
			opB := params[1]
			dest := params[2]
			switch modes[0] {
			case 0:
				opA = mem[opA]
			case 2:
				opA = mem[relBase+opA]
			}
			switch modes[1] {
			case 0:
				opB = mem[opB]
			case 2:
				opB = mem[relBase+opB]
			}
			switch modes[2] {
			case 2:
				dest = relBase + dest
			}
			mem[dest] = opA + opB
			pc += 4
		case 2:
			params := getParam(mem[pc:], 3)
			opA := params[0]
			opB := params[1]
			dest := params[2]
			switch modes[0] {
			case 0:
				opA = mem[opA]
			case 2:
				opA = mem[relBase+opA]
			}
			switch modes[1] {
			case 0:
				opB = mem[opB]
			case 2:
				opB = mem[relBase+opB]
			}
			switch modes[2] {
			case 2:
				dest = relBase + dest
			}
			mem[dest] = opA * opB
			pc += 4
		case 3:
			params := getParam(mem[pc:], 1)
			dest := params[0]
			switch modes[0] {
			case 2:
				dest = relBase + dest
			}
			mem[dest] = <-input
			pc += 2
		case 4:
			params := getParam(mem[pc:], 1)
			out := params[0]
			switch modes[0] {
			case 0:
				out = mem[out]
			case 2:
				out = mem[relBase+out]
			}
			output <- out
			pc += 2
		case 5:
			params := getParam(mem[pc:], 2)
			opA := params[0]
			opB := params[1]
			switch modes[0] {
			case 0:
				opA = mem[opA]
			case 2:
				opA = mem[relBase+opA]
			}
			switch modes[1] {
			case 0:
				opB = mem[opB]
			case 2:
				opB = mem[relBase+opB]
			}
			if opA != 0 {
				pc += opB - pc
			} else {
				pc += 3
			}
		case 6:
			params := getParam(mem[pc:], 2)
			opA := params[0]
			opB := params[1]
			switch modes[0] {
			case 0:
				opA = mem[opA]
			case 2:
				opA = mem[relBase+opA]
			}
			switch modes[1] {
			case 0:
				opB = mem[opB]
			case 2:
				opB = mem[relBase+opB]
			}
			if opA == 0 {
				pc += opB - pc
			} else {
				pc += 3
			}
		case 7:
			params := getParam(mem[pc:], 3)
			opA := params[0]
			opB := params[1]
			dest := params[2]
			switch modes[0] {
			case 0:
				opA = mem[opA]
			case 2:
				opA = mem[relBase+opA]
			}
			switch modes[1] {
			case 0:
				opB = mem[opB]
			case 2:
				opB = mem[relBase+opB]
			}
			switch modes[2] {
			case 2:
				dest = relBase + dest
			}
			if opA < opB {
				mem[dest] = 1
			} else {
				mem[dest] = 0
			}
			pc += 4
		case 8:
			params := getParam(mem[pc:], 3)
			opA := params[0]
			opB := params[1]
			dest := params[2]
			switch modes[0] {
			case 0:
				opA = mem[opA]
			case 2:
				opA = mem[relBase+opA]
			}
			switch modes[1] {
			case 0:
				opB = mem[opB]
			case 2:
				opB = mem[relBase+opB]
			}
			switch modes[2] {
			case 2:
				dest = relBase + dest
			}
			if opA == opB {
				mem[dest] = 1
			} else {
				mem[dest] = 0
			}
			pc += 4
		case 9:
			params := getParam(mem[pc:], 1)
			opA := params[0]
			switch modes[0] {
			case 0:
				opA = mem[opA]
			case 2:
				opA = mem[relBase+opA]
			}
			relBase += opA
			pc += 2
		case 99:
			halt <- true
		default:
			halt <- true
		}
	}
}

func getParam(program []int64, param int64) []int64 {
	params := make([]int64, param)
	for i := int64(0); i < param; i++ {
		params[i] = program[i+1]
	}
	return params
}

type Point struct {
	x int
	y int
	z int
}

func (p *Point) Move(n [3]int) {
	p.x += n[0]
	p.y += n[1]
	p.z += n[2]
}

func (p Point) Angle(t Point) (angle float64) {
	angle = math.Atan2(float64(t.x-p.x), float64(t.y-p.y)) * 180 / math.Pi
	if angle < 0 {
		angle += 360
	}
	return
}

func (p Point) Distance(t Point) int {
	return Abs(p.x-t.x) + Abs(p.y-t.y)
}

func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func GCD(x, y int64) int64 {
	for y != 0 {

		x, y = y, x%y
	}
	return x
}

func RenderGrid(grid map[Point]int64) [][]string {
	minX, minY := 9999999, 99999999
	maxX, maxY := -9999999, -9999999
	for k, _ := range grid {
		if k.y < minY {
			minY = k.y
		}
		if k.x < minX {
			minX = k.x
		}
		if maxY < k.y {
			maxY = k.y
		}
		if maxX < k.x {
			maxX = k.x
		}
	}
	height := maxY - minY + 1
	width := maxX - minX + 1
	image := make([][]string, width)
	for k, v := range grid {
		x := k.x - minX
		y := k.y - minY
		if image[x] == nil {
			image[x] = make([]string, height)
		}
		switch v {
		case 0:
			image[x][y] = " "
		case 1:
			image[x][y] = "#"
		}
	}
	return image
}

func PrintImage(image [][]string) {
	for _, y := range image {
		for _, x := range y {
			fmt.Printf("%s", x)
		}
		fmt.Printf("\n")
	}
}

func RunRobot(grid map[Point]int64, start Point, input <-chan int64, output chan<- int64) {
	direction := 'U'
	location := start

	for {
		// Generate output
		if _, ok := grid[location]; !ok {
			grid[location] = 0
		}
		output <- grid[location]

		// Read input
		instruction := make([]int64, 2)
		for i := range instruction {
			instruction[i] = <-input
		}
		switch instruction[0] {
		case 0:
			grid[location] = 0
		case 1:
			grid[location] = 1
		}
		switch instruction[1] {
		case 0:
			switch direction {
			case 'U':
				direction = 'L'
			case 'L':
				direction = 'D'
			case 'D':
				direction = 'R'
			case 'R':
				direction = 'U'
			}
		case 1:
			switch direction {
			case 'U':
				direction = 'R'
			case 'L':
				direction = 'U'
			case 'D':
				direction = 'L'
			case 'R':
				direction = 'D'
			}
		}

		// Evaluate movement
		switch direction {
		case 'U':
			location.y++
		case 'D':
			location.y--
		case 'L':
			location.x--
		case 'R':
			location.x++
		}
	}
}
