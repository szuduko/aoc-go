package day03

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/szuduko/aoc-go/util"
)

type Number struct {
	value int
	row int
	col []int
}

type Coordinate struct {
	row int
	col int
}

// symbol map instead?? make(map[rune][]Coordinate)
type Symbol struct {
	value rune
	row int
	col int
}

func isDigit(character rune) bool {
	if character == '0' || character == '1' || character == '2' || character == '3' || character == '4' || character == '5' || character == '6' || character == '7' || character == '8' || character == '9' {
		return true
	}
	return false
}

func isSymbol(character rune) bool {
	// @ # $ % & * - + = /
	if character == '@' || character == '#' || character == '$' || character == '%' || character == '&' || character == '*' || character == '-' || character == '+' || character == '=' || character == '/' {
		return true
	}
	return false
}

func Part1() {
	fmt.Println("Part 1")
	lines := util.ReadInput("day03/day03-input.txt")

	numbers := make([]Number, 0)
	symbols := make([]Symbol, 0)

	for row, line := range lines {
		line = strings.TrimRight(line, "\r\n")
		line += "."
		fmt.Println("Line:", line)

		number := Number{value: 0, row: row}
		tempNumber := ""
		for col, char := range line {
			fmt.Printf("%c", char)
			fmt.Printf(" (%d,%d)", row, col)
			fmt.Println()
		
			if isDigit(char) {
				number.col = append(number.col, col)
				tempNumber += string(char)
			} else {
				if len(tempNumber) > 0 {
					value, err := strconv.Atoi(tempNumber)
					util.Check(err)
					tempNumber = ""

					number.value = value
					fmt.Println("Number:", number)
					numbers = append(numbers, number)
					number = Number{value: 0, row: row}
				}
				if isSymbol(char) {
					symbol := Symbol{value: char, row: row, col: col}
					fmt.Println("Symbol:", symbol)
					symbols = append(symbols, symbol)
				}
			}
		}
		fmt.Println()
	}

	sum := 0
	for _, symbol := range symbols {
		fmt.Println("Symbol:", symbol)
		for _, number := range numbers {
			for _, index := range number.col {
				if symbol.row == number.row - 1 || symbol.row == number.row || symbol.row == number.row + 1 {
					if symbol.col == index -1 || symbol.col == index || symbol.col == index + 1 {
						sum += number.value
						break
					}
				}
			}
		}
	}
	fmt.Println("Sum:", sum)
}

func Part2() {
	fmt.Println("Part 2")
}