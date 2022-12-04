package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./day4/input.text")
	if err != nil {
		log.Fatal(err)
	}

	buf := new(strings.Builder)
	_, err = io.Copy(buf, file)
	if err != nil {
		log.Fatal(err)
	}

	var fullyContained int
	textArr := strings.Split(buf.String(), "\n")
	for _, text := range textArr {
		elves := strings.Split(text, ",")
		elf1, elf2 := []int{}, []int{}
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

		if elf1[0] <= elf2[0] && elf1[1] >= elf2[1] {
			fullyContained++
		} else if elf2[0] <= elf1[0] && elf2[1] >= elf1[1] {
			fullyContained++
		}
	}

	log.Println(fullyContained)
}
