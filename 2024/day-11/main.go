package main

import (
	"os"
	"strconv"
	"strings"
)

func blink(stones []string, blinks int) int {
	counts := make(map[string]int)

	for _, stone := range stones {
		counts[stone] = 1
	}

	for i := 0; i < blinks; i++ {
		temp := make(map[string]int)
		for stone, value := range counts {
			if stone == "0" {
				temp["1"] += value
			} else if len(stone)%2 == 0 {
				number_1, _ := strconv.Atoi(stone[len(stone)/2:])
				number_2, _ := strconv.Atoi(stone[:len(stone)/2])

				stone_1 := strconv.Itoa(number_1)
				stone_2 := strconv.Itoa(number_2)

				temp[stone_1] += value
				temp[stone_2] += value
			} else {
				number, _ := strconv.Atoi(stone)
				stone = strconv.Itoa(number * 2024)
				temp[stone] += value
			}
		}
		counts = temp
	}

	count := 0

	for _, value := range counts {
		count += value
	}

	return count
}

func main() {
	input, _ := os.ReadFile(os.Args[1])

	stones := strings.Split(strings.TrimSpace(string(input)), " ")
	println("Part 1:", blink(stones, 25))

	stones = strings.Split(strings.TrimSpace(string(input)), " ")
	println("Part 2:", blink(stones, 75))
}
