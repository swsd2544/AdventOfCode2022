package main

import (
	"fmt"
	"github.com/swsd2544/AdventOfCode2022/reader"
	"log"
	"strconv"
	"strings"
)

const (
	CD = "cd"
	LS = "ls"
)

var root *Directory

type Directory struct {
	parent   *Directory
	name     string
	children []any
}

type File struct {
	name string
	size uint64
}

func BuildDirectory(text string) (*Directory, error) {
	root = &Directory{
		name: "/",
	}
	curr := root
	for _, line := range strings.Split(text, "\n")[1:] {
		line := strings.TrimSpace(line)
		if IsCommand(line) {
			commandType, err := CheckCommandType(line)
			if err != nil {
				return nil, err
			}
			if commandType == CD {
				curr, err = ChangeDirectory(curr, line)
				if err != nil {
					return nil, err
				}
			}
		} else {
			err := curr.AddChild(line)
			if err != nil {
				return nil, err
			}
			continue
		}
	}
	return root, nil
}

func (d *Directory) AddChild(child string) error {
	info := strings.Split(child, " ")
	if strings.HasPrefix(info[0], "dir") {
		d.children = append(d.children, &Directory{name: info[1], parent: d})
		return nil
	}

	size, err := strconv.ParseUint(info[0], 10, 64)
	if err != nil {
		return err
	}
	d.children = append(d.children, &File{name: info[1], size: size})
	return nil
}

func (d *Directory) Size() uint64 {
	var size uint64
	queue := []*Directory{d}
	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, child := range curr.children {
			switch child.(type) {
			case *Directory:
				queue = append(queue, child.(*Directory))
			case *File:
				size += child.(*File).size
			}
		}
	}
	return size
}

func ChangeDirectory(d *Directory, CDCommand string) (*Directory, error) {
	target := strings.Split(CDCommand, " ")[2]
	if target == "/" {
		return root, nil
	}
	if target == ".." {
		return d.parent, nil
	}
	for _, child := range d.children {
		if child, ok := child.(*Directory); ok && child.name == target {
			return child, nil
		}
	}
	return nil, fmt.Errorf("target not found: %s", target)
}

func IsCommand(text string) bool {
	return strings.HasPrefix(text, "$")
}

func CheckCommandType(text string) (string, error) {
	if strings.Contains(text, CD) {
		return CD, nil
	}
	if strings.Contains(text, LS) {
		return LS, nil
	}
	return "", fmt.Errorf("command not found: %s", text)
}

func solveI(text string) uint64 {
	var totalSize uint64
	d, err := BuildDirectory(text)
	if err != nil {
		log.Fatal(err)
	}
	queue := []*Directory{d}
	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, child := range curr.children {
			switch child.(type) {
			case *Directory:
				queue = append(queue, child.(*Directory))
			}
		}
		if size := curr.Size(); size < 100_000 {
			totalSize += size
		}
	}
	return totalSize
}

func solveII(text string) uint64 {
	d, err := BuildDirectory(text)
	if err != nil {
		log.Fatal(err)
	}
	var minSize, capSize uint64
	queue := []*Directory{d}
	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, child := range curr.children {
			switch child.(type) {
			case *Directory:
				queue = append(queue, child.(*Directory))
			}
		}
		size := curr.Size()
		if curr == root {
			availableSpace := 70_000_000 - size
			if availableSpace >= 30_000_000 {
				return 0
			}
			capSize = 30_000_000 - availableSpace
			minSize = size
		}
		if size >= capSize && size < minSize {
			minSize = size
		}
	}
	return minSize
}

func main() {
	text, err := reader.GetTextFromInputFile("./day7/input.text")
	if err != nil {
		log.Fatalf("can't read text: %v", err)
	}

	log.Println(solveI(text))
	log.Println(solveII(text))
}
