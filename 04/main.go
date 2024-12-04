// DAY 04: https://adventofcode.com/2024/day/4

package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main () {
	byte, _ := os.ReadFile("./04/input.txt")
	file := strings.Split(string(byte), "\n")

	part1(file)
	part2(file)
}

func part1 (file []string) {
	var count int
	count = count + checkForXmasHorizontal(file)
	count = count + checkForSamxHorizontal(file)
	count = count + checkForXmasVertical(file)
	count = count + checkForSamxVertical(file)
	count = count + checkForXmasDiagonal(file)
	count = count + checkForSamxDiagonal(file)
	count = count + checkForXmasDiagonalReverse(file)
	count = count + checkForSamxDiagonalReverse(file)

	fmt.Printf("Solution Part 1: %v", count)
	fmt.Println()
}

func checkForXmasHorizontal (file []string) (count int) {
	re := regexp.MustCompile("XMAS")

	for _, line := range file {
		count += len(re.FindAllString(line, -1))
	}

	return
}

func checkForSamxHorizontal (file []string) (count int) {
	re := regexp.MustCompile("SAMX")

	for _, line := range file {
		count += len(re.FindAllString(line, -1))
	}

	return
}

func checkForXmasVertical (file []string) (count int) {
	re := regexp.MustCompile("X")
	
	for i, line := range file {
		charMatchesInCurrentLine := re.FindAllStringIndex(line, -1)

		for _, matchIndexes := range charMatchesInCurrentLine {
			indexOfMatchInCurrentLine := matchIndexes[0]

			if (i < (len(file) - 3)) {
				// Check for character M in line n + 1
				if (string(file[i+1][indexOfMatchInCurrentLine]) != "M") {
					continue
				}

				// Check for character A in line n + 2
				if (string(file[i+2][indexOfMatchInCurrentLine]) != "A") {
					continue
				}

				// Check for character S in line n + 3
				if (string(file[i+3][indexOfMatchInCurrentLine]) != "S") {
					continue
				}
				
				count++
			}
		}
	}

	return
}

func checkForSamxVertical (file []string) (count int) {
	re := regexp.MustCompile("S")
	
	for i, line := range file {
		charMatchesInCurrentLine := re.FindAllStringIndex(line, -1)

		for _, matchIndexes := range charMatchesInCurrentLine {
			indexOfMatchInCurrentLine := matchIndexes[0]

			if (i < (len(file) - 3)) {
				// Check for character A in line n + 1
				if (string(file[i+1][indexOfMatchInCurrentLine]) != "A") {
					continue
				}

				// Check for character M in line n + 2
				if (string(file[i+2][indexOfMatchInCurrentLine]) != "M") {
					continue
				}

				// Check for character X in line n + 3
				if (string(file[i+3][indexOfMatchInCurrentLine]) != "X") {
					continue
				}
				
				count++
			}
		}
	}

	return
}

func checkForXmasDiagonal (file []string) (count int) {
	re := regexp.MustCompile("X")
	
	for i, line := range file {
		charMatchesInCurrentLine := re.FindAllStringIndex(line, -1)

		for _, matchIndexes := range charMatchesInCurrentLine {
			indexOfMatchInCurrentLine := matchIndexes[0]

			if (i < (len(file) - 3) && indexOfMatchInCurrentLine < len(line) - 3) {
				// Check for character M in line n + 1, at position n + 1
				if (string(file[i+1][indexOfMatchInCurrentLine + 1]) != "M") {
					continue
				}

				// Check for character A in line n + 2, at position n + 2
				if (string(file[i+2][indexOfMatchInCurrentLine + 2]) != "A") {
					continue
				}

				// Check for character S in line n + 3, at position n + 3
				if (string(file[i+3][indexOfMatchInCurrentLine + 3]) != "S") {
					continue
				}

				count++
			}
		}
	}

	return
}

func checkForSamxDiagonal (file []string) (count int) {
	re := regexp.MustCompile("S")
	
	for i, line := range file {
		charMatchesInCurrentLine := re.FindAllStringIndex(line, -1)

		for _, matchIndexes := range charMatchesInCurrentLine {
			indexOfMatchInCurrentLine := matchIndexes[0]

			if (i < (len(file) - 3) && indexOfMatchInCurrentLine < len(line) - 3) {
				// Check for character A in line n + 1, at position n + 1
				if (string(file[i+1][indexOfMatchInCurrentLine + 1]) != "A") {
					continue
				}

				// Check for character M in line n + 2, at position n + 2
				if (string(file[i+2][indexOfMatchInCurrentLine + 2]) != "M") {
					continue
				}

				// Check for character X in line n + 3, at position n + 3
				if (string(file[i+3][indexOfMatchInCurrentLine + 3]) != "X") {
					continue
				}
				
				count++
			}
		}
	}

	return
}

func checkForXmasDiagonalReverse (file []string) (count int) {
	re := regexp.MustCompile("X")
	
	for i, line := range file {
		charMatchesInCurrentLine := re.FindAllStringIndex(line, -1)

		for _, matchIndexes := range charMatchesInCurrentLine {
			indexOfMatchInCurrentLine := matchIndexes[0]

			// check if wrong!
			if (i < (len(file) - 3) && indexOfMatchInCurrentLine > 2) {
				// Check for character M in line n + 1, at position n - 1
				if (string(file[i+1][indexOfMatchInCurrentLine - 1]) != "M") {
					continue
				}

				// Check for character A in line n + 2, at position n - 2
				if (string(file[i+2][indexOfMatchInCurrentLine - 2]) != "A") {
					continue
				}

				// Check for character S in line n + 3, at position n - 3
				if (string(file[i+3][indexOfMatchInCurrentLine - 3]) != "S") {
					continue
				}

				count++
			}
		}
	}

	return
}

