package main

import (
	"log"
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

func RunProgram(program []int, input <-chan int, output chan<- int, halt chan<- bool) {
	mem := make([]int, len(program))
	copy(mem, program)
	pc := 0

	for {
		opcode := mem[pc] % 100
		modesNumber := mem[pc] / 100
		modes := make([]int, 3)
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
			if modes[0] == 0 {
				opA = mem[opA]
			}
			if modes[1] == 0 {
				opB = mem[opB]
			}
			mem[dest] = opA + opB
			pc += 4
		case 2:
			params := getParam(mem[pc:], 3)
			opA := params[0]
			opB := params[1]
			dest := params[2]
			if modes[0] == 0 {
				opA = mem[opA]
			}
			if modes[1] == 0 {
				opB = mem[opB]
			}
			mem[dest] = opA * opB
			pc += 4
		case 3:
			params := getParam(mem[pc:], 1)
			dest := params[0]
			// fmt.Print("Entememinput: ")
			// _, err := fmt.Scanf("%d", &id)
			// if err != nil {
			// 	log.Fatalf("Could not convert string to id: %v\n", err)
			// }
			mem[dest] = <-input
			pc += 2
		case 4:
			params := getParam(mem[pc:], 1)
			out := params[0]
			out = mem[out]
			output <- out
			pc += 2
		case 5:
			params := getParam(mem[pc:], 2)
			opA := params[0]
			opB := params[1]
			if modes[0] == 0 {
				opA = mem[opA]
			}
			if modes[1] == 0 {
				opB = mem[opB]
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
			if modes[0] == 0 {
				opA = mem[opA]
			}
			if modes[1] == 0 {
				opB = mem[opB]
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
			if modes[0] == 0 {
				opA = mem[opA]
			}
			if modes[1] == 0 {
				opB = mem[opB]
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
			if modes[0] == 0 {
				opA = mem[opA]
			}
			if modes[1] == 0 {
				opB = mem[opB]
			}
			if opA == opB {
				mem[dest] = 1
			} else {
				mem[dest] = 0
			}
			pc += 4
		case 99:
			halt <- true
		default:
			halt <- true
		}
	}
}

func getParam(program []int, param int) []int {
	params := make([]int, param)
	for i := 0; i < param; i++ {
		params[i] = program[i+1]
	}
	return params
}
