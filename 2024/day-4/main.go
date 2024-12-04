package main

import (
	"os"
	"strings"
)


func count_horizontal(grid [][]string) int {
	count := 0
	
	for _, row := range grid {
		for j, cell := range row {
			if j + 3 < len(row) {
				if cell == "X" && row[j+1] == "M" && row[j+2] == "A" && row[j+3] == "S" {
					count++
				}
			}
		}
	}

	return count
}


func count_vertical(grid [][]string) int {
	transpose := make([][]string, len(grid))
	
	for i, row := range grid {
		transpose[i] = make([]string, len(row))
		for j := 0; j < len(row); j++ {
			transpose[i][j] = grid[j][i]
		}
	}
	
	return count_horizontal(transpose)
}


func count_diagonal() {

}


func main() {
	data, _ := os.ReadFile("input.txt")

	part_1 := 0

	grid := make([][]string, 0)

	for _, line := range strings.Split(string(data), "\n") {
		grid = append(grid, strings.Split(line, ""))
	}

	part_1 += count_horizontal(grid)
	part_1 += count_vertical(grid)

	println(part_1)
}
