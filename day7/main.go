package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	total    int
	operands []int
}

func (e Equation) String() string {
	return fmt.Sprintf("%d: %v", e.total, e.operands)
}

func _int(x string) int {
	i, _ := strconv.Atoi(x)
	return i
}

func parseInput() []Equation {
	scanner := bufio.NewScanner(os.Stdin)
	equations := make([]Equation, 0)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.FieldsFunc(line, func(e rune) bool { return e == ' ' || e == ':' })
		operands := make([]int, len(fields)-1)
		for i, x := range fields[1:] {
			operands[i] = _int(x)
		}
		equations = append(equations, Equation{_int(fields[0]), operands})
	}
	return equations
}

func check(out, cur int, rest []int, ops []func(int, int) int) bool {
	if cur > out {
		return false
	} else if len(rest) > 0 {
		next, rest := rest[0], rest[1:]
		for _, op := range ops {
			if check(out, op(cur, next), rest, ops) {
				return true
			}
		}
		return false
	} else {
		return out == cur
	}
}

func plus(a, b int) int {
	return a + b
}

func times(a, b int) int {
	return a * b
}

func concat(a, b int) int {
	return _int(fmt.Sprintf("%d%d", a, b))
}

func part1(input []Equation) int {
	total := 0
	ops := []func(int, int) int{plus, times}
	for _, eq := range input {
		fmt.Println("Processing equation: ", eq)
		if check(eq.total, eq.operands[0], eq.operands[1:], ops) {
			fmt.Println("  good!")
			total += eq.total
		}
	}
	return total
}

func part2(input []Equation) int {
	total := 0
	ops := []func(int, int) int{plus, times, concat}
	for _, eq := range input {
		fmt.Println("Processing equation:", eq)
		if check(eq.total, eq.operands[0], eq.operands[1:], ops) {
			fmt.Println("  good!")
			total += eq.total
		}
	}
	return total
}

func main() {
	input := parseInput()
	fmt.Println("total:", part1(input))
	fmt.Println("total:", part2(input))
}
