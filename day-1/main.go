package main


import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func calculate_distance(a int, b int) int {
	var distance = a - b
	if distance < 0 {
		distance *= -1
	}
	return distance 
}


func count(list []int, e int) int {
	var count = 0
	for _, number := range(list) {
		if number == e {
			count += 1	
		}
	}
	return count
}


func main() {
	var left [1000]int
	var right [1000]int
	var distance = 0
	var score = 0

	data, _ := os.ReadFile("input.txt")
	
	for i, line := range(strings.Split(string(data), "\n")) {
		if i < 1000 {
			var numbers = strings.Split(line, "   ")
			
			var number_1, _ = strconv.Atoi(numbers[0])
			var number_2, _ = strconv.Atoi(numbers[1])

			left[i] = number_1; 
			right[i] = number_2;
		}
	}

	sort.Ints(left[:])
	sort.Ints(right[:])

	for i := 0; i < 1000; i++ {
		distance += calculate_distance(left[i], right[i])
		score += left[i] * count(right[:], left[i])
		
	}

	fmt.Println("Part 1:", distance)
	fmt.Println("Part 2:", score)
}
