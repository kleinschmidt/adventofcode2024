package main

import (
	"bufio"
	"fmt"
	"iter"
	"os"
)

type Point struct{ x, y int }

func add(a, b Point) Point {
	return Point{a.x + b.x, a.y + b.y}
}

func inbounds(topo [][]byte, pos Point) bool {
	return pos.y >= 0 && pos.y < len(topo) && pos.x >= 0 && pos.x < len(topo[0])
}

func parseInput() ([][]byte, []Point) {
	scanner := bufio.NewScanner(os.Stdin)
	board := make([][]byte, 0)
	trailheads := make([]Point, 0)
	y := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		parsed := make([]byte, len(line))
		for x, c := range line {
			parsed[x] = c - '0'
			if parsed[x] == 0 {
				trailheads = append(trailheads, Point{x, y})
			}
		}
		board = append(board, parsed)
		y++
	}
	return board, trailheads
}

func main() {
	topo, trailheads := parseInput()
	for _, line := range topo {
		fmt.Println(line)
	}
	fmt.Println(trailheads)

	score1 := 0
	score2 := 0

	for _, trailhead := range trailheads {
		fmt.Println("Searching from trialhead", trailhead)
		summits := search(topo, trailhead)
		fmt.Println("Found summits:\n", summits)
		fmt.Println("Trailhead", trailhead, "score", len(summits))
		score1 += len(summits)
		for _, v := range summits {
			score2 += v
		}
	}

	fmt.Println("Total score (part 1):", score1)
	fmt.Println("Total score (part 2):", score2)
}

type Elem[E comparable] struct {
	e    E
	next *Elem[E]
}

type Queue[E comparable] struct {
	first *Elem[E]
	last  *Elem[E]
}

func (queue *Queue[E]) Enqueue(e E) {
	elem := &Elem[E]{e, nil}
	if queue.first == nil {
		queue.first = elem
		queue.last = elem
	} else {
		queue.last.next = elem
		queue.last = queue.last.next
	}
}

func (queue *Queue[E]) Dequeue() *E {
	first := queue.first
	if first == nil {
		return nil
	}
	queue.first = first.next
	if first.next == nil {
		queue.last = nil
	}
	return &first.e
}

// this is a bit iffy: it's destructive!
func (queue *Queue[E]) All() iter.Seq[E] {
	return func(yield func(E) bool) {
		for {
			next := queue.Dequeue()
			if next == nil || !yield(*next) {
				return
			}
		}
	}
}

func NewQueue[E comparable](arr []E) *Queue[E] {
	queue := new(Queue[E])
	for _, e := range arr {
		queue.Enqueue(e)
	}
	return queue
}

func search(topo [][]byte, trailhead Point) map[Point]int {
	queue := NewQueue([]Point{trailhead})
	summits := make(map[Point]int)

	for next := range queue.All() {
		height := topo[next.y][next.x]
		// fmt.Printf("Searching %v (height %v)\n", next, height)
		if height == 9 {
			summits[next]++
		}
		for _, vec := range []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			candidate := add(next, vec)
			if inbounds(topo, candidate) && topo[candidate.y][candidate.x]-height == 1 {
				// fmt.Printf("  queueing %v\n", candidate)
				queue.Enqueue(candidate)
			}
		}
	}

	return summits
}
