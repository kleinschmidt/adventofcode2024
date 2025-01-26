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
		fmt.Println(string(line))
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
		// use 0 as a sentinel value
		return byte(0)
	}
}

func visit(start Point, board [][]byte, visited [][]bool) (area, perimeter int) {
	area = 0
	perimeter = 0

	// stack := []Point{start}
	queue := new(queues.Queue[Point])
	queue.Enqueue(start)

	for queue.HasNext() {
		cur := queue.Dequeue()
		// for len(stack) > 0 {
		// 	// pointer just to match the queue
		// 	cur := &stack[len(stack)-1]
		// 	stack = stack[:len(stack)-1]
		// sicne this location was enqueued, it may have been visited from
		// another neighbor.
		if visited[cur.y][cur.x] {
			continue
		}
		for next := range neighbors(*cur) {
			// fmt.Print("  ", next)
			if !inbounds(next, board) || board[next.y][next.x] != board[cur.y][cur.x] {
				// every out-of-bounds point can only be visited from one
				// in-bounds point ...wait that's not true, if you have a
				// u-shaped region.  but still, every edge that has a different
				// species is a length of fence you need...
				perimeter++
			} else if !visited[next.y][next.x] {
				queue.Enqueue(next)
				// stack = append(stack, next)
			}
		}
		visited[cur.y][cur.x] = true
		area++
	}

	return area, perimeter
}

func printBoard(board [][]byte, visited [][]bool) {
	for y := range board {
		for x := range board[y] {
			if visited[y][x] {
				fmt.Print(".")
			} else {
				fmt.Print(string(board[y][x]))
			}
		}
		println()
	}
}

func totalPerimeter(board [][]byte) int {
	total := 0
	for y := range board {
		for x := range board[y] {
			cur := Point{x, y}
			for neighbor := range neighbors(cur) {
				if get(board, neighbor) != get(board, cur) {
					total++
				}
			}
		}
	}
	return total
}

func part1(board [][]byte, size Point) {
	visited := make([][]bool, size.y)
	for i := range visited {
		visited[i] = make([]bool, size.x)
	}

	total := 0
	total_perimeter := 0
	total_area := 0

	for {
		next := nextUnvisited(visited)
		if next == size {
			break
		}
		area, perimeter := visit(next, board, visited)
		total_area += area
		total_perimeter += perimeter
		fmt.Printf("Plot %v %v: area=%d, perimeter=%d\n", get(board, next), next, area, perimeter)
		total += area * perimeter
	}

	// printBoard(board, visited)

	fmt.Println("total:", total)
	fmt.Println("total_area:", total_area, "expected", size.x*size.y)
	fmt.Println("total_perimeter:", total_perimeter, "expected", totalPerimeter(board))
}
