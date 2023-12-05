package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func PartTwo() int {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	powers := []int{}

	for scanner.Scan() {
		red := 0
		green := 0
		blue := 0
		current_val := 0
		words := strings.Split(scanner.Text(), " ")
		words = words[2:]
		for i, word := range words {
			if i%2 == 0 {
				val, err := strconv.Atoi(word)
				if err != nil {
					log.Fatal("Broken file format")
				}
				current_val = val
			} else {
				word = word[:1]
				if word == "g" {
					if current_val > green {
						green = current_val
					}
				} else if word == "r" {
					if current_val > red {
						red = current_val
					}
				} else {
					if current_val > blue {
						blue = current_val
					}
				}
			}
		}
		powers = append(powers, red*green*blue)

	}
	total := 0
	for _, power := range powers {
		total += power
	}

	return total
}

func PartOne() int {
	red := 12
	green := 13
	blue := 14
	file, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	ids := []int{}

	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		id_s := words[1]
		id_s = id_s[:(len(id_s) - 1)]
		id, err := strconv.Atoi(id_s)
		if err != nil {
			log.Fatal(err)
		}
		words = words[2:]
		current_val := 0
		possible := true
		for i, word := range words {
			if i%2 != 0 {
				word := word[:1]
				fmt.Println(word)
				if word == "g" && current_val > green {
					fmt.Printf("%v: %v\n", word, current_val)
					possible = false
					break
				} else if word == "r" && current_val > red {
					fmt.Printf("%v: %v\n", word, current_val)
					possible = false
					break
				} else if word == "b" && current_val > blue {
					fmt.Printf("%v: %v\n", word, current_val)
					possible = false
					break
				}
			} else {
				val, err := strconv.Atoi(word)
				if err != nil {
					log.Fatal("Broken file format")
				}
				current_val = val
			}
		}
		if possible {
			ids = append(ids, id)
		}
	}
	total := 0
	for _, id := range ids {
		total += id
	}
	return total
}
