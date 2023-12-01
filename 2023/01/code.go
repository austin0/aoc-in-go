package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	var sum int64 = 0

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		numString := searchForNumbers(line)

		numInt, err := strconv.ParseInt(numString, 10, 64)
		if err != nil {
			fmt.Println(err)
			return err
		}

		sum += numInt
	}

	return sum
}

func searchForNumbers(line string) string {
	// search string from both ends inwards until first number is found
	leftIndex, leftString, rightIndex, rightString := 1, "", len(line)-1, ""
	for leftString == "" || rightString == "" {
		leftString, rightString = parseString(line[:leftIndex]), parseString(line[rightIndex:])
		if leftString == "" {
			leftIndex++
		}
		if rightString == "" {
			rightIndex--
		}
	}
	return leftString + rightString
}

func parseString(stringSlice string) string {
	re := regexp.MustCompile("[0-9]{1}|(one|two|three|four|five|six|seven|eight|nine)")
	numSlice := re.FindAllString(stringSlice, -1)
	if numSlice != nil {
		return numberToDigit(numSlice[0])
	}
	return ""
}

func numberToDigit(number string) string {
	switch number {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return number
	}
}
