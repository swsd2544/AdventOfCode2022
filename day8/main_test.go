package main

import (
	"testing"
)

const text = `30373
25512
65332
33549
35390`

func TestSolveI(t *testing.T) {
	want := 21
	visibleTrees := solveI(text)
	if visibleTrees != want {
		t.Errorf("want %d, got %d", want, visibleTrees)
	}
}

func TestSolveII(t *testing.T) {
	want := 8
	highestScenicScore := solveII(text)
	if highestScenicScore != want {
		t.Errorf("want %d, got %d", want, highestScenicScore)
	}
}
