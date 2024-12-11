package main

import (
	"os"
	"strconv"
	"strings"
)

func memoized_blink(stones []string, blinks int) int {
	counts := make(map[string]int)
	count := 0

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

	for _, value := range counts {
		count += value
	}

	return count
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
	blinks_1, _ := strconv.Atoi(os.Args[2])
	blinks_2, _ := strconv.Atoi(os.Args[3])

	stones := strings.Split(strings.TrimSpace(string(input)), " ")
	for i := 0; i < blinks_1; i++ {
		stones = blink(stones)
	}
	println("Part 1:", len(stones))

	stones = strings.Split(strings.TrimSpace(string(input)), " ")
	println("Part 2:", memoized_blink(stones, blinks_2))
}
