package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

type Point struct {
	x, y int
}

type Guard struct {
	location  Point
	direction Point
}

type Board struct {
	obstacles map[Point]bool
	dim       Point
	guard     Guard
	visited   map[Guard]bool
}

func (b *Board) Inbounds(pos Point) bool {
	return pos.x >= 0 && pos.x < b.dim.x && pos.y >= 0 && pos.y < b.dim.y
}

func Add(a, b Point) Point {
	return Point{a.x + b.x, a.y + b.y}
}

func (b *Board) Next() bool {
	cur := b.guard.location
	next := Add(cur, b.guard.direction)
	if b.obstacles[next] {
		b.guard = b.guard.rotCw()
		// next = Add(cur, b.guard.direction)
	} else {
		b.guard.location = next
	}
	// fmt.Print("Next: ", next)
	// defer fmt.Println()
	if !b.Inbounds(next) {
		// fmt.Print(" (out of bounds)")
		return false
	} else if b.visited[b.guard] {
		// fmt.Print(" (visited already)")
		return false
	} else {
		b.visited[b.guard] = true
		return true
	}
}

func parseInput(buff []byte) Board {
	scanner := bufio.NewScanner(bytes.NewReader(buff))
	y := 0
	obstacles := make(map[Point]bool)
	var x int
	var c rune
	var guard Guard
	for scanner.Scan() {
		for x, c = range scanner.Text() {
			loc := Point{x, y}
			switch c {
			case '#':
				obstacles[loc] = true
			case '^':
				guard = Guard{loc, Point{0, -1}}
			case 'v':
				guard = Guard{loc, Point{0, 1}}
			case '>':
				guard = Guard{loc, Point{1, 0}}
			case '<':
				guard = Guard{loc, Point{-1, 0}}
			}
		}
		y += 1
	}
	return Board{obstacles, Point{x + 1, y}, guard, make(map[Guard]bool)}
}

func (b *Board) String() string {
	buff := make([]byte, 0, b.dim.y*(b.dim.x+1))
	var char byte
	for y := range b.dim.y {
		for x := range b.dim.x {
			switch {
			case b.obstacles[Point{x, y}]:
				char = '#'
			case Point{x, y} == b.guard.location:
				switch b.guard.direction {
				case Point{0, -1}:
					char = '^'
				case Point{0, 1}:
					char = 'v'
				case Point{1, 0}:
					char = '>'
				case Point{-1, 0}:
					char = '<'
				}
			default:
				char = '.'
			}
			buff = append(buff, char)
		}
		buff = append(buff, '\n')
	}
	return string(buff)
}

func (g *Guard) rotCw() Guard {
	direction := Point{-g.direction.y, g.direction.x}
	return Guard{g.location, direction}
}

func clear() {
	fmt.Print("\033[2J")
}

func main() {
	clear()
	buff, _ := io.ReadAll(os.Stdin)
	fmt.Println("Input:")
	fmt.Println(string(buff))

	clear()
	board := parseInput(buff)
	fmt.Println("Parsed:")
	fmt.Println(board.String())

	part1(parseInput(buff), 0*time.Millisecond)
	// time.Sleep(1)
	part2(parseInput(buff))
}

func part1(board Board, sleep time.Duration) {
	i := 0
	for board.Next() {
		i++
		if sleep > 0 {
			time.Sleep(sleep)
			clear()
			fmt.Println("Step", i)
			fmt.Println(board.String())
		}
	}
	if board.visited[board.guard] {
		fmt.Println("Stopped because visited already")
	}

	visited := make(map[Point]bool)
	for guard := range board.visited {
		visited[guard.location] = true
	}
	fmt.Println("Visited a total of", len(visited), "locations")
	time.Sleep(time.Second)
}

func part2(board Board) {
	boardStr := board.String()

	// run the board once to narrow down the places we have to check...
	fmt.Println("Running board again...")
	track := make([]Guard, 0)
	for board.Next() {
		track = append(track, board.guard)
	}
	fmt.Println("Checking ", len(track), " candidates")
	checked := make(map[Point]bool)

	// start at the second position (we can't add the obstacle to the start anyway)...
	for _, g := range track[1:] {
		obstacle := g.location
		if _, ok := checked[obstacle]; ok {
			// this was on a previous state in the track, meaning that if we put
			// an obstacle here we may not reach this state
			// continue
		}
		board := parseInput([]byte(boardStr))
		// this is the _previous_ step!  start there.
		// start := track[i]
		// board.guard = start
		// we'll put an obstacle at the current location
		board.obstacles[obstacle] = true
		for board.Next() {
			// run out the board
		}
		if board.visited[board.guard] {
			// fmt.Println("Iter ", i, "\n", boardStart)
			checked[g.location] = true
		} else {
			// checked[g.location] = false
		}
	}

	total := 0
	for _, good := range checked {
		if good {
			total++
		}
	}

	fmt.Println("Found ", total, len(checked), " positions for obstacles")
}
