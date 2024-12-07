package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	s "strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("input_6.txt")
	check(err)
	strDat := string(dat)
	rows := s.Split(strDat, "\n")
	p1Total := 0
	currRowIdx := -1
	currColIdx := -1
	origRowIdx := 0
	origColIdx := 0

	m := map[string][]int{}

	for rowIdx, row := range rows {
		for colIdx, col := range row {
			if col == '^' {
				currRowIdx = rowIdx
				origRowIdx = rowIdx
				currColIdx = colIdx
				origColIdx = colIdx
				break
			}
		}
		if currRowIdx > -1 && currColIdx > -1 {
			break
		}
	}

	// 1 = Up
	// 2 = Down
	// 3 = Left
	// 4 = Right
	currDirection := 1

	for currRowIdx > -1 && currRowIdx < len(rows) && currColIdx > -1 && currColIdx < len(rows[0]) {
		m[strconv.Itoa(currRowIdx)+"|"+strconv.Itoa(currColIdx)] = append(m[strconv.Itoa(currRowIdx)+"|"+strconv.Itoa(currColIdx)], currDirection)
		if currDirection == 1 {
			if currRowIdx-1 > -1 && rows[currRowIdx-1][currColIdx] == '#' {
				currDirection = 4
			} else {
				currRowIdx -= 1
			}
		} else if currDirection == 2 {
			if currRowIdx+1 < len(rows) && rows[currRowIdx+1][currColIdx] == '#' {
				currDirection = 3
			} else {
				currRowIdx += 1
			}
		} else if currDirection == 3 {
			if currColIdx-1 > -1 && rows[currRowIdx][currColIdx-1] == '#' {
				currDirection = 1
			} else {
				currColIdx -= 1
			}
		} else {
			if currColIdx+1 < len(rows[0]) && rows[currRowIdx][currColIdx+1] == '#' {
				currDirection = 2
			} else {
				currColIdx += 1
			}
		}
	}

	p1Total = len(m)
	p2Total := 0

	for rowIdx, row := range rows {
		for colIdx := range row {
			if rows[rowIdx][colIdx] == '#' {
				continue
			}
			currRowIdx = origRowIdx
			currColIdx = origColIdx
			currDirection = 1
			clear(m)

			for currRowIdx > -1 && currRowIdx < len(rows) && currColIdx > -1 && currColIdx < len(rows[0]) {
				if slices.Contains(m[strconv.Itoa(currRowIdx)+"|"+strconv.Itoa(currColIdx)], currDirection) {
					p2Total += 1
					break
				}
				m[strconv.Itoa(currRowIdx)+"|"+strconv.Itoa(currColIdx)] = append(m[strconv.Itoa(currRowIdx)+"|"+strconv.Itoa(currColIdx)], currDirection)

				if currDirection == 1 {
					if currRowIdx-1 > -1 && (rows[currRowIdx-1][currColIdx] == '#' || (currRowIdx-1 == rowIdx && currColIdx == colIdx)) {
						currDirection = 4
					} else {
						currRowIdx -= 1
					}
				} else if currDirection == 2 {
					if currRowIdx+1 < len(rows) && (rows[currRowIdx+1][currColIdx] == '#' || (currRowIdx+1 == rowIdx && currColIdx == colIdx)) {
						currDirection = 3
					} else {
						currRowIdx += 1
					}
				} else if currDirection == 3 {
					if currColIdx-1 > -1 && (rows[currRowIdx][currColIdx-1] == '#' || (currRowIdx == rowIdx && currColIdx-1 == colIdx)) {
						currDirection = 1
					} else {
						currColIdx -= 1
					}
				} else {
					if currColIdx+1 < len(rows[0]) && (rows[currRowIdx][currColIdx+1] == '#' || (currRowIdx == rowIdx && currColIdx+1 == colIdx)) {
						currDirection = 2
					} else {
						currColIdx += 1
					}
				}
			}
		}
	}

	fmt.Println("Part 1 Solution", p1Total)
	fmt.Println("Part 2 Solution", p2Total)
}
