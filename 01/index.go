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

	printTotalDistance(leftCol, rightCol)
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

	for i, _ := range leftCol {
		totalDiff = totalDiff + calcDistance(leftCol[i], rightCol[i])
	}

	fmt.Printf("Total distance: %v", totalDiff)
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

	fmt.Println()
	fmt.Printf("Similarty score: %v", similartyScore)
}

func calcDistance (a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
