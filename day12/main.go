package main

import (
	"bufio"
	"fmt"
	"iter"
	"os"

	"github.com/kleinschmidt/adventofcode2024/queues"
)

type Point struct{ x, y int }

func parseInput() ([][]byte, Point) {
	scanner := bufio.NewScanner(os.Stdin)
	y := 0
	board := make([][]byte, 0)
	for scanner.Scan() {
		line := scanner.Bytes()
		board = append(board, line)
		y++
	}
	return board, Point{len(board[0]), y}
}

func main() {
	board, size := parseInput()
	for _, line := range board {
		fmt.Println(line)
	}
	fmt.Println(size)

	part1(board, size)
}

func nextUnvisited(visited [][]bool) Point {
	for y, line := range visited {
		for x, v := range line {
			if !v {
				return Point{x, y}
			}
		}
	}
	// sentinel value of size/OOB corner
	return Point{len(visited[0]), len(visited)}
}

func inbounds(p Point, board [][]byte) bool {
	return p.x >= 0 && p.x < len(board[0]) && p.y >= 0 && p.y < len(board)
}

func add(a, b Point) Point {
	return Point{a.x + b.x, a.y + b.y}
}

func neighbors(p Point) iter.Seq[Point] {
	directions := [4]Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	return func(yield func(Point) bool) {
		for _, d := range directions {
			pn := add(p, d)
			if !yield(pn) {
				return
			}
		}
		return
	}
}

func get(board [][]byte, p Point) byte {
	if inbounds(p, board) {
		return board[p.y][p.x]
	} else {
		return byte(0)
	}
}

func visit(start Point, board [][]byte, visited [][]bool) (area, perimeter int) {
	area = 0
	perimeter = 0

	queue := new(queues.Queue[Point])
	queue.Enqueue(start)

	for queue.HasNext() {
		cur := queue.Dequeue()
		// fmt.Println(cur)
		if visited[cur.y][cur.x] {
			// fmt.Println("visited")
			continue
		}
		for next := range neighbors(*cur) {
			// fmt.Print("  ", next)
			if !inbounds(next, board) || board[next.y][next.x] != board[cur.y][cur.x] {
				// every out-of-bounds point can only be visited from one in-bounds point
				perimeter++
			} else if !visited[next.y][next.x] {
				queue.Enqueue(next)
			}
		}
		visited[cur.y][cur.x] = true
		area++
	}

	return area, perimeter
}

func part1(board [][]byte, size Point) {
	visited := make([][]bool, size.y)
	for i := range visited {
		visited[i] = make([]bool, size.x)
	}

	total := 0

	for {
		next := nextUnvisited(visited)
		if next == size {
			break
		}
		fmt.Println("Visiting", next)
		area, perimeter := visit(next, board, visited)
		fmt.Println("Area", area, "perimeter", perimeter)
		total += area * perimeter
	}

	fmt.Println("total:", total)
}
