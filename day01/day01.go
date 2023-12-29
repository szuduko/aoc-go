package day01

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/szuduko/aoc-go/util"
)

func Part1() {
	fmt.Println("Part 1")
	lines := util.ReadInput("day01/day01-input.txt")
	
	sum := 0
	re := regexp.MustCompile("[0-9]")

	for _, line := range lines {
		digits := re.FindAllString(line, -1)
		first := string(digits[0])
		last := string(digits[len(digits) - 1])

		number, err := strconv.Atoi(string(first) + string(last))
		
		util.Check(err)

		sum += number
	}

	fmt.Println(sum)
}

func Part2() {
	fmt.Println("Part 2")
	lines := util.ReadInput("day01/day01-input.txt")
	
	sum := 0

	for _, line := range lines {
		line = strings.ToLower(line)

		firstIndex := -1
		lastIndex := -1
		firstDigit := ""
		lastDigit := ""

		numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
		numberMap := make(map[string]string)

		numberMap["one"] = "1"
		numberMap["two"] = "2"
		numberMap["three"] = "3"
		numberMap["four"] = "4"
		numberMap["five"] = "5"
		numberMap["six"] = "6"
		numberMap["seven"] = "7"
		numberMap["eight"] = "8"
		numberMap["nine"] = "9"
		
		for _, number := range numbers {
			index1 := strings.Index(line, number)
			index2 := strings.LastIndex(line, number)
			
			if index1 != -1 {
				if firstIndex == -1 {
					firstIndex = index1
					firstDigit = number
				}
				if firstIndex >= index1 {
					firstIndex = index1
					firstDigit = number
				}
			}

			if index2 != -1 {
				if lastIndex == -1 {
					lastIndex = index2
					lastDigit = number
				}
				if lastIndex < index2 {
					lastIndex = index2
					lastDigit = number
				}
			}
		}

		if len(firstDigit) > 1 {
			firstDigit = numberMap[firstDigit]
		}

		if len(lastDigit) > 1 {
			lastDigit = numberMap[lastDigit]
		}

		fullNumber := firstDigit

		if lastDigit != "0" {
			fullNumber= firstDigit + lastDigit
		}

		number, err := strconv.Atoi(fullNumber)
		
		util.Check(err)
		
		sum += number
	}

	fmt.Println(sum)
}