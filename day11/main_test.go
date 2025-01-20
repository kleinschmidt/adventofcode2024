package main

import (
	"strconv"
	"testing"
)

func TestDigits(t *testing.T) {
	for i := range 1000 {
		lenStr := len(strconv.Itoa(i))
		if lenStr != digits(i) {
			t.Errorf("i=%v, got %v, want %v", i, digits(i), lenStr)
		}
	}
}
