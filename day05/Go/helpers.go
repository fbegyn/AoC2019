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
