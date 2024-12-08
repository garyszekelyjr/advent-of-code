package main


import (
	"os"
	"strings"
)


func part_2_count(grid [][]string) int {
	count := 0

	for i, row := range grid {
		for j, cell := range row {
			if i + 1 < len(grid) && j + 1 < len(grid[i]) && grid[i + 1][j + 1] == "A" {
				if cell == "M" && i + 2 < len(grid) && j + 2 < len(grid[i]) && grid[i + 2][j + 2] == "S" ||
					cell == "S" && i + 2 < len(grid) && j + 2 < len(grid[i]) && grid[i + 2][j + 2] == "M" {
					if j + 2 < len(grid[i]) && grid[i][j + 2] == "S" &&
						i + 2 < len(grid) && grid[i + 2][j] == "M" {
						count++
					}

					if j + 2 < len(grid[i]) && grid[i][j + 2] == "M" && 
						i + 2 < len(grid) && grid[i + 2][j] == "S" {
						count++
					}
				}
			}
		}
	}
	
	return count
}


func part_1_count(grid [][]string, x int, y int) int {
	count := 0

	if y + 1 < len(grid[x]) && grid[x][y + 1] == "M" {
		if y + 2 < len(grid[x]) && grid[x][y + 2] == "A" {
			if y + 3 < len(grid[x]) && grid[x][y + 3] == "S" {
				count++
			}
		}
	}

	if y - 1 >= 0 && grid[x][y - 1] == "M" {
		if y - 2 >= 0 && grid[x][y - 2] == "A" {
			if y - 3 >= 0 && grid[x][y - 3] == "S" {
				count++
			}
		
		}
	}

	if x + 1 < len(grid) && grid[x + 1][y] == "M" {
		if x + 2 < len(grid) && grid[x + 2][y] == "A" {
			if x + 3 < len(grid) && grid[x + 3][y] == "S" {
				count++
			}
		}
	}

	if x - 1 >= 0 && grid[x - 1][y] == "M" {
		if x - 2 >= 0 && grid[x - 2][y] == "A" {
			if x - 3 >= 0 && grid[x - 3][y] == "S" {
				count++
			}
		}
	}

	if x + 1 < len(grid) && y + 1 < len(grid[x]) && grid[x + 1][y + 1] == "M" {
		if x + 2 < len(grid) && y + 2 < len(grid[x]) && grid[x + 2][y + 2] == "A" {
			if x + 3 < len(grid) && y + 3 < len(grid[x]) && grid[x + 3][y + 3] == "S" {
				count++
			}
		}
	}

	if x - 1 >= 0 && y - 1 >= 0 && grid[x - 1][y - 1] == "M" {
		if x - 2 >= 0 && y - 2 >= 0 && grid[x - 2][y - 2] == "A" {
			if x - 3 >= 0 && y - 3 >= 0 && grid[x - 3][y - 3] == "S" {
				count++
			}
		}
	}

	if x + 1 < len(grid) && y - 1 >= 0 && grid[x + 1][y - 1] == "M" {
		if x + 2 < len(grid) && y - 2 >= 0 && grid[x + 2][y - 2] == "A" {
			if x + 3 < len(grid) && y - 3 >= 0 && grid[x + 3][y - 3] == "S" {
				count++
			}
		}
	}

	if x - 1 >= 0 && y + 1 < len(grid[x]) && grid[x - 1][y + 1] == "M" {
		if x - 2 >= 0 && y + 2 < len(grid[x]) && grid[x - 2][y + 2] == "A" {
			if x - 3 >= 0 && y + 3 < len(grid[x]) && grid[x - 3][y + 3] == "S" {
				count++
			}
		}
	}

	return count
}


func main() {
	data, _ := os.ReadFile("input.txt")

	part_1 := 0

	grid := make([][]string, 0)

	for _, line := range strings.Split(string(data), "\n") {
		if line != "" {
			grid = append(grid, strings.Split(line, ""))
		}
	}

	for i, row := range grid {
		for j, cell := range row {
			if cell == "X" {
				part_1 += part_1_count(grid, i, j)
			}
		}
	} 

	println("Part 1:", part_1)
	println("Part 2:", part_2_count(grid))
}

