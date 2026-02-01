package main

import (
	"os"
	"strings"
)

func construct(garden [][]string, plots [][]int) {

}

func main() {
	input, _ := os.ReadFile(os.Args[1])

	garden := make([][]string, 0)
	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			break
		}

		garden = append(garden, strings.Split(line, ""))
	}

	for _, row := range garden {
		for _, cell := range row {
			print(cell)
		}
		println()
	}
}
