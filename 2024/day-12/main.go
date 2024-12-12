package main

import (
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile(os.Args[1])

	garden := make([][]string, 0)

	for i, line := range strings.Split(string(input), "\n") {
		if line != "" {
			garden = append(garden, make([]string, 0))
			for _, character := range strings.Split(line, "") {
				garden[i] = append(garden[i], character)
			}
		}
	}

}
