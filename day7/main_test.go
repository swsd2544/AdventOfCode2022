package main

import (
	"testing"
)

const text = `
	$ cd /
	$ ls
	dir a
	14848514 b.txt
	8504156 c.dat
	dir d
	$ cd a
	$ ls
	dir e
	29116 f
	2557 g
	62596 h.lst
	$ cd e
	$ ls
	584 i
	$ cd ..
	$ cd ..
	$ cd d
	$ ls
	4060174 j
	8033020 d.log
	5626152 d.ext
	7214296 k
	`

func TestSolveI(t *testing.T) {
	var want uint64 = 95_437
	totalSize := solveI(text)
	if totalSize != want {
		t.Errorf("want %d, got %d", want, totalSize)
	}
}

func TestSolveII(t *testing.T) {
	var want uint64 = 24933642
	totalSize := solveII(text)
	if totalSize != want {
		t.Errorf("want %d, got %d", want, totalSize)
	}
}
