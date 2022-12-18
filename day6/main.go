package main

import (
	"github.com/swsd2544/AdventOfCode2022/reader"
	"log"
	"strings"
)

func AllCharactersUnique(text string) bool {
	hashset := map[rune]bool{}
	for _, char := range text {
		_, ok := hashset[char]
		if ok {
			return false
		}
		hashset[char] = true
	}
	return true
}

func solveI(text string) int {
	var marker string
	for i := 0; i < len(text); i++ {
		character := text[i : i+1]
		marker += character
		if !AllCharactersUnique(marker) {
			notUniqueIndex := strings.Index(marker, character)
			marker = marker[notUniqueIndex+1:]
		}
		if len(marker) < 4 {
			continue
		}
		return i + 1
	}
	return -1
}

func solveII(text string) int {
	var marker string
	for i := 0; i < len(text); i++ {
		character := text[i : i+1]
		marker += character
		if !AllCharactersUnique(marker) {
			notUniqueIndex := strings.Index(marker, character)
			marker = marker[notUniqueIndex+1:]
		}
		if len(marker) < 14 {
			continue
		}
		return i + 1
	}
	return -1
}

func main() {
	text, err := reader.GetTextFromInputFile("./day6/input.text")
	if err != nil {
		log.Fatalf("can't read text: %v", err)
	}

	log.Println(solveI(text))
	log.Println(solveII(text))
}