func checkForSamxDiagonalReverse (file []string) (count int) {
	re := regexp.MustCompile("S")
	
	for i, line := range file {
		charMatchesInCurrentLine := re.FindAllStringIndex(line, -1)

		for _, matchIndexes := range charMatchesInCurrentLine {
			indexOfMatchInCurrentLine := matchIndexes[0]

			// check if wrong!
			if (i < (len(file) - 3) && indexOfMatchInCurrentLine > 2) {
				// Check for character A in line n + 1, at position n - 1
				if (string(file[i+1][indexOfMatchInCurrentLine - 1]) != "A") {
					continue
				}

				// Check for character M in line n + 2, at position n - 2
				if (string(file[i+2][indexOfMatchInCurrentLine - 2]) != "M") {
					continue
				}

				// Check for character X in line n + 3, at position n - 3
				if (string(file[i+3][indexOfMatchInCurrentLine - 3]) != "X") {
					continue
				}

				count++
			}
		}
	}

	return
}

func part2 (file []string) {
	var count int

	count = count + checkForMasMas(file)
	count = count + checkForMasSam(file)
	count = count + checkForSamMas(file)
	count = count + checkForSamSam(file)

	fmt.Printf("Solution Part 2: %v", count)
	fmt.Println()
}


func checkForMasMas (file []string) (count int) {
	re := regexp.MustCompile("M")
	
	for i, line := range file {
		charMatchesInCurrentLine := re.FindAllStringIndex(line, -1)

		for _, matchIndexes := range charMatchesInCurrentLine {
			indexOfMatchInCurrentLine := matchIndexes[0]

			if (i < (len(file) - 2) && indexOfMatchInCurrentLine < len(line) - 2) {
				// Check for character M in line n, at position n + 1
				if (string(file[i][indexOfMatchInCurrentLine + 2]) != "M") {
					continue
				}

				// Check for character A in line n + 1, at position n + 1
				if (string(file[i+1][indexOfMatchInCurrentLine + 1]) != "A") {
					continue
				}

				// Check for character S in line n+2, at position n
				if (string(file[i+2][indexOfMatchInCurrentLine]) != "S") {
					continue
				}

				// Check for character S in line n+2, at position n + 1
				if (string(file[i+2][indexOfMatchInCurrentLine +2]) != "S") {
					continue
				}

				count++
			}
		}
	}

	return
}

func checkForMasSam (file []string) (count int) {
	re := regexp.MustCompile("M")
	
	for i, line := range file {
		charMatchesInCurrentLine := re.FindAllStringIndex(line, -1)

		for _, matchIndexes := range charMatchesInCurrentLine {
			indexOfMatchInCurrentLine := matchIndexes[0]

			if (i < (len(file) - 2) && indexOfMatchInCurrentLine < len(line) - 2) {
				// Check for character S in line n, at position n + 1
				if (string(file[i][indexOfMatchInCurrentLine + 2]) != "S") {
					continue
				}

				// Check for character A in line n + 1, at position n + 1
				if (string(file[i+1][indexOfMatchInCurrentLine + 1]) != "A") {
					continue
				}

				// Check for character M in line n+2, at position n
				if (string(file[i+2][indexOfMatchInCurrentLine]) != "M") {
					continue
				}

				// Check for character S in line n+2, at position n + 1
				if (string(file[i+2][indexOfMatchInCurrentLine +2]) != "S") {
					continue
				}

				count++
			}
		}
	}

	return
}

func checkForSamMas (file []string) (count int) {
	re := regexp.MustCompile("S")
	
	for i, line := range file {
		charMatchesInCurrentLine := re.FindAllStringIndex(line, -1)

		for _, matchIndexes := range charMatchesInCurrentLine {
			indexOfMatchInCurrentLine := matchIndexes[0]

			if (i < (len(file) - 2) && indexOfMatchInCurrentLine < len(line) - 2) {
				// Check for character M in line n, at position n + 1
				if (string(file[i][indexOfMatchInCurrentLine + 2]) != "M") {
					continue
				}

				// Check for character A in line n + 1, at position n + 1
				if (string(file[i+1][indexOfMatchInCurrentLine + 1]) != "A") {
					continue
				}

				// Check for character S in line n+2, at position n
				if (string(file[i+2][indexOfMatchInCurrentLine]) != "S") {
					continue
				}

				// Check for character S in line n+2, at position n + 1
				if (string(file[i+2][indexOfMatchInCurrentLine +2]) != "M") {
					continue
				}

				count++
			}
		}
	}

	return
}

func checkForSamSam (file []string) (count int) {
	re := regexp.MustCompile("S")
	
	for i, line := range file {
		charMatchesInCurrentLine := re.FindAllStringIndex(line, -1)

		for _, matchIndexes := range charMatchesInCurrentLine {
			indexOfMatchInCurrentLine := matchIndexes[0]

			if (i < (len(file) - 2) && indexOfMatchInCurrentLine < len(line) - 2) {
				// Check for character S in line n, at position n + 1
				if (string(file[i][indexOfMatchInCurrentLine + 2]) != "S") {
					continue
				}

				// Check for character A in line n + 1, at position n + 1
				if (string(file[i+1][indexOfMatchInCurrentLine + 1]) != "A") {
					continue
				}

				// Check for character M in line n+2, at position n
				if (string(file[i+2][indexOfMatchInCurrentLine]) != "M") {
					continue
				}

				// Check for character M in line n+2, at position n + 1
				if (string(file[i+2][indexOfMatchInCurrentLine +2]) != "M") {
					continue
				}

				count++
			}
		}
	}

	return
}