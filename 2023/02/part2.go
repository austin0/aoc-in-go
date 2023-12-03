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

var sum int

func run(part2 bool, input string) any {
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		minimumCubesNeeded := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		cubesSlice := cubesRegex.FindAllString(line, -1)

		for _, cube := range cubesSlice {
			cubeColour := colourRegex.FindString(cube)
			cubeNumberString := numberRegex.FindString(cube)
			cubeNumberInt, err := strconv.Atoi(cubeNumberString)
			if err != nil {
				return err
			}
			if cubeNumberInt > minimumCubesNeeded[cubeColour] {
				minimumCubesNeeded[cubeColour] = cubeNumberInt
			}
		}

		gamePower := minimumCubesNeeded["red"] * minimumCubesNeeded["green"] * minimumCubesNeeded["blue"]
		sum += gamePower
	}

	return sum
}
