package main

import (
	"fmt"
	"os"
	s "strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkForXmas(curr string, colIdx int, rowIdx int, xmasMap []string, direction int) int {
	total := 0
	switch curr {
	case "X":
		if rowIdx-1 >= 0 && colIdx-1 >= 0 && xmasMap[rowIdx-1][colIdx-1] == 'M' {
			total += checkForXmas("M", colIdx-1, rowIdx-1, xmasMap, 1)
		}
		if rowIdx-1 >= 0 && xmasMap[rowIdx-1][colIdx] == 'M' {
			total += checkForXmas("M", colIdx, rowIdx-1, xmasMap, 2)
		}
		if colIdx+1 < len(xmasMap[0]) && rowIdx-1 >= 0 && xmasMap[rowIdx-1][colIdx+1] == 'M' {
			total += checkForXmas("M", colIdx+1, rowIdx-1, xmasMap, 3)
		}
		if colIdx-1 >= 0 && xmasMap[rowIdx][colIdx-1] == 'M' {
			total += checkForXmas("M", colIdx-1, rowIdx, xmasMap, 4)
		}
		if colIdx+1 < len(xmasMap[0]) && xmasMap[rowIdx][colIdx+1] == 'M' {
			total += checkForXmas("M", colIdx+1, rowIdx, xmasMap, 5)
		}
		if rowIdx+1 < len(xmasMap) && colIdx-1 >= 0 && xmasMap[rowIdx+1][colIdx-1] == 'M' {
			total += checkForXmas("M", colIdx-1, rowIdx+1, xmasMap, 6)
		}
		if rowIdx+1 < len(xmasMap) && xmasMap[rowIdx+1][colIdx] == 'M' {
			total += checkForXmas("M", colIdx, rowIdx+1, xmasMap, 7)
		}
		if rowIdx+1 < len(xmasMap) && colIdx+1 < len(xmasMap[0]) && xmasMap[rowIdx+1][colIdx+1] == 'M' {
			total += checkForXmas("M", colIdx+1, rowIdx+1, xmasMap, 8)
		}
	case "M":
		if direction == 1 && rowIdx-1 >= 0 && colIdx-1 >= 0 && xmasMap[rowIdx-1][colIdx-1] == 'A' {
			total += checkForXmas("A", colIdx-1, rowIdx-1, xmasMap, direction)
		}
		if direction == 2 && rowIdx-1 >= 0 && xmasMap[rowIdx-1][colIdx] == 'A' {
			total += checkForXmas("A", colIdx, rowIdx-1, xmasMap, direction)
		}
		if direction == 3 && colIdx+1 < len(xmasMap[0]) && rowIdx-1 >= 0 && xmasMap[rowIdx-1][colIdx+1] == 'A' {
			total += checkForXmas("A", colIdx+1, rowIdx-1, xmasMap, direction)
		}
		if direction == 4 && colIdx-1 >= 0 && xmasMap[rowIdx][colIdx-1] == 'A' {
			total += checkForXmas("A", colIdx-1, rowIdx, xmasMap, direction)
		}
		if direction == 5 && colIdx+1 < len(xmasMap[0]) && xmasMap[rowIdx][colIdx+1] == 'A' {
			total += checkForXmas("A", colIdx+1, rowIdx, xmasMap, direction)
		}
		if direction == 6 && rowIdx+1 < len(xmasMap) && colIdx-1 >= 0 && xmasMap[rowIdx+1][colIdx-1] == 'A' {
			total += checkForXmas("A", colIdx-1, rowIdx+1, xmasMap, direction)
		}
		if direction == 7 && rowIdx+1 < len(xmasMap) && xmasMap[rowIdx+1][colIdx] == 'A' {
			total += checkForXmas("A", colIdx, rowIdx+1, xmasMap, direction)
		}
		if direction == 8 && rowIdx+1 < len(xmasMap) && colIdx+1 < len(xmasMap[0]) && xmasMap[rowIdx+1][colIdx+1] == 'A' {
			total += checkForXmas("A", colIdx+1, rowIdx+1, xmasMap, direction)
		}
	case "A":
		if (direction == 1 && rowIdx-1 >= 0 && colIdx-1 >= 0 && xmasMap[rowIdx-1][colIdx-1] == 'S') ||
			(direction == 2 && rowIdx-1 >= 0 && xmasMap[rowIdx-1][colIdx] == 'S') ||
			(direction == 3 && colIdx+1 < len(xmasMap[0]) && rowIdx-1 >= 0 && xmasMap[rowIdx-1][colIdx+1] == 'S') ||
			(direction == 4 && colIdx-1 >= 0 && xmasMap[rowIdx][colIdx-1] == 'S') ||
			(direction == 5 && colIdx+1 < len(xmasMap[0]) && xmasMap[rowIdx][colIdx+1] == 'S') ||
			(direction == 6 && rowIdx+1 < len(xmasMap) && colIdx-1 >= 0 && xmasMap[rowIdx+1][colIdx-1] == 'S') ||
			(direction == 7 && rowIdx+1 < len(xmasMap) && xmasMap[rowIdx+1][colIdx] == 'S') ||
			(direction == 8 && rowIdx+1 < len(xmasMap) && colIdx+1 < len(xmasMap[0]) && xmasMap[rowIdx+1][colIdx+1] == 'S') {
			return 1
		}
	}
	return total
}

func checkForCrossmass(colIdx int, rowIdx int, xmasMap []string) int {
	if colIdx == 0 || rowIdx == 0 || colIdx == len(xmasMap[0])-1 || rowIdx == len(xmasMap)-1 {
		return 0
	}

	if ((xmasMap[rowIdx-1][colIdx-1] == 'M' && xmasMap[rowIdx+1][colIdx+1] == 'S') && ((xmasMap[rowIdx-1][colIdx+1] == 'M' && xmasMap[rowIdx+1][colIdx-1] == 'S') || (xmasMap[rowIdx+1][colIdx-1] == 'M' && xmasMap[rowIdx-1][colIdx+1] == 'S'))) ||
		((xmasMap[rowIdx+1][colIdx+1] == 'M' && xmasMap[rowIdx-1][colIdx-1] == 'S') && ((xmasMap[rowIdx-1][colIdx+1] == 'M' && xmasMap[rowIdx+1][colIdx-1] == 'S') || (xmasMap[rowIdx+1][colIdx-1] == 'M' && xmasMap[rowIdx-1][colIdx+1] == 'S'))) {
		return 1
	}

	return 0
}

func main() {
	dat, err := os.ReadFile("input_4.txt")
	check(err)
	strDat := string(dat)
	xmasMap := s.Split(strDat, "\n")
	p1Total := 0
	p2Total := 0

	for rowIdx, map_row := range xmasMap {
		for colIdx, map_col := range map_row {
			if map_col == 'X' {
				p1Total += checkForXmas(string(map_col), colIdx, rowIdx, xmasMap, 0)
			} else if map_col == 'A' {
				p2Total += checkForCrossmass(colIdx, rowIdx, xmasMap)
			}
		}
	}

	fmt.Println("Part 1 Solution", p1Total)
	fmt.Println("Part 2 Solution", p2Total)
}
