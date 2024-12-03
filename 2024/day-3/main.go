package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)


func main() {
	data, _ := os.ReadFile("input.txt")
	r, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)

	part_1 := 0
	part_2 := 0

	for _, matches := range r.FindAllStringSubmatch(string(data), -1) {
		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		part_1 += x * y
	}

	for _, do := range strings.Split(string(data), "do()") {
		do = strings.Split(do, "don't()")[0]
		for _, matches := range r.FindAllStringSubmatch(do, -1) {
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])
			part_2 += x * y
		}
	}

	println("Part 1:", part_1)
	println("Part 2:", part_2)
}

