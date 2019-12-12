package main

import (
	"bufio"
	"fmt"
	"log"
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

	moons := getMoons(*input)
	part1(moons, *steps)
	part2(moons)
}

func part1(initMoons []Moon, steps int) {
	moons := make([]Moon, len(initMoons))
	copy(moons, initMoons)
	for i := 0; i < steps; i++ {
		for i := range moons {
			for b, t := range moons {
				if i == b {
					continue
				}
				moons[i].updateVelocity(t)
			}
		}
		for i := range moons {
			moons[i].move()
		}
	}
	energy := 0
	for _, m := range moons {
		energy += m.energy()
	}
	fmt.Printf("Total energy after %d steps: %d\n", steps, energy)
}

func part2(initMoons []Moon) {
	moons := make([]Moon, len(initMoons))
	copy(moons, initMoons)
	repeat := false
	coordEqual := make([]bool, 3)
	coordIter := make([]int64, 3)
	for !repeat {
		// time step
		for i := range moons {
			for b, t := range moons {
				if i == b {
					continue
				}
				moons[i].updateVelocity(t)
			}
		}
		for i := range moons {
			moons[i].move()
		}

		// look if a component of the Moon is equal to the initial state. If we find at which point each
		// of the components is equal to the original state we can use Lowest Common Multiplier to find
		// how many iterations will be needed to reach the initial state again.
		for c := 0; c < 3; c++ {
			if !coordEqual[c] {
				if compareMoons(initMoons, moons, c) {
					coordEqual[c] = true
				} else {
					coordIter[c]++
				}
			}
		}
		repeat = coordEqual[0] && coordEqual[1] && coordEqual[2]
	}

	for i := range coordIter {
		coordIter[i]++
	}

	iters := LCM(coordIter[0], coordIter[1], coordIter[2])
	fmt.Printf("It took %d iterations for history to repeat.\n", iters)
}

func compareMoons(src, tar []Moon, comp int) bool {
	moonsEq := make([]bool, len(src))
	equalLocation := false
	equalVel := false
	for i := range src {
		switch comp {
		case 0:
			equalLocation = src[i].location.x == tar[i].location.x
		case 1:
			equalLocation = src[i].location.y == tar[i].location.y
		case 2:
			equalLocation = src[i].location.z == tar[i].location.z
		}
		if src[i].velocity[comp] == tar[i].velocity[comp] {
			equalVel = true
		}
		moonsEq[i] = equalLocation && equalVel
	}

	equal := true
	for _, b := range moonsEq {
		equal = equal && b
	}
	return equal
}

type Moon struct {
	location Point
	velocity [3]int
}

func (m *Moon) move() {
	m.location.Move(m.velocity)
}

func (m *Moon) updateVelocity(t Moon) {
	if m.location.x < t.location.x {
		m.velocity[0]++
	} else if t.location.x < m.location.x {
		m.velocity[0]--
	}
	if m.location.y < t.location.y {
		m.velocity[1]++
	} else if t.location.y < m.location.y {
		m.velocity[1]--
	}
	if m.location.z < t.location.z {
		m.velocity[2]++
	} else if t.location.z < m.location.z {
		m.velocity[2]--
	}
}

func (m *Moon) energy() int {
	location := m.location
	potEnergy := Abs(location.x) + Abs(location.y) + Abs(location.z)
	kinEnergy := Abs(m.velocity[0]) + Abs(m.velocity[1]) + Abs(m.velocity[2])
	return potEnergy * kinEnergy
}

func getMoons(fiDsc string) []Moon {
	fi := OpenFile(fiDsc)
	defer fi.Close()

	var moons []Moon

	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		line = strings.Trim(strings.Trim(line, ">"), "<")
		var x, y, z int
		_, err := fmt.Sscanf(line, "x=%d, y=%d, z=%d", &x, &y, &z)
		if err != nil {
			log.Fatalf("Failed to parse moon location: %v\n", err)
		}
		coord := Point{x, y, z}
		moons = append(moons, Moon{coord, [3]int{0, 0, 0}})
	}
	return moons
}
