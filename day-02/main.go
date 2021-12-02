package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// var input []string = []string{
// 	"forward 5",
// 	"down 5",
// 	"forward 8",
// 	"up 3",
// 	"down 8",
// 	"forward 2",
// }

const (
	Forward = "f"
	Down    = "d"
	Up      = "u"
)

func readFile(path string) []string {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	defer f.Close()

	var depths []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		depths = append(depths, scanner.Text())
	}

	return depths
}

func part1(input []string) int {
	depth, forward := 0, 0

	for _, i := range input {
		direction := string(i[0])
		value, _ := strconv.Atoi(string(i[len(i)-1]))

		switch direction {
		case Forward:
			forward += value
		case Down:
			depth += value
		case Up:
			depth -= value
		}
	}

	return depth * forward
}

func part2(input []string) int {
	depth, forward, aim := 0, 0, 0

	for _, i := range input {
		direction := string(i[0])
		value, _ := strconv.Atoi(string(i[len(i)-1]))

		switch direction {
		case Forward:
			forward += value
			if aim != 0 {
				depth += aim * value
			}
		case Down:
			aim += value
		case Up:
			aim -= value
		}
	}

	return depth * forward
}

func main() {
	input := readFile("./input.txt")

	fmt.Printf("val -> %d\n", part1(input))
	fmt.Printf("val -> %d\n", part2(input))
}
