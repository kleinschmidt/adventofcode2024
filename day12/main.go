package main

import (
	"bufio"
	"fmt"
	"iter"
	"os"

	"github.com/kleinschmidt/adventofcode2024/queues"
)

type Point struct{ x, y int }
type Vector struct{ loc, dir Point }

func parseInput() ([][]byte, Point) {
	scanner := bufio.NewScanner(os.Stdin)
	y := 0
	board := make([][]byte, 0)
	for scanner.Scan() {
		// real input is big enough that it fills the buffer so we must copy to
		// avoid overwriting early input with later input...
		bytes := scanner.Bytes()
		line := make([]byte, len(bytes))
		copy(line, bytes)
		// fmt.Println(string(line))
		board = append(board, line)
		y++
	}
	return board, Point{len(board[0]), y}
}

func main() {
	board, size := parseInput()
	part1(board, size)
	part2(board, size)
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

func clockwise(p Point) Point {
	return Point{p.y, -p.x}
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

func get(board [][]byte, p Point) byte {
	if inbounds(p, board) {
		return board[p.y][p.x]
	} else {
		// use 0 as a sentinel value
		return byte(0)
	}
}

func visit(start Point, board [][]byte, visited [][]bool) (area int, edges map[Vector]bool) {
	area = 0
	edges = make(map[Vector]bool)

	queue := new(queues.Queue[Point])
	visited[start.y][start.x] = true
	queue.Enqueue(start)

	for queue.HasNext() {
		cur := *queue.Dequeue()
		for dir, next := range neighbors(cur) {
			if !inbounds(next, board) || board[next.y][next.x] != board[cur.y][cur.x] {
				edges[Vector{cur, dir}] = true
			} else if !visited[next.y][next.x] {
				visited[next.y][next.x] = true
				queue.Enqueue(next)
			}
		}
		area++
	}

	return area, edges
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
		area, edges := visit(next, board, visited)
		total += area * len(edges)
	}

	fmt.Println("total:", total)
}

func sides(edges map[Vector]bool) int {
	total := 0

	for edge := range edges {
		rh := Vector{add(edge.loc, clockwise(edge.dir)), edge.dir}
		if !edges[rh] {
			total++
		}
	}

	return total
}

func part2(board [][]byte, size Point) {
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
		area, edges := visit(next, board, visited)
		// we only count edges that have no "right-hand" neighbor

		total += area * sides(edges)
	}

	fmt.Println("total:", total)
}
