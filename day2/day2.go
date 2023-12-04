package day2

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
    "fmt"
)

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
            if i % 2 != 0 {
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
