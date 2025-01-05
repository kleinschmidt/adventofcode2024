package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := io.ReadAll(os.Stdin)
	// part1(input)
	part2(input)
}

func part1(input []byte) {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	totalSafe := 0
	for scanner.Scan() {
		line := parseLine(scanner.Text())
		if safe(line, 0) {
			totalSafe += 1
		}
	}
	fmt.Printf("Total of %d safe\n", totalSafe)
}

func part2(input []byte) {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	totalSafe := 0
	for scanner.Scan() {
		line := parseLine(scanner.Text())
		if safe(line, 1) || safe(line[1:], 0) {
			totalSafe += 1
		}
	}
	fmt.Printf("Total of %d safe\n", totalSafe)
}

func safe(report []int, skips int) bool {
	var diff int
	prev := report[0]
	pos := diffPos(report)
	for i, cur := range report[1:] {
		diff = cur - prev
		if diff == 0 || (diff >= 0) != pos || diff > 3 || diff < -3 {
			skips -= 1
			if skips < 0 {
				fmt.Println("Not safe: ", report, "(i:", i+1, ", cur:", cur, ")")
				return false
			} else {
				fmt.Println("Skipped: ", report, "(i:", i+1, ", cur:", cur, ")")
			}
		} else {
			prev = cur
		}
	}
	fmt.Println("Safe: ", report)
	return true
}

func parseLine(line string) []int {
	fields := strings.Fields(line)
	nums := make([]int, len(fields))
	for i, field := range fields {
		num, _ := strconv.Atoi(field)
		nums[i] = num
	}
	return nums
}

func diffPos(report []int) bool {
	pos := 0
	for i, x := range report[1:] {
		if x > report[i] {
			pos += 1
		} else {
			pos -= 1
		}
	}
	return pos > 0
}
