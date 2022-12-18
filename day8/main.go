package main

import (
	"fmt"
	"github.com/swsd2544/AdventOfCode2022/reader"
	"log"
	"strings"
)

func IsTreeAtPosVisible(trees []string, row, col int) (bool, error) {
	if row <= 0 || row >= len(trees)-1 || col <= 0 || col >= len(trees[0])-1 {
		return false, fmt.Errorf("error position: %d, %d", row, col)
	}

	// top
	topVisible := true
	for i := 0; i < row; i++ {
		topVisible = topVisible && trees[row][col] > trees[i][col]
	}

	// left
	leftVisible := true
	for i := 0; i < col; i++ {
		leftVisible = leftVisible && trees[row][col] > trees[row][i]
	}

	// right
	rightVisible := true
	for i := len(trees[0]) - 1; i > col; i-- {
		rightVisible = rightVisible && trees[row][col] > trees[row][i]
	}

	// bottom
	bottomVisible := true
	for i := len(trees) - 1; i > row; i-- {
		bottomVisible = bottomVisible && trees[row][col] > trees[i][col]
	}

	return leftVisible || topVisible || rightVisible || bottomVisible, nil
}

func getScenicScoreAtPos(trees []string, row, col int) (int, error) {
	if row <= 0 || row >= len(trees)-1 || col <= 0 || col >= len(trees[0])-1 {
		return 0, fmt.Errorf("error position: %d, %d", row, col)
	}

	// top
	var topScenicScore int
	for i := row - 1; i >= 0; i-- {
		topScenicScore++
		if trees[row][col] <= trees[i][col] {
			break
		}
	}

	// left
	var leftScenicScore int
	for i := col - 1; i >= 0; i-- {
		leftScenicScore++
		if trees[row][col] <= trees[row][i] {
			break
		}
	}

	// right
	var rightScenicScore int
	for i := col + 1; i < len(trees[0]); i++ {
		rightScenicScore++
		if trees[row][col] <= trees[row][i] {
			break
		}
	}

	// bottom
	var bottomScenicScore int
	for i := row + 1; i < len(trees); i++ {
		bottomScenicScore++
		if trees[row][col] <= trees[i][col] {
			break
		}
	}

	return topScenicScore * leftScenicScore * rightScenicScore * bottomScenicScore, nil
}

func solveI(text string) int {
	lines := strings.Split(text, "\n")
	numberOfVisibleTrees := len(lines)*2 + len(lines[0])*2 - 4
	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[0])-1; j++ {
			isVisible, err := IsTreeAtPosVisible(lines, i, j)
			if err != nil {
				log.Fatalln(err)
			}
			if isVisible {
				numberOfVisibleTrees++
			}
		}
	}
	return numberOfVisibleTrees
}

func solveII(text string) int {
	lines := strings.Split(text, "\n")
	var highestScenicScore int
	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[0])-1; j++ {
			scenicScore, err := getScenicScoreAtPos(lines, i, j)
			if err != nil {
				log.Fatalln(err)
			}
			if scenicScore > highestScenicScore {
				highestScenicScore = scenicScore
			}
		}
	}
	return highestScenicScore
}

func main() {
	text, err := reader.GetTextFromInputFile("./day8/input.text")
	if err != nil {
		log.Fatalf("can't read text: %v", err)
	}

	log.Println(solveI(text))
	log.Println(solveII(text))
}
