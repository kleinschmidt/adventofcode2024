package main

import (
	"bufio"
	"bytes"
	"cmp"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := io.ReadAll(os.Stdin)
	parseInput(bytes)
	order, manuals := parseInput(bytes)
	fmt.Println(order)
	fmt.Println("\n---- Part 1: ----")
	part1(order, manuals)
	fmt.Println("\n---- Part 2: ----")
	part2(order, manuals)
}

func parseInput(input []byte) (order map[int][]int, manuals [][]int) {
	order = make(map[int][]int)
	manuals = make([][]int, 0)

	scanner := bufio.NewScanner(bytes.NewReader(input))
	// parse the X|Y orderings
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		fields := strings.FieldsFunc(line, func(r rune) bool { return r == '|' })
		first := _int(fields[0])
		second := _int(fields[1])
		order[first] = append(order[first], second)
	}

	// parse the X,Y,Z,... manuals
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.FieldsFunc(line, func(r rune) bool { return r == ',' })
		manual := make([]int, len(fields))
		for i, f := range fields {
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

type Pair struct{ a, b int }

func (p Pair) Delta() int {
	diff := p.a - p.b
	if diff < 0 {
		diff = -diff
	}
	return diff
}

func (p Pair) String() string {
	return fmt.Sprintf("(%d, %d)", p.a, p.b)
}

func violations(order map[int][]int, manual []int) (bad []Pair) {
	bad = make([]Pair, 0)
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

	fmt.Println("Checking manual ", manual)
	for i, page := range manual {
		// check forward ordering for violations (if any page that must occur _after_ page comes earlier)
		for _, next := range order[page] {
			idx, ok := index[next]
			if ok && idx < i {
				bad = append(bad, Pair{i, idx})
				fmt.Printf("Violation: %d (%d) precedes %d (%d)\n", next, idx, page, i)
			}
		}
	}

	return bad
}

func part1(order map[int][]int, manuals [][]int) (total int) {
	total = 0
	for _, manual := range manuals {
		badPairs := violations(order, manual)
		fmt.Println("Found violations:", badPairs)
		if len(badPairs) == 0 {
			total += manual[len(manual)/2]
		}
	}
	fmt.Println("Total of middle pages: ", total)
	return total
}

func fixManualMiddlePage(order map[int][]int, manual []int) (middle int) {
	badPairs := violations(order, manual)
	fmt.Println("Found violations:", badPairs)
	if len(badPairs) == 0 {
		return 0
	}
	for len(badPairs) > 0 {
		p := slices.MaxFunc(badPairs, func(p1, p2 Pair) int {
			return cmp.Compare(p1.Delta(), p2.Delta())
		})
		fmt.Printf("Worst violation: %v\n", p)
		manual[p.a], manual[p.b] = manual[p.b], manual[p.a]
		badPairs = violations(order, manual)
	}
	return manual[len(manual)/2]
}

func part2(order map[int][]int, manuals [][]int) (total int) {
	total = 0
	for _, manual := range manuals {
		total += fixManualMiddlePage(order, manual)
	}
	fmt.Println("Total of middle pages: ", total)
	return total
}
