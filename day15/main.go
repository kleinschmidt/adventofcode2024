package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kleinschmidt/adventofcode2024/points"
	"github.com/kleinschmidt/adventofcode2024/queues"
)

type Board struct {
	board [][]byte
	robot points.Point
}

func parseInput() (Board, string) {
	scanner := bufio.NewScanner(os.Stdin)
	board := make([][]byte, 0)
	y := 0
	var robot points.Point
	for scanner.Scan() && scanner.Text() != "" {
		line := make([]byte, len(scanner.Bytes()))
		copy(line, scanner.Text())
		for x, c := range line {
			if c == '@' {
				robot = points.Point{X: x, Y: y}
			}
		}
		board = append(board, line)
		y++
	}

	instructions := ""
	for scanner.Scan() {
		line := scanner.Text()
		instructions += line
	}

	return Board{board, robot}, instructions
}

func printBoard(board Board) {
	for _, line := range board.board {
		fmt.Println(string(line))
	}
}

var directions = map[byte]points.Point{
	'v': {X: 0, Y: 1},
	'>': {X: 1, Y: 0},
	'<': {X: -1, Y: 0},
	'^': {X: 0, Y: -1},
}

func (b *Board) move(dir points.Point) {
	cur := b.robot
	b.robot = move(b.board, cur, dir)
}

func move(board [][]byte, cur points.Point, dir points.Point) points.Point {
	next := points.Add(cur, dir)
	p := *points.Get(board, next)
	// fmt.Printf("Moving: %v -> %v=%s (%v)\n", cur, next, string(p), dir)
	switch p {
	case '#':
	// nothing to do here
	case '.':
		// we can move into this space; swap cur and next on the board
		board[cur.Y][cur.X], board[next.Y][next.X] = board[next.Y][next.X], board[cur.Y][cur.X]
		// return next to indicate that we moved successfully
		cur = next
	case 'O':
		// must check that this crate can move
		if move(board, next, dir) != next {
			board[cur.Y][cur.X], board[next.Y][next.X] = board[next.Y][next.X], board[cur.Y][cur.X]
			// return next to indicate that we moved successfully
			cur = next
		}
	}
	return cur
}

func main() {
	board, instructions := parseInput()
	printBoard(board)
	fmt.Println("\n", board.robot, instructions)

	// part1(board, instructions)
	part2(board, instructions)
}

func part1(board Board, instructions string) {
	for _, inst := range instructions {
		dir := directions[byte(inst)]
		// fmt.Printf("Instruction %d: %s (%v)\n", i, string(inst), dir)
		board.move(dir)
		// printBoard(board)
		// fmt.Println(board.robot)
	}

	total := 0
	for y, line := range board.board {
		for x, c := range line {
			if c == 'O' {
				total += 100*y + x
			}
		}
	}

	fmt.Println("Total:", total)
}

func expandBoard(board Board) Board {
	newb := make([][]byte, 0)
	for _, line := range board.board {
		newLine := make([]byte, len(line)*2)
		for i, c := range line {
			var c0, c1 byte
			switch c {
			case 'O':
				c0 = '['
				c1 = ']'
			case '@':
				c0 = '@'
				c1 = '.'
			default:
				c0 = c
				c1 = c
			}
			newLine[i*2] = c0
			newLine[i*2+1] = c1
		}
		newb = append(newb, newLine)
	}
	return Board{newb, points.Point{X: board.robot.X * 2, Y: board.robot.Y}}
}

func part2(board Board, instructions string) {
	board = expandBoard(board)
	fmt.Println(board.robot)
	printBoard(board)

	for _, inst := range instructions {
		dir := directions[byte(inst)]
		// fmt.Printf("Instruction %d: %s (%v)\n", i, string(inst), dir)
		board.move2(dir)
		// printBoard(board)
		// fmt.Println(board.robot)
		// fmt.Println(string(inst))
		// printBoard(board)
	}

	printBoard(board)

	total := 0
	for y, line := range board.board {
		for x, c := range line {
			if c == '[' {
				total += 100*y + x
			}
		}
	}

	fmt.Println("Total:", total)
}

/*
need to split up the move into two steps:
1. figure out if we can move
2. do the move

both need to do a search of all connected positions

basically we need to find all the "columns" (starting points to move from).
*/

func (board *Board) move2(dir points.Point) {
	// which points need to move
	moves := make(map[points.Point]byte)
	work := queues.NewQueue([]points.Point{board.robot})
	for work.HasNext() {
		cur := *work.Dequeue()

		switch *points.Get(board.board, cur) {
		case '.':
			// nothing to do here
			continue
		case '#':
			// can't move!
			return
		case '[': // depending on direction: need to add RHS and cur + dir to queue
			otherHalf := points.Add(cur, points.Point{X: 1, Y: 0})
			if *points.Get(board.board, otherHalf) != ']' {
				panic("bad state!")
			}
			_, ok := moves[otherHalf]
			if !ok {
				work.Enqueue(otherHalf)
			}
		case ']': // depending on direction: need to add LHS and cur + dir to queue
			otherHalf := points.Add(cur, points.Point{X: -1, Y: 0})
			if *points.Get(board.board, otherHalf) != '[' {
				panic("bad state!")
			}
			_, ok := moves[otherHalf]
			if !ok {
				work.Enqueue(otherHalf)
			}
		case '@':
		}

		moves[cur] = *points.Get(board.board, cur)
		next := points.Add(cur, dir)
		work.Enqueue(next)
	}

	// first, clear all existing contents of src. tiles
	for src := range moves {
		board.board[src.Y][src.X] = '.'
	}

	// then make the moves
	for src, c := range moves {
		dest := points.Add(src, dir)
		board.board[dest.Y][dest.X] = c
	}

	robot := board.robot
	board.board[robot.Y][robot.X] = '.'

	board.robot = points.Add(robot, dir)
}
