package main

import (
	"os"
	"strconv"
	"strings"
)

func score(grid [][]int, x int, y int, distinct bool) int {
	paths := [][]int{{x, y}}
	goals := [][]int{}
	for len(paths) > 0 {
		x, y := paths[0][0], paths[0][1]
		current := grid[x][y]
		paths = paths[1:]

		if current == 9 {
			if distinct {
				contains := false
				for _, goal := range goals {
					if goal[0] == x && goal[1] == y {
						contains = true
						break
					}
				}
				if !contains {
					goals = append(goals, []int{x, y})
				}
			} else {
				goals = append(goals, []int{x, y})
			}
		}

		if x > 0 {
			left := grid[x-1][y]
			if left-current == 1 {
				paths = append(paths, []int{x - 1, y})
			}
		}

		if x < len(grid)-1 {
			right := grid[x+1][y]
			if right-current == 1 {
				paths = append(paths, []int{x + 1, y})
			}
		}

		if y > 0 {
			up := grid[x][y-1]
			if up-current == 1 {
				paths = append(paths, []int{x, y - 1})
			}
		}

		if y < len(grid[x])-1 {
			down := grid[x][y+1]
			if down-current == 1 {
				paths = append(paths, []int{x, y + 1})
			}
		}
	}

	return len(goals)
}

func main() {
	input, _ := os.ReadFile(os.Args[1])
	rows := strings.Split(string(input), "\n")
	rows = rows[:len(rows)-1]

	grid := make([][]int, len(rows))
	for i, line := range rows {
		grid[i] = make([]int, len(line))
		for j, cell := range strings.Split(line, "") {
			grid[i][j], _ = strconv.Atoi(cell)
		}
	}

	part_1 := 0
	part_2 := 0

	for i, row := range grid {
		for j, cell := range row {
			if cell == 0 {
				part_1 += score(grid, i, j, true)
				part_2 += score(grid, i, j, false)
			}
		}
	}

	println("Part 1:", part_1)
	println("Part 2:", part_2)
}
