package main

import (
	"github.com/swsd2544/AdventOfCode2022/reader"
	"log"
	"strconv"
	"strings"
)

func convertInputToElves(text string) ([]int, error) {
	elves := make([]int, 0)
	for _, elf := range strings.Split(text, "\n\n") {
		totalCalories := 0
		for _, caloriesText := range strings.Split(elf, "\n") {
			calories, err := strconv.Atoi(caloriesText)
			if err != nil {
				return nil, err
			}
			totalCalories += calories
		}
		elves = append(elves, totalCalories)
	}

	return elves, nil
}

func findTopThreeHighestCalories(elves []int) [3]int {
	var topThree [3]int
	for i := 0; i < 3; i++ {
		var highestCalories int
		for _, calories := range elves {
			if calories > highestCalories && (i == 0 || calories < topThree[i-1]) {
				highestCalories = calories
			}
		}
		topThree[i] = highestCalories
	}
	return topThree
}

func main() {
	text, err := reader.GetTextFromInputFile("./day1/input.text")
	if err != nil {
		log.Fatalf("can't read text: %v", err)
	}

	elves, err := convertInputToElves(text)
	if err != nil {
		log.Fatalf("input file isn't in the correct format: %v", err)
	}

	topThree := findTopThreeHighestCalories(elves)
	log.Printf("Elf calories rank:\n1. %d\n2. %d\n3. %d", topThree[0], topThree[1], topThree[2])
	log.Printf("Total caloires: %d\n", topThree[0]+topThree[1]+topThree[2])
}
