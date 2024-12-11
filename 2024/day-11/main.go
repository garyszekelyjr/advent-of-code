package main

import (
	"os"
	"strconv"
	"strings"
)

func process(stones []string, blinks int) int {
	return 0
}

func blink(stones []string) []string {
	temp := make([]string, 0)
	for _, stone := range stones {
		if stone == "0" {
			temp = append(temp, "1")
		} else if len(stone)%2 == 0 {
			number_1, _ := strconv.Atoi(stone[len(stone)/2:])
			number_2, _ := strconv.Atoi(stone[:len(stone)/2])

			stone_1 := strconv.Itoa(number_1)
			stone_2 := strconv.Itoa(number_2)

			temp = append(temp, stone_2, stone_1)
		} else {
			number, _ := strconv.Atoi(stone)
			temp = append(temp, strconv.Itoa(number*2024))
		}
	}
	return temp
}

func main() {
	input, _ := os.ReadFile(os.Args[1])
	blinks, _ := strconv.Atoi(os.Args[2])
	stones := strings.Split(strings.TrimSpace(string(input)), " ")

	for i := 0; i < blinks; i++ {
		stones = blink(stones)
	}

	println("Part 1:", len(stones))
}
