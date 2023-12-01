package solution

import (
	"strconv"
	"strings"
	"unicode"
)

func Part1(data string) int {
	lines := strings.Split(data, "\n")

	total := 0

	for _, line := range lines {
		total += getCalibrationValue(line)
	}
	return total
}

func getCalibrationValue(line string) int {
	var first, last byte = '#', '#'
	for i := 0; i < len(line); i++ {
		if unicode.IsDigit(rune(line[i])) {
			first = line[i]
			break
		}
	}
	for j := len(line) - 1; j >= 0; j-- {
		if unicode.IsDigit(rune(line[j])) {
			last = line[j]
			break
		}
	}
	if first == '#' {
		return 0
	}
	digit, _ := strconv.Atoi(string([]byte{first, last}))
	return digit
}
