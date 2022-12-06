package main

import (
	"github.com/swsd2544/AdventOfCode2022/reader"
	"log"
	"strings"
)

func getPriorities(char rune) int {
	if char >= 'a' && char <= 'z' {
		return int(char-'a') + 1
	} else if char >= 'A' && char <= 'Z' {
		return int(char-'A') + 27
	}
	return -1
}

func findSharedItemInPriorities(text, text2, text3 string) int {
	sharedItem := map[rune]bool{}
	for _, char := range text {
		if strings.ContainsRune(text2, char) && strings.ContainsRune(text3, char) {
			sharedItem[char] = true
		}
	}

	var priorities int
	for k := range sharedItem {
		priorities += getPriorities(k)
	}

	return priorities
}

func rucksackReorganization(text string) (int, error) {
	var totalPriorities int
	stringArray := strings.Split(text, "\n")
	for i := 0; i <= len(stringArray)-3; i += 3 {
		totalPriorities += findSharedItemInPriorities(stringArray[i], stringArray[i+1], stringArray[i+2])
	}

	return totalPriorities, nil
}

func main() {
	text, err := reader.GetTextFromInputFile("input.text")
	if err != nil {
		log.Fatalf("can't read text: %v", err)
	}

	totalPriorities, err := rucksackReorganization(text)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(totalPriorities)
}
