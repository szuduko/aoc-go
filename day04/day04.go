package day04

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/szuduko/aoc-go/util"
)

type card struct {
	number int
	winningNumbers []int
}

func collectCards(lines []string) []card {
	cards := make([]card, 0)

	for _, line := range lines {
		line = strings.ToLower(line)
		temp := strings.Split(line, ": ")
		
		rawCardNumber := strings.ReplaceAll(temp[0], " ", "")
		rawCardNumber = strings.TrimLeft(rawCardNumber, "card")

		cardNumber, err := strconv.Atoi(rawCardNumber)
		util.Check(err)

		fmt.Println(cardNumber)

		numbers := strings.Split(temp[1], " | ")

		winning := sanitizeNumber(numbers[0])

		owned := sanitizeNumber(numbers[1])

		winningNumbers := collectWinningNumbers(winning, owned)

		card := card{number: cardNumber, winningNumbers: winningNumbers}

		cards = append(cards, card)
	}

	return cards
}

func sanitizeNumber(numbers string) []int {
	numbers = strings.TrimRight(numbers, "\r\n ")
	numbers = strings.TrimLeft(numbers, " ")

	sanitisedNumbers := make([]int, 0)

	for _, number := range strings.Split(numbers, " ") {
		if number != "" {
			number, err := strconv.Atoi(number)
			util.Check(err)

			sanitisedNumbers = append(sanitisedNumbers, number)
		}
	}

	fmt.Println("SanitisedNumbers:", sanitisedNumbers)

	return sanitisedNumbers
}

func collectWinningNumbers(winning []int, owned []int) []int {
	winningNumbers := make([]int, 0)
	for _, number := range owned {
		if slices.Contains(winning, number) {
			winningNumbers = append(winningNumbers, number)
		}
	}
	return winningNumbers
}

func Part1() {
	fmt.Println("Part 1")
	lines := util.ReadInput("day04/day04-input.txt")

	cards := collectCards(lines)

	points := 0.0
	for _, card := range cards {
		fmt.Println(card)
		
		if len(card.winningNumbers) > 1 {
			points += math.Pow(2, float64(len(card.winningNumbers) - 1))
		}
		if len(card.winningNumbers) == 1 {
			points += 1
		}
	}
	fmt.Println("Points:", points)
}

func Part2() {
	fmt.Println("Part 2")
}