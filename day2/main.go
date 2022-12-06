package main

import (
	"fmt"
	"github.com/swsd2544/AdventOfCode2022/reader"
	"log"
	"strings"
)

func calculateRoundScore(elf, player string) int {
	var score int

	if player == "Y" {
		score = 3
	} else if player == "Z" {
		score = 6
	}

	switch elf {
	case "A":
		if player == "X" {
			score += 3
		} else if player == "Y" {
			score += 1
		} else if player == "Z" {
			score += 2
		}
	case "B":
		if player == "X" {
			score += 1
		} else if player == "Y" {
			score += 2
		} else if player == "Z" {
			score += 3
		}
	case "C":
		if player == "X" {
			score += 2
		} else if player == "Y" {
			score += 3
		} else if player == "Z" {
			score += 1
		}
	}

	return score
}

func calculateGameScore(text string) int {
	var score int
	for _, round := range strings.Split(text, "\n") {
		inputs := strings.Split(round, " ")
		elf, player := inputs[0], inputs[1]
		score += calculateRoundScore(elf, player)
	}
	return score
}

func main() {
	text, err := reader.GetTextFromInputFile("./day2/input.text")
	if err != nil {
		log.Fatalf("can't read text: %v", err)
	}

	fmt.Printf("Total score: %d\n", calculateGameScore(string(text)))
}
