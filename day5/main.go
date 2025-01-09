package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	// "slices"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := io.ReadAll(os.Stdin)
	parseInput(bytes)
	order, revOrder, manuals := parseInput(bytes)
	fmt.Println(order)
	fmt.Println(revOrder)
	part1(order, revOrder, manuals)
}

func parseInput(input []byte) (order map[int][]int, revOrder map[int][]int, manuals [][]int) {
	order = make(map[int][]int)
	revOrder = make(map[int][]int)
	manuals = make([][]int, 0)

	scanner := bufio.NewScanner(bytes.NewReader(input))
	// parse the X|Y orderings
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		fields := strings.FieldsFunc(line, func (r rune) bool {return r == '|'})
		first := _int(fields[0])
		second := _int(fields[1])
		order[first] = append(order[first], second)
		revOrder[second] = append(revOrder[second], first)
	}

	// parse the X,Y,Z,... manuals
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.FieldsFunc(line, func (r rune) bool {return r == ','})
		manual := make([]int, len(fields))
		for i, f := range(fields) {
			manual[i] = _int(f)
		}
		manuals = append(manuals, manual)
	}

	return
}

func _int(x string) int {
	xx, _ := strconv.Atoi(x)
	return xx
}

func part1(order map[int][]int, revOrder map[int][]int, manuals [][]int) (total int) {
	total = 0
	for _, manual := range manuals {
		// strategy here is:
		// three different lookup indices:
		// - for page number, all pages that must be later
		// - for page number, all pages that must be earlier
		// - for page in a manual, the order (index).
		//
		// then we iterate through the manual, and check that the index is
		// less/greater where required
		index := make(map[int]int)
		for i, page := range manual {
			index[page] = i
		}

		okay := true

		fmt.Println("Checking manual ", manual)
		for i, page := range manual {
			nexts := order[page]
			prevs := revOrder[page]
			for _, next := range nexts {
				if !okay {break}
				idx, ok := index[next]
				if ok && idx < i {
					fmt.Printf("Violation: %d (%d) precedes %d (%d)\n", next, idx, page, i)
					okay = false
				}
			}

			if !okay {break}
			for _, prev := range prevs {
				if !okay {break}
				idx, ok := index[prev]
				if ok && idx > i {
					fmt.Printf("Violation: %d (%d) follows %d (%d)\n", prev, idx, page, i)
					okay = false
				}
			}					
		}

		if okay {
			total += manual[len(manual) / 2]
		}
	}
	fmt.Println("Total of middle pages: ", total)
	return total
}

func part2() {}
