package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
)

func main() {
	level := part1("./input.txt")
	level2 := part2("./input.txt")
	fmt.Printf("Total fuel level: %d\n", level)
	fmt.Printf("Total fuel level (with fuel): %d\n", int(level2))

}

func part1(file string) (level int) {
	fi := OpenFile(file)
	defer fi.Close()

	var fuelLevels []int

	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		line := scanner.Text()
		rocketMass, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Failed to parse line into number: %v", err)
		}
		fuelNeeded := (rocketMass / 3) - 2
		fuelLevels = append(fuelLevels, fuelNeeded)
	}

	return SumOfIntArray(fuelLevels)
}

func part2(file string) (level int) {
	fi := OpenFile(file)
	defer fi.Close()

	var fuelLevels []int

	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		line := scanner.Text()
		mass, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Failed to parse line into number: %v", err)
		}
		fuelNeeded := calcFuel(mass)
		fuelLevels = append(fuelLevels, fuelNeeded)
	}

	return SumOfIntArray(fuelLevels)
}

func calcFuel(mass int) (f int) {
	fuelNeeded := (mass / 3) - 2
	if fuelNeeded <= 0 {
		return 0
	}
	return fuelNeeded + calcFuel(fuelNeeded)
}
