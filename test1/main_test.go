package main

import "testing"

func TestSum(tt *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{5, 5, 10},
	}
	for _, t := range tables {
		total := Sum(t.x, t.y)
		if total != t.n {
			tt.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", t.x, t.y, total, t.n)
		}
	}
}
