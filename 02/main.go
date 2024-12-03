// DAY 1 – https://adventofcode.com/2024/day/2

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	byte, _ := os.ReadFile("./input.txt")

	parsedFile := string(byte)
	splitFile := strings.Split(parsedFile, "\n")

	var safeRecords int

	for _, report := range splitFile {
		var parsedReport []int
		isSafe := true
		levels := strings.Split(report, " ")

		for _, level := range levels {
			parsedLevel, _ := strconv.Atoi(level)
			parsedReport = append(parsedReport, parsedLevel)
		}

		if isReportSafe(parsedReport) == false {
			isSafe = false
		}

		if isSafe == true {
			safeRecords++
		}
	}

	fmt.Printf("Safe Records: %v", safeRecords)
}

func isReportSafe (parsedReport []int) (isRecordSafe bool) {
	isRecordSafe = false

	if checkForAllDecreasing(parsedReport) == true {
		isRecordSafe = true
		return
	}
	
	if checkForAllIncreasing(parsedReport) == true {
		isRecordSafe = true
		return
	}

	return
}

func checkForAllDecreasing (parsedReport []int) (isSafe bool) {
	isSafe = true

	for i, currentValue := range parsedReport {
		if i < (len(parsedReport) - 1) {
			if parsedReport[i + 1] >= currentValue {
				isSafe = false
				return
			}
		}
	}

	for i, currentValue := range parsedReport {
		if i < (len(parsedReport) - 1) {
			diff := (currentValue - parsedReport[i + 1])

			if (diff < 0) {
				diff = diff * -1
			}

			if  diff < 1 || diff > 3 {
				isSafe = false
			}
		}
	}

	return
}

func checkForAllIncreasing (parsedReport []int) (isSafe bool) {
	isSafe = true

	for i, currentValue := range parsedReport {
		if i < (len(parsedReport) - 1) {
			if parsedReport[i + 1] <= currentValue {
				isSafe = false
				return
			}
		}
	}

	for i, currentValue := range parsedReport {
		if i < (len(parsedReport) - 1) {
			diff := (currentValue - parsedReport[i + 1])

			if (diff < 0) {
				diff = diff * -1
			}

			if  diff < 1 || diff > 3 {
				isSafe = false
			}
		}
	}

	return
}

