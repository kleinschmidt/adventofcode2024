package main

import (
	"bytes"
	"fmt"
	"io"
	"iter"
	"os"
	"strconv"
	"strings"
)

func parseInput() []string {
	line, _ := io.ReadAll(os.Stdin)
	stones := strings.Fields(string(line))
	return stones
}

type List[E any] struct {
	prev *List[E]
	e    E
	next *List[E]
}

func NewList[E any](arr []E) *List[E] {
	if len(arr) == 0 {
		return nil
	}
	list := &List[E]{nil, arr[0], nil}
	cur := list
	for _, e := range arr[1:] {
		cur.next = &List[E]{cur, e, nil}
		cur = cur.next
	}
	return list
}

func (l *List[E]) All() iter.Seq[E] {
	return func(yield func(E) bool) {
		cur := l
		for cur != nil {
			// pre-assign this inc ase `yield` does anything to `next`
			next := cur.next
			if !yield(cur.e) {
				return
			}
			cur = next
		}
	}
}

func (l *List[E]) AllList() iter.Seq[*List[E]] {
	return func(yield func(*List[E]) bool) {
		cur := l
		for cur != nil {
			// pre-assign this inc ase `yield` does anything to `next`
			next := cur.next
			if !yield(cur) {
				return
			}
			cur = next
		}
	}
}

func (l List[E]) String() string {
	buf := new(bytes.Buffer)
	fmt.Fprint(buf, "List[ ")
	for e := range l.All() {
		fmt.Fprint(buf, e, " ")
	}
	fmt.Fprint(buf, "]")
	return buf.String()
}

func main() {
	stones := parseInput()
	fmt.Println(stones)
	list := NewList(stones)
	fmt.Println(list)

	part1(list)
}

func timeLeadingZeros(x string) string {
	for len(x) > 1 && x[0] == '0' {
		x = x[1:]
	}
	return x
}

func part1(stones *List[string]) {
	var n int
	for i := range 25 {
		n = 0
		for l := range stones.AllList() {
			// always have at least stone!
			n++
			switch {
			case l.e == "0":
				l.e = "1"
			case len(l.e)%2 == 0:
				// split this one so we have an extra output
				n++
				half := len(l.e) / 2
				first := l.e[:half]
				second := timeLeadingZeros(l.e[half:])
				l.e = first
				l.next = &List[string]{l, second, l.next}
			case len(l.e)%2 == 1:
				i, err := strconv.Atoi(l.e)
				if err != nil {
					fmt.Println("Error converting:", err)
				}
				l.e = strconv.Itoa(i * 2024)
			}
		}
		fmt.Println(i+1, "blinks,", n, "stones")
	}
	fmt.Println("Total stones:", n)
}

func part2(stones []string) {
	cache := make(map[int][]int)
	for stone := range stones {

	}
}
