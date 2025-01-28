package main

import "testing"

func TestInbounds(t *testing.T) {
	corner := Point{3, 3}
	board := make([][]byte, corner.y)
	for y := range board {
		board[y] = make([]byte, corner.x)
	}
	shouldInbounds := []Point{{0, 0}, {2, 0}, {0, 2}, {2, 2}}
	for _, p := range shouldInbounds {
		if !inbounds(p, board) {
			t.Error("Expected inbounds: ", p)
		}
	}

	shouldOutbounds := []Point{{-1, 0}, {0, -1}, {-1, -1}, {2, 3}, {3, 2}, {3, 3}, {-1, 2}, {0, 3}, {2, -1}, {3, 0}}
	for _, p := range shouldOutbounds {
		if inbounds(p, board) {
			t.Error("Expected not inbounds: ", p)
		}
	}
}
