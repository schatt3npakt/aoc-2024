// DAY 3: https://adventofcode.com/2024/day/3

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main () {
	file, _ := os.Open("./03/input.txt")
	sc := bufio.NewScanner(file)
	
	part1(sc)
	part2()
}

func part1 (sc *bufio.Scanner) {
	var operations []string
	var total int

	for sc.Scan() {
		re := regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)")
		txt := sc.Text()
		matches := re.FindAllString(txt, -1)
		operations = append(operations, matches...)
	}

	for _, exp := range operations {
		parsedExp := strings.Split(exp, ",")
		re := regexp.MustCompile("\\d{1,3}")
		firstValue, _ := strconv.Atoi(re.FindString(parsedExp[0]))
		secondValue, _ := strconv.Atoi(re.FindString(parsedExp[1]))

		total = total + (firstValue * secondValue)
	}

	fmt.Printf("Solution Part 1: %v", total)
	fmt.Println()
}

func part2 () {
	var operations []string
	var total int

	re := regexp.MustCompile("don't.*")

	byte, _ := os.ReadFile("./03/input.txt")
	fileString := strings.ReplaceAll(string(byte), "\n", "")
	fileString = strings.ReplaceAll(fileString, "don't()", "\ndon't()")
	fileString = strings.ReplaceAll(fileString, "do()", "\ndo()")
	fileString = re.ReplaceAllString(fileString, "")

	expressionRe := regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)")
	matches := expressionRe.FindAllString(fileString, -1)
	operations = append(operations, matches...)

	for _, exp := range operations {
		parsedExp := strings.Split(exp, ",")
		re := regexp.MustCompile("\\d{1,3}")
		firstValue, _ := strconv.Atoi(re.FindString(parsedExp[0]))
		secondValue, _ := strconv.Atoi(re.FindString(parsedExp[1]))

		total = total + (firstValue * secondValue)
	}

	fmt.Printf("Solution Part 2: %v", total)
	fmt.Println()
}