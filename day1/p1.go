package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	getTotalCalibration()
}

func getTotalCalibration() {
	file, ok := os.ReadFile("./day1p1")
	checkError(ok)
	total := 0
	for _, line := range strings.Split(string(file), "\n") {
		digit := countCalibrationLine(line)
		total += digit
	}
	fmt.Printf("%d\n", total)
}

func checkError(ok error) {
	if ok != nil {
		panic(ok)
	}
}

func countCalibrationLine(line string) int {
	first := getFirstDigit(line)
	last := getLastDigit(line)
	fmt.Printf("%d%d\n", first, last)
	return first*10 + last
}

func getFirstDigit(line string) int {
	len := len(line) - 1
	for i := range line {
		if isDigit(line[i]) {
			return int(line[i] - '0')
		}
		if digit := isDigitString(line[i:min(len, i+6)]); digit != -1 {
			return digit
		}

	}
	return 0
}

func getLastDigit(line string) int {
	len := len(line) - 1
	for i := range line {
		if isDigit(line[len-i]) {
			return int(line[len-i] - '0')
		}

		if digit := isDigitString(line[len-i : len+1]); digit != -1 {
			return digit
		}
	}
	return 0
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isDigitString(line string) int {
	switch {
	case strings.HasPrefix(line, "one"):
		return 1
	case strings.HasPrefix(line, "two"):
		return 2
	case strings.HasPrefix(line, "three"):
		return 3
	case strings.HasPrefix(line, "four"):
		return 4
	case strings.HasPrefix(line, "five"):
		return 5
	case strings.HasPrefix(line, "six"):
		return 6
	case strings.HasPrefix(line, "seven"):
		return 7
	case strings.HasPrefix(line, "eight"):
		return 8
	case strings.HasPrefix(line, "nine"):
		return 9
	default:
		return -1
	}
}
