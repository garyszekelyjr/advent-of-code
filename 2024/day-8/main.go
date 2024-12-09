package main

import (
	"math"
	"os"
	"strings"
)

func distance(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func line(x1 float64, y1 float64, x2 float64, y2 float64) (float64, float64) {
	if x2-x1 == 0 {
		return -1.0, -1.0
	}
	m := (y2 - y1) / (x2 - x1)
	b := y1 - (m * x1)
	return m, b
}

func antennas(grid [][]string) [][]float64 {
	antennas := make([][]float64, 0)
	for i, row1 := range grid {
		for j, cell1 := range row1 {
			if cell1 == "." {
				continue
			}

			for k, row2 := range grid {
				for l, cell2 := range row2 {
					if i == k && j == l {
						continue
					}

					if cell1 == cell2 {
						antennas = append(antennas, []float64{float64(i), float64(j), float64(k), float64(l)})
					}
				}
			}
		}
	}
	return antennas
}

func main() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")
	lines = lines[:len(lines)-1]

	grid := make([][]string, len(lines))
	for i, line := range lines {
		grid[i] = make([]string, len(line))
		for j, char := range line {
			grid[i][j] = string(char)
		}
	}

	part_1 := 0
	part_2 := 0

	antennas := antennas(grid)

	for i := 0.0; i < float64(len(grid)); i++ {
		row := grid[int(i)]
		for j := 0.0; j < float64(len(row)); j++ {
			for _, antenna := range antennas {
				if antenna[0] == i && antenna[1] == j {
					continue
				}

				m, b := line(antenna[0], antenna[1], antenna[2], antenna[3])
				m1, b1 := line(antenna[0], antenna[1], i, j)
				m2, b2 := line(antenna[2], antenna[3], i, j)

				tolerance := 0.01

				if math.Abs(m2-m) <= tolerance &&
					math.Abs(m1-m) <= tolerance &&
					math.Abs(b2-b) <= tolerance &&
					math.Abs(b1-b) <= tolerance {
					if distance(antenna[0], antenna[1], i, j)*2 == distance(antenna[2], antenna[3], i, j) {
						part_1++
						break
					}
				}
			}

			for _, antenna := range antennas {
				if antenna[0] == i && antenna[1] == j {
					part_2++
					break
				}

				m, b := line(antenna[0], antenna[1], antenna[2], antenna[3])
				m1, b1 := line(antenna[0], antenna[1], i, j)
				m2, b2 := line(antenna[2], antenna[3], i, j)

				tolerance := 0.01

				if math.Abs(m2-m) <= tolerance &&
					math.Abs(m1-m) <= tolerance &&
					math.Abs(b2-b) <= tolerance &&
					math.Abs(b1-b) <= tolerance {
					part_2++
					break
				}
			}
		}
	}

	println("Part 1:", part_1)
	println("Part 2:", part_2)
}
