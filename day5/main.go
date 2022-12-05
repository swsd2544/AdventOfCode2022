package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func getCrates(text string) map[int][]string {
	crates := map[int][]string{}

	textSplit := strings.Split(text, "\n")
	for i := len(textSplit) - 2; i >= 0; i-- {
		for j := 3; j <= len(textSplit[i]); j += 4 {
			crate := strings.TrimSpace(textSplit[i][j-3 : j])
			if crate != "" {
				crates[j/4] = append(crates[j/4], crate)
			}
		}
	}

	return crates
}

func takeAction(crates map[int][]string, procedure string) error {
	procedureArr := strings.Split(procedure, " ")
	n, err := strconv.Atoi(string(procedureArr[1]))
	if err != nil {
		return err
	}

	from, err := strconv.Atoi(string(procedureArr[3]))
	if err != nil {
		return err
	}

	to, err := strconv.Atoi(string(procedureArr[5]))
	if err != nil {
		return err
	}

	from--
	to--

	crates[to] = append(crates[to], crates[from][len(crates[from])-n:]...)
	crates[from] = crates[from][:len(crates[from])-n]

	return nil
}

func main() {
	buf, err := ioutil.ReadFile("./day5/input.text")
	if err != nil {
		log.Fatal(err)
	}

	text := strings.Split(string(buf), "\n\n")
	crates := getCrates(text[0])
	for _, procedure := range strings.Split(text[1], "\n") {
		err := takeAction(crates, procedure)
		if err != nil {
			log.Fatal(err)
		}
	}

	for i := 0; i < len(crates); i++ {
		crate := string(crates[i][len(crates[i])-1][1])
		fmt.Print(crate)
	}

	fmt.Print("\n")
}
