package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

var numberRegex = regexp.MustCompile("[0-9]+")
var cubesRegex = regexp.MustCompile("([0-9]+ (red|green|blue)){1}")
var colourRegex = regexp.MustCompile("(red|green|blue){1}")
var gameRegex = regexp.MustCompile("([0-9]+:){1}")

var cubesLimits = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

var sum int

func run(part2 bool, input string) any {
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		possible := true

		gameNumberString := gameRegex.FindString(line)
		gameNumberString = strings.Replace(gameNumberString, ":", "", 1)

		gameNumberInt, err := strconv.Atoi(gameNumberString)
		if err != nil {
			return err
		}

		cubesSlice := cubesRegex.FindAllString(line, -1)
		for _, cube := range cubesSlice {
			cubeColour := colourRegex.FindString(cube)
			cubeNumberString := numberRegex.FindString(cube)
			cubeNumberInt, err := strconv.Atoi(cubeNumberString)
			if err != nil {
				return err
			}
			if cubeNumberInt > cubesLimits[cubeColour] {
				possible = false
			}
		}
		if possible {
			sum += gameNumberInt
		}
	}

	return sum
}
