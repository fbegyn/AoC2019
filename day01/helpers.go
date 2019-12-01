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
