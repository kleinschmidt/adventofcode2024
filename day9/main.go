package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func feedSpaces(c chan int, input []int) {
	// we need to know the index of each empty space
	i := 0
	pos := 0
	for i < len(input)-1 {
		// first one is a file
		fmt.Printf("Pos %d, skipping %d -> ", pos, input[i])
		pos += input[i]
		fmt.Print(pos)
		i++
		fmt.Printf(". Feeding %d: ", input[i])
		for range input[i] {
			pos += 1
			fmt.Printf("%d ", pos)
			c <- pos
		}
		fmt.Println()
		i++
	}
	close(c)
}

func feedFileBlocks(c chan [2]int, input []int, total int) {
	// work backwards from the last position
	pos := total
	for i := len(input) - 1; i >= 0; i-- {
		switch i % 2 {
		case 0:
			file := i / 2
			for range input[i] {
				// pre-decrement because we've started at the position _after_
				// the last block
				pos -= 1
				c <- [2]int{pos, file}
			}
		case 1:
			pos -= input[i]
		}
	}
	close(c)
}

func parseInput() ([]int, int) {
	input, _ := io.ReadAll(os.Stdin)
	str := strings.TrimSpace(string(input))
	out := make([]int, len(str))
	total := 0
	for i, c := range str {
		out[i] = int(c - '0')
		total += out[i]
	}
	return out, total
}

func main() {
	input, total := parseInput()
	part1(input, total)
}

func feed(c chan int, input []int, revPos int) {
	fileBlocks := make(chan [2]int)
	go feedFileBlocks(fileBlocks, input, revPos)
	pos := 0
	file := 0
	for i, blockLen := range input {
		for range blockLen {
			switch i % 2 {
			// evens are files
			case 0:
				file = i / 2
			// odds are gaps
			case 1:
				fileBlock, ok := <-fileBlocks
				if ok {
					revPos = fileBlock[0]
					file = fileBlock[1]
				}
			}
			c <- file
			pos++
			if pos >= revPos {
				close(c)
				return
			}
		}
	}
	close(c)
	return
}

func part1(input []int, revPos int) int {
	total := 0
	files := make(chan int)
	go feed(files, input, revPos)
	pos := 0
	for file := range files {
		total += file * pos
		pos++
	}
	fmt.Println("Total:", total)
	return total
}
