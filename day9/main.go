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
	// part1(input, total)
	part2(input, total)
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

// left closed/right open interval
type Interval struct {
	start, stop int
}

func size(i Interval) int {
	return i.stop - i.start
}

func checksum(i Interval, fileNum int) int {
	return (i.stop + i.start - 1) * size(i) * fileNum / 2
}

func update(board []byte, i Interval, char byte) {
	for j := i.start; j < i.stop; j++ {
		board[j] = char
	}
}

func part2(input []int, total int) {
	// process files in reverse order
	// move to first contiguous open block that's at least as large

	// need to _iterate_ files, _update_ their positions
	// need to _shrink_ or _remove_ gaps that are filled, and _join_ gaps that are vacated

	files := make([]Interval, 0)
	gaps := make([]Interval, 0)
	pos := 0
	// board := make([]byte, total)
	// var char byte
	for i, size := range input {
		interval := Interval{pos, pos + size}
		pos += size
		switch i % 2 {
		case 0:
			files = append(files, interval)
			// char = byte(i/2) + '0'
		case 1:
			gaps = append(gaps, interval)
			// char = '.'
		}
		// update(board, interval, char)
	}

	// fmt.Println(string(board))

	check := 0

	for fileNum := len(files) - 1; fileNum >= 0; fileNum-- {
		file := files[fileNum]
		fileSize := size(file)
		// fmt.Printf("File %v at %v (size %v)\n", fileNum, file, fileSize)
		for gapIdx, gap := range gaps {
			// fmt.Printf("  Gap %v at %v (size %v)\n", gapIdx, gap, size(gap))
			if gap.start >= file.start {
				break
			} else if size(gap) >= fileSize {
				// update(board, file, '.')
				file = Interval{gap.start, gap.start + fileSize}
				// update(board, file, byte(fileNum)+'0')
				gap = Interval{gap.start + fileSize, gap.stop}
				// fmt.Printf("  Moving file to %v, new gap: %v\n", file, gap)
				gaps[gapIdx] = gap
				// fmt.Println(string(board))
				break
			}
		}
		cs := checksum(file, fileNum)
		// fmt.Printf("Checksum for file %v at %v: %v\n", fileNum, file, cs)
		check += cs
	}

	fmt.Println("Total checksum:", check)

	// files can only be 1-9 in size.  gaps may become bigger than that!  we
	// only ever need to know the _first_ gap of a given size, so we can keep an
	// index of "first gap that can accomodate a file of size X" into the list
	// of gaps (e.g.,
	//
	// then for each file of size F:
	//   get the first gap that is big enough, of size G starting at g
	//   advance start of gap g+=F
	//   update the gap index:
	//     iterate gaps g' starting at g
	//     if idx[size(g')] > g', set it to g'
	//     if size(g') >= G, stop

}
