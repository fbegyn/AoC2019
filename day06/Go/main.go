package main

import (
	"bufio"
	"fmt"
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

	orbits := getOrbits(*input)

	orbTotal := part1(orbits)
	fmt.Printf("Total direct and indirect orbits: %d\n", orbTotal)
	jumps := part2(orbits)
	fmt.Printf("Total jumps: %d\n", jumps)
}

func getOrbits(file string) (orbits map[string]string) {
	fi := OpenFile(file)
	defer fi.Close()
	orbits = make(map[string]string)
	// setup input scanner
	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		line := scanner.Text()
		obs := strings.SplitN(line, ")", 2)
		orbits[obs[1]] = obs[0]
	}
	return
}

func part1(orbits map[string]string) (orbTotal int) {
	for _, orbit := range orbits {
		for {
			orbTotal++
			p, ok := orbits[orbit]
			if !ok {
				break
			}
			orbit = p
		}
	}
	return
}

func part2(orbits map[string]string) (jumps int) {
	path1, path2 := make([]string, 0), make([]string, 0)
	o, ok := "YOU", false
	for {
		o, ok = orbits[o]
		if !ok {
			break
		}
		path1 = append(path1, o)
	}

	o = "SAN"
	for {
		o, ok = orbits[o]
		if !ok {
			break
		}
		path2 = append(path2, o)
	}
	for i, j := len(path1)-1, len(path2)-1; i >= 0 && j >= 0; {
		if path1[i] == path2[j] {
			jumps++
		} else {
			break
		}
		i--
		j--
	}
	jumps = len(path1) + len(path2) - jumps*2
	return
}
