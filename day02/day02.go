package day02

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/szuduko/aoc-go/util"
)

type Subset struct {
	red int
	green int
	blue int
}

type Game struct {
	id int
	subsets []Subset
}

const (
	MAX_RED = 12
	MAX_GREEN = 13
	MAX_BLUE = 14
)

func Part1() {
	fmt.Println("Part 1")
	lines := util.ReadInput("day02/day02-input.txt")

	// The Elf would first like to know which games would have been possible if the bag contained only 12 red cubes, 13 green cubes, and 14 blue cubes?

	games := make([]Game, 0, len(lines))
	for _, line := range lines {
		line = strings.ToLower(strings.ReplaceAll(line, " ", ""))
		fmt.Println("Line:", line)
		// Split Game number: by delimeter ": "
		subsets := strings.Split(line, ":")
		
		game_id, err := strconv.Atoi(strings.Replace(subsets[0], "game", "", -1))
		util.Check(err)
		fmt.Println("Game Number:", game_id)

		// Split subsets by delimeter "; "
		fmt.Println("Raw Subset:", subsets[1])
		split_subsets := strings.Split(subsets[1], ";")
		// Create subsets slice with size of len(split_subsets)
		subsets_struct := make([]Subset, 0, len(split_subsets))

		for i, subset := range split_subsets {
			fmt.Println("Subset Number:", i)
			fmt.Println("Subset:", subset)
			// Identify the colour values, find index of occurences "red" "green" blue" (order by lowest) then remove all alphabetic characters and spaces
			// Loop over numbers determine match based on red, green, blue
			idx_red := 10000 
			temp := strings.Index(subset, "red")
			if temp != -1 {
				idx_red  = temp
			}
			idx_green := 10000
			temp = strings.Index(subset, "green")
			if temp != -1 {
				idx_green = temp
			}
			idx_blue := 10000
			temp = strings.Index(subset, "blue")
			if temp != -1 {
				idx_blue = temp
			}

			fmt.Printf("idx_red: %d, idx_green: %d, idx_blue: %d\n", idx_red, idx_green, idx_blue)

			colour_values_map := make(map[string]int)
			colour_values_map["red"] = 0
			colour_values_map["green"] = 0
			colour_values_map["blue"] = 0
			first_key := ""
			second_key := ""
			third_key := ""

			if idx_red < idx_green && idx_red < idx_blue {
				first_key = "red"
				if idx_green < idx_blue {
					second_key = "green"
					third_key = "blue"
				} else  {
					second_key = "blue"
					third_key = "green"
				}
			} else if idx_green < idx_red && idx_green < idx_blue {
				first_key = "green"
				if idx_red < idx_blue {
					second_key = "red"
					third_key = "blue"
				} else  {
					second_key = "blue"
					third_key = "red"
				}
			} else {
				first_key = "blue"
				if idx_red < idx_green {
					second_key = "red"
					third_key = "green"
				} else  {
					second_key = "green"
					third_key = "red"
				}
			}

			fmt.Println("Order of Colours:", first_key, second_key, third_key)
			
			// Possibly a regex issue? no clue
			re := regexp.MustCompile("[a-zA-Z_'\r']+")

			// Blindly assigning values to the colour_values_map using the keys array for ordering
			subset_values := strings.Split(subset, ",")
			for j, value := range subset_values {
				number_value, err := strconv.Atoi(re.ReplaceAllString(value, ""))
				util.Check(err)
				if (j == 0) {
					colour_values_map[first_key] = number_value
				} else if (j == 1) {
					colour_values_map[second_key] = number_value
				} else {
					colour_values_map[third_key] = number_value
				}
				fmt.Println("Assign Colour Value:", colour_values_map)
			}

			fmt.Println("Complete Colour Values Map:", colour_values_map)
			subset_struct := Subset {
				red: colour_values_map["red"],
				green: colour_values_map["green"],
				blue: colour_values_map["blue"],
			}

			subsets_struct = append(subsets_struct, subset_struct)
		}
		fmt.Println("Subsets:", subsets_struct)
		
		game := Game {
			id: game_id,
			subsets: subsets_struct,
		}

		games = append(games, game)
	}
	fmt.Println(games)

	sum_ids := 0
	for _, game := range games {
		game_passed := true
		for _, subset := range game.subsets {
			if subset.red > MAX_RED || subset.green > MAX_GREEN || subset.blue > MAX_BLUE {
				game_passed = false
				break
			}
		}

		if game_passed {
			sum_ids += game.id
		}
	}
	fmt.Println("Sum of IDs:", sum_ids)
}

