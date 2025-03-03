package main

import (
	"bufio"
	"flag"
	"fmt"
	"iter"
	"os"
	"regexp"
	"strconv"
)

type Point struct{ x, y int }

type Robot struct{ pos, vel Point }

func MustInt(x string) int {
	i, err := strconv.Atoi(x)
	if err != nil {
		panic(fmt.Sprint("Cannot convert string to int:", x))
	}
	return i
}

func parseInput() []Robot {
	scanner := bufio.NewScanner(os.Stdin)
	re := regexp.MustCompile("p=([-0-9]+),([-0-9]+) v=([-0-9]+),([-0-9]+)")
	robots := make([]Robot, 0)
	for scanner.Scan() {
		matches := re.FindStringSubmatch(scanner.Text())
		pos := Point{MustInt(matches[1]), MustInt(matches[2])}
		vel := Point{MustInt(matches[3]), MustInt(matches[4])}
		robots = append(robots, Robot{pos, vel})
	}
	return robots
}

func pmod(a, b int) int {
	r := a % b
	if r < 0 {
		r += b
	}
	return r
}

func (r *Robot) Move(steps int, board Point) Robot {
	r.pos = Point{
		pmod(r.pos.x+steps*r.vel.x, board.x),
		pmod(r.pos.y+steps*r.vel.y, board.y),
	}
	return *r
}

func counts(robots []Robot) map[Point]int {
	counts := make(map[Point]int)
	for _, r := range robots {
		counts[r.pos] += 1
	}
	return counts
}

func printBoard(robots []Robot, size Point) {
	counts := counts(robots)
	for y := range size.y {
		for x := range size.x {
			c := counts[Point{x, y}]
			if c > 0 {
				fmt.Print(c)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	var x, y int
	flag.IntVar(&x, "w", 0, "")
	flag.IntVar(&y, "h", 0, "")
	flag.Parse()
	size := Point{x, y}
	robots := parseInput()

	part1(robots, size)
	part2(robots, size)
}

func quadrant(p Point, size Point) Point {
	mid := Point{size.x / 2, size.y / 2}
	var qx, qy int
	if p.x > mid.x {
		qx = 1
	} else if p.x < mid.x {
		qx = -1
	}

	if p.y > mid.y {
		qy = 1
	} else if p.y < mid.y {
		qy = -1
	}

	if qx == 0 || qy == 0 {
		return Point{0, 0}
	} else {
		return Point{qx, qy}
	}
}

func part1(robotsIn []Robot, size Point) {
	printBoard(robotsIn, size)
	robots := make([]Robot, len(robotsIn))
	for i, r := range robotsIn {
		robots[i] = r.Move(100, size)
	}
	fmt.Println("--------")
	printBoard(robots, size)

	c := counts(robots)
	quadcounts := make(map[Point]int)
	for p, x := range c {
		q := quadrant(p, size)
		if (q != Point{0, 0}) {
			quadcounts[q] += x
		}
	}

	fmt.Println(quadcounts)
	total := 1
	for _, c := range quadcounts {
		total *= c
	}
	fmt.Println("Safety factor:", total)
}

func add(a, b Point) Point {
	return Point{a.x + b.x, a.y + b.y}
}

func neighbors(p Point) iter.Seq2[Point, Point] {
	directions := [4]Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	return func(yield func(Point, Point) bool) {
		for _, d := range directions {
			pn := add(p, d)
			if !yield(d, pn) {
				return
			}
		}
		return
	}
}

func robotsWithNeighbors(robots []Robot) int {
	occupancy := make(map[Point]bool)
	count := 0
	for _, r := range robots {
		occupancy[r.pos] = true
	}
	for _, r := range robots {
		for _, n := range neighbors(r.pos) {
			if occupancy[n] {
				count++
				break
			}
		}
	}
	return count
}

func part2(robots []Robot, size Point) {
	printBoard(robots, size)
	count := robotsWithNeighbors(robots)
	fmt.Println("robots with neighbors:", count, "/", len(robots))
	nIter := 0
	maxSoFar := count
	for count < len(robots)/2 {
		nIter++
		for i := range robots {
			robots[i].Move(1, size)
		}
		count = robotsWithNeighbors(robots)
		if count > maxSoFar {
			maxSoFar = count
		}
		// fmt.Println(nIter, ": robots with neighbors:", count, "/", len(robots), "(max so far:", maxSoFar, ")")
	}
	fmt.Println("--------")
	printBoard(robots, size)
	fmt.Println("Seconds elapsed:", nIter)

}
