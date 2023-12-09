package day4

import (
	"bufio"
	"log"
	"math"
	"os"
	"slices"
	"strings"
)

func PartOne() int {
	file, err := os.Open("day4/input.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := 0
		line := scanner.Text()[10:]
		vals := strings.Fields(line)
		have := vals[:10]
		winning := vals[11:]
		for _, num := range have {
			if slices.Contains(winning, num) {
				matches += 1
			}
		}
		total += int(math.Pow(2, float64(matches-1)))
	}
	return total
}
