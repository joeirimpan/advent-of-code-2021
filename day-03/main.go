package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// var input []string = []string{
// 	"00100",
// 	"11110",
// 	"10110",
// 	"10111",
// 	"10101",
// 	"01111",
// 	"00111",
// 	"11100",
// 	"10000",
// 	"11001",
// 	"00010",
// 	"01010",
// }

func readFile(path string) []string {
	f, err := os.Open(path)
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

func sum(input []string) []int {
	sum := make([]int, len(input[0]))

	for _, bin := range input {
		for i, num := range bin {
			if num == '0' {
				sum[i] += 1
			}
		}
	}
	return sum
}

func part1(input []string) int64 {
	n := len(input)

	gamma := ""
	eplison := ""
	for _, s := range sum(input) {
		if s < n/2 {
			gamma += "1"
			eplison += "0"
		} else {
			gamma += "0"
			eplison += "1"
		}
	}

	g, _ := strconv.ParseInt(gamma, 2, 64)
	// complement works!
	e, _ := strconv.ParseInt(eplison, 2, 64)
	return g * e
}

func filterWith(input []string, bitPos int, having byte) []string {
	var out []string
	for _, bin := range input {
		if bin[bitPos] == having {
			out = append(out, bin)
		}
	}
	return out
}

func part2(input []string) int64 {
	return o2rating(input) * co2rating(input)
}

func co2rating(input []string) int64 {
	return filter(input, "CO2", 0)
}

func o2rating(input []string) int64 {
	return filter(input, "O2", 0)
}

func filter(input []string, rating string, pass int) int64 {
	n := len(input)

	// tail
	if n == 1 {
		d, _ := strconv.ParseInt(input[0], 2, 64)
		return d
	}

	colSum := sum(input)[pass]

	var data []string
	switch {
	case colSum < n/2, colSum == n/2:
		if rating == "O2" {
			data = filterWith(input, pass, '1')
		} else if rating == "CO2" {
			data = filterWith(input, pass, '0')
		}

	case colSum > n/2:
		if rating == "O2" {
			data = filterWith(input, pass, '0')
		} else if rating == "CO2" {
			data = filterWith(input, pass, '1')
		}
	}

	return filter(data, rating, pass+1)
}

func main() {
	input := readFile("./input.txt")

	fmt.Printf("val -> %d\n", part1(input))
	fmt.Printf("val -> %d\n", part2(input))
}
