// DAY XX

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main () {
	file, _ := os.Open("./example/input.txt")
	sc := bufio.NewScanner(file)

	part1(sc)
	part2(sc)
}

func part1 (sc *bufio.Scanner) {
	for sc.Scan() {
		fmt.Println(sc.Text())
	}

	fmt.Println("Solution Part 1: ")
}

func part2 (sc *bufio.Scanner) {
	for sc.Scan() {
		fmt.Println(sc.Text())
	}

	fmt.Println("Solution Part 1: ")
}