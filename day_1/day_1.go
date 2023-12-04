package day_1

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Part_two() int {
    num_map := map[string]rune{"one": '1', "two": '2', "three": '3', "four": '4', "five": '5',
        "six": '6', "seven": '7', "eight": '8', "nine": '9'}
    file, error := os.Open("input.txt")
	if error != nil {
		log.Fatal(error)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
    numbers := []int{}
    for scanner.Scan() {
        nums := []rune{}
        line := scanner.Text()
        if line == "" {
            break
        }
        for i, character := range line {
            if character >= '0' && character <= '9' {
                nums = append(nums, character)
            }
            for key, value := range num_map {
                if strings.HasPrefix(line[i:], key) {
                    nums = append(nums, value)
                }
            }
        }
        first := int(nums[0] - '0') * 10
        second := int(nums[len(nums) - 1] - '0')
        numbers = append(numbers, first + second)
    }
    total := 0
    for _, num := range numbers {
        total += num
    }
	return total
}

func Part_one() int {
    file, error := os.Open("input.txt")
	if error != nil {
		log.Fatal(error)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	numbers_arr := []int{}
	reg, error := regexp.Compile("[^0-9]+")
	if error != nil {
		log.Fatal(error)
	}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		nums := reg.ReplaceAllString(line, "")
		num_string := strings.Split(nums, "")
		memes := num_string[0]
		memes += num_string[len(num_string)-1]
		num, error := strconv.Atoi(memes)
		if error != nil {
			log.Fatal(error)
		}
		numbers_arr = append(numbers_arr, num)
	}

	total := 0
	for _, v := range numbers_arr {
		total += v
	}
	return total
}
