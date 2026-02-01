package main

import (
	"os"
	"slices"
	"strconv"
	"strings"
)

func generate_memory(files string) []string {
	memory := make([]string, 0)

	for i, file := range strings.Split(files, "") {
		length, _ := strconv.Atoi(file)
		if i % 2 == 0 {
			for j := 0; j < length; j++ {
				memory = append(memory, strconv.Itoa(i / 2))
			}
		} else {
			for j := 0; j < length; j++ {
				memory = append(memory, ".")
			}
		}
	}

	return memory
}

func move_memory_1(memory []string) []string {
	i := len(memory) - 1
	for {
		next_space := next_space(memory, i)
		if next_space != -1 {
			memory[next_space] = memory[i]
			memory[i] = "."
		} else {
			break
		}
		i--
	}

	return memory
}

func next_space(memory []string, i int) int {
	for j := 0; j < i; j++ {
		if memory[j] == "." {
			return j 
		}
	}
	return -1 
}

func move_memory_2(memory []string) []string {
	i := len(memory) - 1
	processed := make([]string, 0)
	
	for {
		if i == -1 {
			break
		}

		if memory[i] != "." {
			id := memory[i]
			block := id
			spaces := []int{i}

			if !slices.Contains(processed, id) {	
				for {
					if i == 0 || memory[i - 1] != id {
						break
					}

					i--
					block += id 
					spaces = append(spaces, i)
				}

				next_spaces := next_spaces(memory, i, block, id)
				
				if len(next_spaces) == (len(block) / len(id)) {
					for _, space := range next_spaces {
						memory[space] = id 
					}
					for _, space := range spaces {
						memory[space] = "."
					}
				}

				processed = append(processed, id)
			}

		}
		i--
	}

	return memory
}

func next_spaces(memory []string, i int, block string, id string) []int {
	next_spaces := make([]int, 0)
	
	for j := 0; j < i; j++ {
		if memory[j] == "." {
			next_spaces = append(next_spaces, j)
		} else {
			next_spaces = []int{}
		}

		if len(next_spaces) == (len(block) / len(id)) {
			break
		}

	}

	return next_spaces

}

func checksum_memory(memory []string) int {
	checksum := 0
	
	for i, file := range memory {
		if file == "." {
			continue	
		}
		id, _ := strconv.Atoi(file)
		checksum += i * id
	}
	
	return checksum
}

func main() {
	input, _ := os.ReadFile(os.Args[1])
	files := strings.TrimSpace(string(input))
	memory := generate_memory(files)
	memory = move_memory_1(memory)
	println("Part 1:", checksum_memory(memory))
	memory = generate_memory(files)
	memory = move_memory_2(memory)
	println("Part 2:", checksum_memory(memory))
}
