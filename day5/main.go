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
	bytes, _ := io.ReadAll(os.Stdin)
	parseInput(bytes)
	order, _ := parseInput(bytes)
	fmt.Println(order)
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
		fields := strings.FieldsFunc(line, func (r rune) bool {return r == '|'})
		first := _int(fields[0])
		second := _int(fields[1])
		order[first] = append(order[first], second)
		fmt.Println(line, " -> ", fields)
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
		fmt.Println(line, " -> ", manual)		
	}

	return
}

func _int(x string) int {
	xx, _ := strconv.Atoi(x)
	return xx
}

func part1(order map[int][]int, manuals [][]int) {
	for _, manual := range manuals {
		index := make(map[int]int)
		for i, page := range manual {
			index[page] = i
		}

		for i, page := range manual {
			orders := order[page]
			
		}
	}
}

func part2() {}
