package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

//var depths []int = []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

func readFile(path string) []int {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	defer f.Close()

	var depths []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("error converting int: %v", err)
		}

		depths = append(depths, i)
	}

	return depths
}

func part1(depths []int) int {
	increased := 0
	for i := 0; i < len(depths); i++ {
		// Skip first
		if i == 0 {
			continue
		}

		if depths[i] > depths[i-1] {
			increased++
		}
	}

	return increased
}

func part2(depths []int) int {
	increased := 0
	for i := 1; i < len(depths)-2; i++ {
		a, d := depths[i-1], depths[i+2]

		if d > a {
			increased++
		}
	}

	return increased
}

func main() {
	depths := readFile("./input.txt")

	fmt.Printf("part1 -> %d\n", part1(depths))
	fmt.Printf("part1 -> %d\n", part2(depths))
}
