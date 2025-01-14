package main

import (
	"bufio"
	"fmt"
	"os"

	"gonum.org/v1/gonum/stat/combin"
)

type Point struct{ x, y int }

func isAlphaNum(c rune) bool {
	return 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9'
}

func antinodes(a, b Point) (aa, bb Point) {
	delta := Point{a.x - b.x, a.y - b.y}
	aa = Point{a.x + delta.x, a.y + delta.y}
	bb = Point{b.x - delta.x, b.y - delta.y}
	return
}

func resonances(a, b, corner Point) []Point {
	points := make([]Point, 0)
	delta := Point{a.x - b.x, a.y - b.y}
	g := gcd(delta.x, delta.y)
	if g > 1 {
		delta = Point{delta.x / g, delta.y / g}
	}
	// fmt.Println(a, b, delta)
	for inbounds(a, corner) {
		// fmt.Println("  ", a)
		points = append(points, a)
		a = Point{a.x + delta.x, a.y + delta.y}
	}
	for inbounds(b, corner) {
		// fmt.Println("  ", b)
		points = append(points, b)
		b = Point{b.x - delta.x, b.y - delta.y}
	}
	return points
}

func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	if a == b {
		return a
	} else if a < b {
		return gcd(a, b-a)
	} else {
		return gcd(a-b, b)
	}
}

func inbounds(p, corner Point) bool {
	return p.x >= 0 && p.x < corner.x && p.y >= 0 && p.y < corner.y
}

func parseInput() (locations map[rune][]Point, corner Point) {
	scanner := bufio.NewScanner(os.Stdin)
	locations = make(map[rune][]Point)
	y := 0
	var x int
	var c rune
	for scanner.Scan() {
		line := scanner.Text()
		for x, c = range line {
			if isAlphaNum(c) {
				loc, ok := locations[c]
				if !ok {
					loc = make([]Point, 0)
				}
				locations[c] = append(loc, Point{x, y})
			}
		}
		y++
	}
	return locations, Point{x + 1, y}
}

func part1(locations map[rune][]Point, corner Point) {
	antis := make(map[Point]bool)
	for _, locs := range locations {
		// fmt.Printf("Processing antinodes for %v (%d locations)\n", c, len(locs))
		pairs := combin.NewCombinationGenerator(len(locs), 2)
		pair := make([]int, 2)
		for pairs.Next() {
			pairs.Combination(pair)
			a, b := locs[pair[0]], locs[pair[1]]
			aa, bb := antinodes(a, b)
			if inbounds(aa, corner) {
				antis[aa] = true
			}
			if inbounds(bb, corner) {
				antis[bb] = true
			}
			// fmt.Printf("  Processing pair: %v <-> %v, antinodes %v (%v) and %v (%v)\n", a, b, aa, inbounds(aa, corner), bb, inbounds(bb, corner))
		}
	}
	//
	fmt.Printf("Found %d antinodes\n", len(antis))
}

func part2(locations map[rune][]Point, corner Point) {
	antis := make(map[Point]bool)
	for _, locs := range locations {
		pairs := combin.NewCombinationGenerator(len(locs), 2)
		pair := make([]int, 2)
		for pairs.Next() {
			pairs.Combination(pair)
			a, b := locs[pair[0]], locs[pair[1]]
			for _, p := range resonances(a, b, corner) {
				antis[p] = true
			}
		}
	}
	fmt.Printf("Found %d antinodes\n", len(antis))
}

func main() {
	input, corner := parseInput()
	fmt.Println(input)
	fmt.Println(corner)
	part1(input, corner)
	fmt.Println(gcd(-48, 18))
	part2(input, corner)
}
