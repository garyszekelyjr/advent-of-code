package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)


func base_3(number int, length int) []int {
	ternary := make([]int, 0)
	for i := 0; i < length; i++ {
		bit := (number % 3)
		ternary = append(ternary, bit)
		number = number / 3
	}
	return ternary
}


func base_2(number int, length int) []int {
	binary := make([]int, 0)
	for i := 0; i < length; i++ {
		bit := (number % 2)
		binary = append(binary, bit)
		number = number / 2
	}
	return binary
}


func main() {
	input, _ := os.ReadFile("input.txt")
	rows := strings.Split(string(input), "\n")
	rows = rows[:len(rows)-1]

	part_1 := 0
	part_2 := 0

	for _, row := range rows {
		result, _ := strconv.Atoi(strings.Split(row, ":")[0])
		numbers := strings.Split(strings.Split(row, ": ")[1], " ")
	
		for i := 0; i < int(math.Pow(2, float64(len(numbers) - 1))); i++ {
			operators := base_2(i, len(numbers) - 1)

			output, _ := strconv.Atoi(numbers[0])
			for j := 1; j < len(numbers); j ++ {
				number, _ := strconv.Atoi(numbers[j])
				if operators[j - 1] == 0 {
					output = output + number	
				} else {
					output = output * number 
				}
			}

			if output == result {
				part_1 += result
				break
			}
		}
		

		for i := 0; i < int(math.Pow(3, float64(len(numbers) - 1))); i++ {
			operators := base_3(i, len(numbers) - 1)

			output, _ := strconv.Atoi(numbers[0])
			for j := 1; j < len(numbers); j ++ {
				number, _ := strconv.Atoi(numbers[j])
				if operators[j - 1] == 0 {
					output = output + number	
				} else if operators[j - 1] == 1 {
					output = output * number 
				} else {
					output, _ = strconv.Atoi(fmt.Sprintf("%d%d", output, number))
				}
			}

			if output == result {
				part_2 += result
				break
			}
		}

	}

	println("Part 1:", part_1)
	println("Part 2:", part_2)
}

