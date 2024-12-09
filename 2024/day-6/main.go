package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func reset(grid [][]string, rows []string) [][]string {
	for i, line := range rows {
		columns := strings.Split(line, "")
		for j, cell := range columns {
			grid[i][j] = cell
		}
	}
	return grid
}

func position(grid [][]string) (int, int) {
	for i, row := range grid {
		for j, cell := range row {
			if cell == "^" {
				return i, j
			}
		}
	}
	return -1, -1
}

func walk(grid [][]string) int {
	positions := 1
	up := make([]string, 0)
	right := make([]string, 0)
	down := make([]string, 0)
	left := make([]string, 0)

	x, y := position(grid)

	for {
		next_x := x
		next_y := y
		next_orientation := grid[x][y]

		switch grid[x][y] {
		case "^":
			if x-1 == -1 {
				return positions
			}

			if grid[x-1][y] == "#" {
				next_orientation = ">"
				if slices.Contains(up, fmt.Sprintf("%d,%d", x-1, y)) {
					return 0
				}
				up = append(up, fmt.Sprintf("%d,%d", x-1, y))
			} else {
				next_x = x - 1
				next_orientation = "^"
			}
		case ">":
			if y+1 == len(grid[x]) {
				return positions
			}

			if grid[x][y+1] == "#" {
				next_orientation = "v"
				if slices.Contains(right, fmt.Sprintf("%d,%d", x, y+1)) {
					return 0
				}
				right = append(right, fmt.Sprintf("%d,%d", x, y+1))
			} else {
				next_y = y + 1
				next_orientation = ">"
			}
		case "v":
			if x+1 == len(grid) {
				return positions
			}

			if grid[x+1][y] == "#" {
				next_orientation = "<"
				if slices.Contains(down, fmt.Sprintf("%d,%d", x+1, y)) {
					return 0
				}
				down = append(down, fmt.Sprintf("%d,%d", x+1, y))
			} else {
				next_x = x + 1
				next_orientation = "v"
			}
		case "<":
			if y-1 == -1 {
				return positions
			}

			if grid[x][y-1] == "#" {
				next_orientation = "^"
				if slices.Contains(left, fmt.Sprintf("%d,%d", x, y-1)) {
					return 0
				}
				left = append(left, fmt.Sprintf("%d,%d", x, y-1))
			} else {
				next_y = y - 1
				next_orientation = "<"
			}
		}

		if grid[next_x][next_y] == "." {
			positions += 1
		}
		grid[x][y] = "X"
		grid[next_x][next_y] = next_orientation
		x, y = next_x, next_y
	}
}

func main() {
	input, _ := os.ReadFile("input.txt")
	rows := strings.Split(string(input), "\n")
	rows = rows[:len(rows)-1]
	grid := make([][]string, len(rows))
	for i, line := range rows {
		columns := strings.Split(line, "")
		grid[i] = make([]string, len(columns))
		for j, cell := range columns {
			grid[i][j] = cell
		}
	}

	println("Part 1:", walk(grid))

	part_2 := 0

	grid = reset(grid, rows)
	for i := range grid {
		row := grid[i]
		for j := range row {
			cell := grid[i][j]
			if cell == "." {
				grid[i][j] = "#"
				if walk(grid) == 0 {
					part_2 += 1
				}
				grid = reset(grid, rows)
			}
		}
	}

	println("Part 2:", part_2)
}
