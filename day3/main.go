package main

import (
	"io"
	"log"
	"os"
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

func rucksackReorganization(r io.Reader) (int, error) {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, r)
	if err != nil {
		return 0, err
	}

	var totalPriorities int
	strarr := strings.Split(buf.String(), "\n")
	for i := 0; i <= len(strarr)-3; i += 3 {
		totalPriorities += findSharedItemInPriorities(strarr[i], strarr[i+1], strarr[i+2])
	}

	return totalPriorities, nil
}

func main() {
	file, err := os.Open("./day3/input.text")
	if err != nil {
		log.Fatal(err)
	}
	totalPriorities, err := rucksackReorganization(file)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(totalPriorities)
}
