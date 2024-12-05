package main

import (
	"os"
	"strconv"
	"strings"
)


func swap(numbers []string, before int, after int) []string {
	temp := numbers[before]
	numbers[before] = numbers[after]
	numbers[after] = temp
	return numbers
}


func valid(numbers []string, rules []string) (bool, int, int) {
	for _, rule := range rules {
		if rule == "" {
			continue
		}
		before := strings.Split(rule, "|")[0]
		after := strings.Split(rule, "|")[1]

		before_i := index(numbers, before)
		after_i := index(numbers, after)

		if before_i == -1 || after_i == -1 {
			continue
		}

		if before_i > after_i {
			return false, before_i, after_i
		}
	}
	return true, -1, -1

}


func index(list []string, target string) int {
	for i, v := range list {
		if v == target {
			return i
		}
	}
	return -1
}


func main() {
	input, _ := os.ReadFile("input.txt")	
	rules, _ := os.ReadFile("rules.txt")

	part_1 := 0
	part_2 := 0

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}
		numbers := strings.Split(line, ",")
		correct, before, after := valid(numbers, strings.Split(string(rules), "\n"))
		if correct{ 
			middle, _ := strconv.Atoi(numbers[(len(numbers) - 1) / 2])
			part_1 += middle
		} else {
			numbers = swap(numbers, before, after)
			for {
				correct, before, after = valid(numbers, strings.Split(string(rules), "\n"))
				if correct {
					middle, _ := strconv.Atoi(numbers[(len(numbers) - 1) / 2])
					part_2 += middle
					break
				}
				numbers = swap(numbers, before, after)
			}
		}
	}


	println("Part 1:", part_1)
	println("Part 2:", part_2)
}
