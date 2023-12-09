package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func checkSides(row []rune, column int) bool {
	left := row[column-1]
	right := row[column+3]
	if (right >= '0' && right <= '9') || right == '.' {
		if (left >= '0' && left <= '9') || left == '.' {
			return true
		}
	}
	return false
}

func checkRow(row []rune, start int, end int) bool {
	for index := start; index < end+1; index++ {
		curr := row[index]
		if !(curr >= '0' && curr <= '9') && curr != '.' {
			return false
		}
	}
	return true
}

// I give up on this, the solution is in sight but... this challenge day is so unfun
func PartOne() int {
	file, err := os.Open("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	rows := []string{}
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}
	rsize := len(rows[0]) - 3

	nums := []int{}
	skip := 0
	for r, row := range rows {
		for c, col := range row {
			if skip > 0 {
				skip -= 1
			}
			if col >= '0' && col <= '9' && skip == 0 {
				if r > 0 && r < rsize {
					if checkRow([]rune(rows[r-1]), c-1, c+3) && checkRow([]rune(rows[r+1]), c-1, c-3) && checkSides([]rune(rows[r]), c) {
						num, err := strconv.Atoi(string(row[c]) + string(row[c+1]) + string(row[c+2]))
						if err != nil {
							log.Fatal("Failed to convert runes to number")
						}
						nums = append(nums, num)
					}
				} else if r > 0 {

				} else {

				}

				skip = 3
			}
			fmt.Printf("%val\n", c)
		}
	}

	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
