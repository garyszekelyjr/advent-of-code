package main


import (
	"os"
	"strconv"
	"strings"
)


func abs(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}


func is_safe(levels []int) bool {
	increasing := levels[0] < levels[1]
	for i, curr := range(levels) {
		if i == len(levels) - 1 {
			break
		}

		next := levels[i + 1]
		difference := abs(curr - next)

		if difference < 1 || difference > 3 {
			return false
		}


		if increasing {
			if curr > next {
				return false
			}
		} else {
			if curr < next {
				return false
			}
		}
	}
	return true
}


func main() {
	data, _ := os.ReadFile("input.txt")

	part_1 := 0
	part_2 := 0

	for _, line := range(strings.Split(string(data), "\n")) {
		if line != "" {
			elements := strings.Split(line, " ")
			levels := make([]int, len(elements))
			for i, element := range(elements) {
				levels[i], _ = strconv.Atoi(element)
			}
			if is_safe(levels) {
				part_1 += 1
			} else {
				for i := 0; i < len(levels); i++ {
					sublevels := make([]int, len(levels))
					copy(sublevels, levels)
					if is_safe(append(sublevels[:i], sublevels[i + 1:]...)) {
						part_2 += 1
						break
					}
				}
			}
		}
	}

	println("Part 1:", part_1)
	println("Part 2:", part_1 + part_2)
}

