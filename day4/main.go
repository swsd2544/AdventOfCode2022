package main

import (
	"github.com/swsd2544/AdventOfCode2022/reader"
	"log"
	"strconv"
	"strings"
)

func main() {
	text, err := reader.GetTextFromInputFile("./day4/input.text")
	if err != nil {
		log.Fatalf("can't read text: %v", err)
	}

	var overlapped int
	textArr := strings.Split(text, "\n")
	for _, text := range textArr {
		elves := strings.Split(text, ",")
		var elf1 []int
		var elf2 []int
		for _, numStr := range strings.Split(elves[0], "-") {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}
			elf1 = append(elf1, num)
		}
		for _, numStr := range strings.Split(elves[1], "-") {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}
			elf2 = append(elf2, num)
		}

		if (elf1[0] >= elf2[0] && elf1[0] <= elf2[1]) || (elf1[1] >= elf2[0] && elf1[1] <= elf2[1]) {
			overlapped++
		} else if (elf2[0] >= elf1[0] && elf2[0] <= elf1[1]) || (elf2[1] >= elf1[0] && elf2[1] <= elf1[1]) {
			overlapped++
		}
	}

	log.Println(overlapped)
}
