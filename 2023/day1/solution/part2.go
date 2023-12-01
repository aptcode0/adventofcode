package solution

import (
	"strconv"
	"strings"
	"unicode"
)

var digits = []string{"one", "two", "three", "four", "five",
	"six", "seven", "eight", "nine"}
var letterToDigits = digitsMap()

func Part2(data string) int {
	lines := strings.Split(data, "\n")

	total := 0

	for _, line := range lines {
		total += getNewCalibrationValue(line)
	}
	return total
}

func digitsMap() map[string]int {
	m := map[string]int{}
	for i, d := range digits {
		m[d] = i + 1
	}
	return m
}
func getNewCalibrationValue(line string) int {
	var first, last byte = '#', '#'

	for i := 0; i < len(line); i++ {
		isDigit := unicode.IsDigit(rune(line[i]))
		var curr byte = '#'
		if isDigit {
			curr = line[i]

		} else if digit := findLetterDigit(line, i); digit != -1 {
			curr = byte(digit + '0')
		}
		if curr == '#' {
			continue
		}
		if first == '#' {
			first = curr
		}
		last = curr
	}

	if first == '#' {
		return 0
	}
	digit, _ := strconv.Atoi(string([]byte{first, last}))
	return digit
}

func findLetterDigit(s string, start int) int {
	for l, d := range letterToDigits {
		end := start + len(l)
		if end > len(s) {
			continue
		}

		if s[start:end] == l {
			return d
		}
	}
	return -1
}
