package points

import "iter"

type Point struct{ X, Y int }

func Add(a, b Point) Point {
	return Point{a.X + b.X, a.Y + b.Y}
}

func Clockwise(p Point) Point {
	return Point{p.Y, -p.X}
}

func Neighbors(p Point) iter.Seq2[Point, Point] {
	directions := [4]Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	return func(yield func(Point, Point) bool) {
		for _, d := range directions {
			pn := Add(p, d)
			if !yield(d, pn) {
				return
			}
		}
		return
	}
}

func Inbounds[E any](p Point, board [][]E) bool {
	return p.X >= 0 && p.X < len(board[0]) && p.Y >= 0 && p.Y < len(board)
}

func Get[E any](board [][]E, p Point) *E {
	if Inbounds(p, board) {
		return &board[p.Y][p.X]
	} else {
		// use 0 as a sentinel value
		return nil
	}
}