func Part2() {
	fmt.Println("Part 2")
	lines := util.ReadInput("day02/day02-input.txt")

	// The Elf would first like to know which games would have been possible if the bag contained only 12 red cubes, 13 green cubes, and 14 blue cubes?

	games := make([]Game, 0, len(lines))
	for _, line := range lines {
		line = strings.ToLower(strings.ReplaceAll(line, " ", ""))
		fmt.Println("Line:", line)
		// Split Game number: by delimeter ": "
		subsets := strings.Split(line, ":")
		
		game_id, err := strconv.Atoi(strings.Replace(subsets[0], "game", "", -1))
		util.Check(err)
		fmt.Println("Game Number:", game_id)

		// Split subsets by delimeter "; "
		fmt.Println("Raw Subset:", subsets[1])
		split_subsets := strings.Split(subsets[1], ";")
		// Create subsets slice with size of len(split_subsets)
		subsets_struct := make([]Subset, 0, len(split_subsets))

		for i, subset := range split_subsets {
			fmt.Println("Subset Number:", i)
			fmt.Println("Subset:", subset)
			// Identify the colour values, find index of occurences "red" "green" blue" (order by lowest) then remove all alphabetic characters and spaces
			// Loop over numbers determine match based on red, green, blue
			idx_red := 10000 
			temp := strings.Index(subset, "red")
			if temp != -1 {
				idx_red  = temp
			}
			idx_green := 10000
			temp = strings.Index(subset, "green")
			if temp != -1 {
				idx_green = temp
			}
			idx_blue := 10000
			temp = strings.Index(subset, "blue")
			if temp != -1 {
				idx_blue = temp
			}

			fmt.Printf("idx_red: %d, idx_green: %d, idx_blue: %d\n", idx_red, idx_green, idx_blue)

			colour_values_map := make(map[string]int)
			colour_values_map["red"] = 0
			colour_values_map["green"] = 0
			colour_values_map["blue"] = 0
			first_key := ""
			second_key := ""
			third_key := ""

			if idx_red < idx_green && idx_red < idx_blue {
				first_key = "red"
				if idx_green < idx_blue {
					second_key = "green"
					third_key = "blue"
				} else  {
					second_key = "blue"
					third_key = "green"
				}
			} else if idx_green < idx_red && idx_green < idx_blue {
				first_key = "green"
				if idx_red < idx_blue {
					second_key = "red"
					third_key = "blue"
				} else  {
					second_key = "blue"
					third_key = "red"
				}
			} else {
				first_key = "blue"
				if idx_red < idx_green {
					second_key = "red"
					third_key = "green"
				} else  {
					second_key = "green"
					third_key = "red"
				}
			}

			fmt.Println("Order of Colours:", first_key, second_key, third_key)
			
			// Possibly a regex issue? no clue
			re := regexp.MustCompile("[a-zA-Z_'\r']+")

			// Blindly assigning values to the colour_values_map using the keys array for ordering
			subset_values := strings.Split(subset, ",")
			for j, value := range subset_values {
				number_value, err := strconv.Atoi(re.ReplaceAllString(value, ""))
				util.Check(err)
				if (j == 0) {
					colour_values_map[first_key] = number_value
				} else if (j == 1) {
					colour_values_map[second_key] = number_value
				} else {
					colour_values_map[third_key] = number_value
				}
				fmt.Println("Assign Colour Value:", colour_values_map)
			}

			fmt.Println("Complete Colour Values Map:", colour_values_map)
			subset_struct := Subset {
				red: colour_values_map["red"],
				green: colour_values_map["green"],
				blue: colour_values_map["blue"],
			}

			subsets_struct = append(subsets_struct, subset_struct)
		}
		fmt.Println("Subsets:", subsets_struct)
		
		game := Game {
			id: game_id,
			subsets: subsets_struct,
		}

		games = append(games, game)
	}
	fmt.Println(games)

	games_power_sum := 0
	for _, game := range games {
		largest_red := 0
		largest_green := 0
		largest_blue := 0

		for _, subset := range game.subsets {
			if subset.red > largest_red {
				largest_red = subset.red
			}
			if subset.green > largest_green {
				largest_green = subset.green
			}
			if subset.blue > largest_blue {
				largest_blue = subset.blue
			}
		}

		games_power_sum += largest_red * largest_green * largest_blue
	}

	fmt.Println("Games Power Sum:", games_power_sum)
}
