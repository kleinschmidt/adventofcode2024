package main

import "testing"

func TestPmod(t *testing.T) {
	in := []int{-3, -2, -1, 0, 1, 2, 3}
	out := []int{0, 1, 2, 0, 1, 2, 0}
	for i := range in {
		got := pmod(in[i], 3)
		if got != out[i] {
			t.Errorf("pmod(%d): expected %d, got %d", in[i], out[i], got)
		}
	}
}

func TestQuadrant(t *testing.T) {
	size := Point{5, 7}
	var qTrue Point
	for x := range size.x {
		for y := range size.y {
			p := Point{x, y}
			q := quadrant(p, size)
			if x == 2 || y == 3 {
				qTrue = Point{0, 0}
			} else {
				var qx, qy int
				if x < 2 {
					qx = -1
				} else {
					qx = 1
				}
				if y < 3 {
					qy = -1
				} else {
					qy = 1
				}
				qTrue = Point{qx, qy}
			}
			if q != qTrue {
				t.Errorf("quadrant(%v, %v): expected %v, got %v", p, size, qTrue, q)
			}
		}
	}
}
