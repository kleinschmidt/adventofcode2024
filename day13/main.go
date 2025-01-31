package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// problems are of the form:
//
// A: (Ax, Ay)
// B: (Bx, By)
//
// target: (x, y)
//
// we need to find a, b such that
//
// a * A + b * B = (x, y)
//
// or
// [A' B'] [a; b] = [x; y]
//
// so
//
// [a; b] = [A' B']^-1 [x; y]
//
// [A' B'] = [Ax Bx
//            Ay By]
//
// so [A' B']^-1 = 1/(Ax*By - Bx*Ay) [By -Bx
//                                    -Ay Ax]
//
// if the determinant is zero, the matrix is rank deficient and we have to
// determine whether either of the individual sets of coefficients is a
// solution, and if both are, which is the "cheaper" one.
//
// if the determinant is non-zero, then we have to see whether the (unique)
// coefficients are integral; if not, there is no solution!
//
// looking at the input, I don't _think_ we're at risk for integer overflow...

type Point struct{ x, y int }

func MustInt(x string) int {
	i, err := strconv.Atoi(x)
	if err != nil {
		panic(fmt.Sprint("Cannot convert string to int:", x))
	}
	return i
}

type Machine struct{ a, b, prize Point }

var re = regexp.MustCompile("X[+=]([0-9]+), Y[+=]([0-9]+)")

func parseLine(line string) Point {
	matches := re.FindStringSubmatch(line)
	return Point{MustInt(matches[1]), MustInt(matches[2])}
}

func parseInput() []Machine {
	scanner := bufio.NewScanner(os.Stdin)
	machines := make([]Machine, 0)
	for scanner.Scan() {
		A := parseLine(scanner.Text())
		scanner.Scan()
		B := parseLine(scanner.Text())
		scanner.Scan()
		prize := parseLine(scanner.Text())
		machines = append(machines, Machine{A, B, prize})
		// skip blank line
		scanner.Scan()
	}
	return machines
}

func main() {
	machines := parseInput()
	// for _, m := range machines {
	// 	fmt.Printf("%+v\n", m)
	// 	solve(m)
	// }
	part1(machines)
	part2(machines)
}

func part1(machines []Machine) {
	total := 0
	for _, m := range machines {
		total += solve(m, 0)
	}
	fmt.Println("total:", total)
}

func part2(machines []Machine) {
	fmt.Println("Part 2:")
	total := 0
	for _, m := range machines {
		total += solve(m, 10000000000000)
	}
	fmt.Println("total:", total)
}

func div(a, b int) (q int, ok bool) {
	ok = a%b == 0
	q = a / b
	return
}

func solve(m Machine, extra int) int {
	det := m.a.x*m.b.y - m.a.y*m.b.x
	if det == 0 {
		panic("zero determinant!")
	} else {
		fmt.Println("determinant", det)
		x := m.prize.x + extra
		y := m.prize.y + extra
		a, a_ok := div(m.b.y*x-m.b.x*y, det)
		b, b_ok := div(-m.a.y*x+m.a.x*y, det)
		if a_ok && b_ok {
			fmt.Println("a:", a, "b:", b)
			return a*3 + b
		} else {
			fmt.Println("no solution")
			return 0
		}
	}
}
