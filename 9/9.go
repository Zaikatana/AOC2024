package main

import (
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("input_9.txt")
	check(err)
	strDat := string(dat)
	blocks := []string{}
	currId := 0
	m := map[string][]int{}
	emptySpaces := [][]int{}
	for cIdx, c := range strDat {
		count, countErr := strconv.Atoi(string(c))
		check(countErr)
		if cIdx%2 == 0 {
			for range count {
				blocks = append(blocks, strconv.Itoa(currId))
			}
			m[strconv.Itoa(currId)] = []int{len(blocks) - count, len(blocks)}
			currId += 1
		} else {
			for range count {
				blocks = append(blocks, ".")
			}
			emptySpaces = append(emptySpaces, []int{len(blocks) - count, len(blocks)})
		}
	}

	p2Total := 0
	p2Blocks := make([]string, len(blocks))
	_ = copy(p2Blocks, blocks)

	endPtr := len(blocks) - 1

	for i := 0; i < len(blocks); i++ {
		if i > endPtr {
			break
		}
		if blocks[i] == "." {
			temp := blocks[endPtr]
			blocks[endPtr] = blocks[i]
			blocks[i] = temp
			for blocks[endPtr] == "." {
				endPtr -= 1
			}
		}
	}

	p1Total := 0
	p1Blocks := make([]string, len(blocks))
	_ = copy(p1Blocks, blocks)
	for i := 0; i < len(p1Blocks); i++ {
		if p1Blocks[i] == "." {
			break
		}

		curr, currErr := strconv.Atoi(string(p1Blocks[i]))
		check(currErr)
		p1Total += (i * curr)
	}

	for j := currId - 1; j >= 0; j-- {
		idStart := m[strconv.Itoa(j)][0]
		idEnd := m[strconv.Itoa(j)][1]

		idCount := idEnd - idStart
		for _, emptySpace := range emptySpaces {
			if emptySpace[0] >= idStart {
				break
			}
			if emptySpace[1]-emptySpace[0] >= idCount {
				emptySpaceStart := emptySpace[0]
				for k := idStart; k < idEnd; k++ {
					temp := p2Blocks[emptySpaceStart]
					p2Blocks[emptySpaceStart] = p2Blocks[k]
					p2Blocks[k] = temp
					emptySpaceStart += 1
				}
				emptySpace[0] = emptySpaceStart
				break
			}
		}
	}

	for i := 0; i < len(p2Blocks); i++ {
		if string(p2Blocks[i]) != "." {
			curr, currErr := strconv.Atoi(string(p2Blocks[i]))
			check(currErr)
			p2Total += (i * curr)
		}
	}

	fmt.Println("Part 1 solution", p1Total)
	fmt.Println("Part 2 solution", p2Total)
}
