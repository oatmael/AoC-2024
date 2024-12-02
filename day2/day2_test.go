package main

import "testing"

func TestPart1(t *testing.T) {
	data := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	};

	if answer := processSafeCount(data); answer != 2 {
		t.Fatalf(`Got %d, expected 2`, answer);
	}
}