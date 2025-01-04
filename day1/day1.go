package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, _ := io.ReadAll(os.Stdin)

	part1(input)


	// part 2
	// scanner := bufio.NewScanner(bytes.NewReader(input))

	part2(input)
	
}

func part1(input []byte) {
	first := make([]int, 0)
	second := make([]int, 0)

	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		one, _ := strconv.Atoi(fields[0])
		two, _ := strconv.Atoi(fields[1])
		first = append(first, one)
		second = append(second, two)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	slices.Sort(first)
	slices.Sort(second)

	total := 0
	for i, one := range first {
		two := second[i]
		// total += two - one
		if one < two {
			total += two - one
		} else {
			total += one - two
		}
	}

	fmt.Printf("Part 1: Total distance: %d\n", total)
}

func part2(input []byte) {
	first := make([]int, 0)
	counts := make(map[int]int)
	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		one, _ := strconv.Atoi(fields[0])
		two, _ := strconv.Atoi(fields[1])
		first = append(first, one)
		// retrieval of non-existent key gives zero value
		counts[two] += 1
	}

	total := 0
	for _, one := range first {
		total += counts[one] * one
	}
	fmt.Printf("Part 2: Total %d\n", total)
}
