package main

import (
	"fmt"
)

func makePassword(n int) []int {
	password := make([]int, 6)
	for i := 0; i < 6; i++ {
		password[5-i] = int(n % 10)
		n /= 10
	}
	return password
}

func checkPassword(password []int) (part1, part2 bool) {
	countNumbers := map[int]int{}
	for _, number := range password {
		countNumbers[number]++
	}
	for _, frequency := range countNumbers {
		if frequency == 2 {
			part1 = true
		}
		if frequency >= 2 {
			part2 = true
		}
	}
	return
}

func main() {
	var (
		solution1, solution2 int
	)

	input := []int{193651, 649729}
	for i := input[0]; i <= input[1]; {
		if inc, password, newNum := isIncreasing(i); inc {
			part1, part2 := checkPassword(password)
			if part1 {
				solution1++
			}
			if part2 {
				solution2++
			}
			i++
		} else {
			i = int(newNum)

		}
	}
	fmt.Printf("Part1:\t%v\nPart2:\t%v", solution1, solution2)
}

func isIncreasing(n int) (bool, []int, int) {
	var pos int
	increasing := true
	password := makePassword(n)
	for i := 0; i < 5; i++ {
		if password[i] > password[i+1] {
			increasing = false
			pos = i + 1
			break
		}
	}
	//create next valid number
	newNumber := make([]int, 6)
	if !increasing {
		for i := 0; i <= pos; i++ {
			newNumber[i] = password[i]
		}
		newNumber[pos] = newNumber[pos-1]
	}
	newNum := newNumber[0]
	for i := 1; i < 6; i++ {
		newNum = newNum*10 + newNumber[i]
	}
	return increasing, password, newNum
}
