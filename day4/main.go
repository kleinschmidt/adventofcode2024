package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// bytes, _ := io.ReadAll(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0)
	for scanner.Scan() {
		// fmt.Printf("%s\n", scanner.Bytes())
		lines = append(lines, scanner.Text())
	}
	// fmt.Println(lines)
	// fmt.Println("Is it S: ", lines[1][1] == 'S')

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	total := 0
	for x, line := range lines {
		for y, _ := range line {
			for dx := -1; dx <= 1; dx += 1 {
				for dy := -1; dy <= 1; dy += 1 {
					total += search(lines, x, y, dx, dy, "XMAS")
				}
			}
		}
	}
	fmt.Println("Total \"XMAS\": ", total)
}

func search(lines []string, x int, y int, dx int, dy int, text string) int {
	if len(text) == 0 {
		// fmt.Printf("Found at %d, %d (delta %d, %d)\n", x, y, dx, dy)
		return 1
	} else if dx == 0 && dy == 0 || x < 0 || x >= len(lines) || y < 0 || y >= len(lines[x]) {
		return 0
	} else if lines[x][y] == text[0] {
		return search(lines, x+dx, y+dy, dx, dy, text[1:])
	} else {
		return 0
	}
}

func part2(lines []string) {
	total := 0
	for x, line := range lines {
		for y, _ := range line {
			for dx := -1; dx <= 1; dx += 2 {
				for dy := -1; dy <= 1; dy += 2 {
					if search(lines, x, y, dx, dy, "MAS") == 1 {
						total += search(lines, x+2*dx, y, -dx, dy, "MAS")
						total += search(lines, x, y+2*dy, dx, -dy, "MAS")
					}
				}
			}
		}
	}
	// we've double counted so divide by two
	fmt.Println("Total \"X-MAS\": ", total/2)
}
