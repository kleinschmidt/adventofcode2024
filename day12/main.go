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

func visit(start Point, board [][]byte, visited [][]bool) (area, perimeter int, region map[Point]bool) {
	area = 0
	perimeter = 0

	region = make(map[Point]bool)

	queue := new(queues.Queue[Point])
	visited[start.y][start.x] = true
	region[start] = true
	queue.Enqueue(start)

	for queue.HasNext() {
		cur := queue.Dequeue()
		for next := range neighbors(*cur) {
			if !inbounds(next, board) || board[next.y][next.x] != board[cur.y][cur.x] {
				perimeter++
			} else if !visited[next.y][next.x] {
				visited[next.y][next.x] = true
				region[next] = true
				queue.Enqueue(next)
			} else {
				// visited, inbounds, and same species: should be in region!
				if !region[next] {
					panic(fmt.Sprintf("Expected %v in region", next))
				}
			}
		}
		area++
	}

	return area, perimeter, region
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
	total_regions := 0

	regions := make([]map[Point]bool, 0)

	for {
		next := nextUnvisited(visited)
		if next == size {
			break
		}
		area, perimeter, region := visit(next, board, visited)
		regions = append(regions, region)
		total_area += area
		total_perimeter += perimeter
		fmt.Printf("Plot %v %v: area=%d, perimeter=%d\n", string(get(board, next)), next, area, perimeter)
		total += area * perimeter
		total_regions += 1
	}

	// printBoard(board, visited)

	fmt.Println("total:", total)
	fmt.Println("total_area:", total_area, "expected", size.x*size.y)
	fmt.Println("total_perimeter:", total_perimeter, "expected", totalPerimeter(board))
	fmt.Println("total_regions:", total_regions)

	checkRegions(board, regions)
}

func checkRegions(board [][]byte, regions []map[Point]bool) {
	for _, region := range regions {
		for p := range region {
			neighbors := []Point{{p.x, p.y + 1}, {p.x + 1, p.y}, {p.x, p.y - 1}, {p.x - 1, p.y}}
			for _, n := range neighbors {
				if n.y < 0 || n.y >= len(board) || n.x < 0 || n.x >= len(board[0]) {
					// out of bounds
					continue
				} else if !region[n] && (board[p.y][p.x] == board[n.y][n.x]) {
					fmt.Println("Regions should be merged:", p, n, string(board[p.y][p.x]), string(board[n.y][n.x]))
				}
			}
		}
	}
}
