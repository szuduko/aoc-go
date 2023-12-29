package util

import (
	"os"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadInput(filepath string) []string {
	data, err := os.ReadFile(filepath)
	Check(err)

	lines := strings.Split(string(data), "\n")

	return lines
}