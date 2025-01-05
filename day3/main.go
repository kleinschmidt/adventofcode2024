package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	buf := new(strings.Builder)
	io.Copy(buf, os.Stdin)
	input := buf.String()
	part1(input)
	part2(input)
}

func part1(input string) {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	total := 0
	for _, match := range matches {
		first, _ := strconv.Atoi(match[1])
		second, _ := strconv.Atoi(match[2])
		total += first * second
	}
	fmt.Println("Total: ", total)
}

func part2(input string) {
	re := regexp.MustCompile(`(mul|do|don't)\((\d+)?,?(\d+)?\)`)
	matches := re.FindAllStringSubmatch(input, -1)
	fmt.Println(matches)
	enabled := true
	total := 0
	for _, match := range matches {
		fmt.Println(match[0])
		if match[1] == "do" {
			fmt.Println("  enabled")
			enabled = true
		} else if match[1] == "don't" {
			fmt.Println("  disabled")
			enabled = false
		} else if enabled {
			first, _ := strconv.Atoi(match[2])
			second, _ := strconv.Atoi(match[3])
			fmt.Printf("  %d * %d\n", first, second)
			total += first * second
		}
	}
	fmt.Println("Total: ", total)
}
