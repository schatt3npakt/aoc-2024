// DAY 1 – https://adventofcode.com/2024/day/1

package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	byte, _ := os.ReadFile("./input.txt")

	leftCol, rightCol := parseList(byte)

	// Solution: Part 1
	printTotalDistance(leftCol, rightCol)

	// Solution: Part 2
	printSimilarityScore(leftCol, rightCol)
}

func parseList (file []byte) (leftCol []int, rightCol []int) {
	parsedFile := string(file)
	splitString := strings.Split(parsedFile, "\n")

	for _, v := range splitString {
		colValues := strings.Split(v, "   ")

		leftVal, _ := strconv.Atoi(colValues[0])
		leftCol = append(leftCol, leftVal)

		rightVal, _ := strconv.Atoi(colValues[1])
		rightCol = append(rightCol, rightVal)
	}

	slices.Sort(leftCol)
	slices.Sort(rightCol)

	return
}

func printTotalDistance (leftCol []int, rightCol []int) {
	var totalDiff int

	for i := range leftCol {
		totalDiff = totalDiff + calcDistance(leftCol[i], rightCol[i])
	}

	fmt.Printf("Total distance: %v", totalDiff)
	fmt.Println()
}

func printSimilarityScore (leftCol []int, rightCol []int ) {
	var similartyScore int

	for _, valueLeft := range leftCol {
		count := 0

		for _, valueRight := range rightCol {
			if valueLeft == valueRight {
				count ++
			}
		}

		similartyScore = similartyScore + (valueLeft * count)
	}

	fmt.Printf("Similarty score: %v", similartyScore)
	fmt.Println()
}

func calcDistance (a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
