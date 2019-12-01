package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	src, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("failed to read file into scanner: %v", err)
	}
	defer src.Close()

	var fuelLevels []int
	var fuelLevels2 []float64

	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		line := scanner.Text()
		rocketMass, err := strconv.ParseFloat(line, 64)
		if err != nil {
			log.Fatalf("Failed to parse line into numbber: %v", err)
		}
		fuelNeeded := int(math.Floor(rocketMass/3) - 2)
		fuelLevels = append(fuelLevels, fuelNeeded)
		fuelLevels2 = append(fuelLevels2, calcFuel(rocketMass))

	}

	level := 0
	for _, v := range fuelLevels {
		level += v
	}

	level2 := float64(0)
	for _, v := range fuelLevels2 {
		level2 += v
	}

	fmt.Printf("Total fuel level: %d\n", level)
	fmt.Printf("Total fuel level (with fuel): %d\n", int(level2))
}

func calcFuel(mass float64) (f float64) {
	fuelNeeded := math.Floor(mass/3) - 2
	if fuelNeeded <= 0 {
		return 0
	} else {
		return fuelNeeded + calcFuel(fuelNeeded)
	}
}
