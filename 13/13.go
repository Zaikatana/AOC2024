package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	s "strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Button struct {
	X     int
	Y     int
	Value int
}

type Value struct {
	Val    int
	APress int
	BPress int
}

func retrieveCoord(str string) int {
	r, _ := regexp.Compile(`[0-9]+`)
	num := r.FindAllString(str, -1)

	numInt, _ := strconv.Atoi(num[0])

	return numInt
}

func main() {
	dat, err := os.ReadFile("input_13.txt")
	check(err)
	strDat := string(dat)
	p1Total := 0
	p2Total := 0

	list := s.Split(strDat, "\n\n")

	// More Comp sci-ey type solution attempting to create a 2D array utilising tabulation to solve it like a knapsack problem
	// Naturally when you add 10000000000000 to the prize X and Y values, theres not enough memory
	for _, item := range list {
		itemConfig := s.Split(item, "\n")
		aConfig := s.Split(itemConfig[0], " ")
		bConfig := s.Split(itemConfig[1], " ")
		prizeConfig := s.Split(itemConfig[2], " ")
		prizeX := retrieveCoord(prizeConfig[1][2:])
		prizeY := retrieveCoord(prizeConfig[2][2:])

		buttonA := Button{X: retrieveCoord(aConfig[2][2:]), Y: retrieveCoord(aConfig[3][2:]), Value: 3}
		buttonB := Button{X: retrieveCoord(bConfig[2][2:]), Y: retrieveCoord(bConfig[3][2:]), Value: 1}

		dp := make([][]Value, prizeY+1)
		for i := 0; i < prizeY+1; i++ {
			dp[i] = make([]Value, prizeX+1)
			for j := 0; j < prizeX+1; j++ {
				dp[i][j] = Value{Val: 0, APress: 0, BPress: 0}
			}
		}

		dp[buttonA.Y][buttonA.X] = Value{Val: buttonA.Value, APress: 1, BPress: 0}
		dp[buttonB.Y][buttonB.X] = Value{Val: buttonB.Value, APress: 0, BPress: 1}

		for rowIdx, row := range dp {
			for colIdx, col := range row {
				if colIdx >= buttonA.X && rowIdx >= buttonA.Y {
					if col.Val == 0 && dp[rowIdx-buttonA.Y][colIdx-buttonA.X].Val != 0 {
						dp[rowIdx][colIdx] = Value{Val: dp[rowIdx-buttonA.Y][colIdx-buttonA.X].Val + buttonA.Value, APress: dp[rowIdx-buttonA.Y][colIdx-buttonA.X].APress + 1, BPress: dp[rowIdx-buttonA.Y][colIdx-buttonA.X].BPress}
					} else {
						if min(col.Val, dp[rowIdx-buttonA.Y][colIdx-buttonA.X].Val+buttonA.Value) != col.Val {
							dp[rowIdx][colIdx] = Value{Val: dp[rowIdx-buttonA.Y][colIdx-buttonA.X].Val + buttonA.Value, APress: dp[rowIdx-buttonA.Y][colIdx-buttonA.X].APress + 1, BPress: dp[rowIdx-buttonA.Y][colIdx-buttonA.X].BPress}
						}
					}
				}
				if colIdx >= buttonB.X && rowIdx >= buttonB.Y {
					if col.Val == 0 && dp[rowIdx-buttonB.Y][colIdx-buttonB.X].Val != 0 {
						dp[rowIdx][colIdx] = Value{Val: dp[rowIdx-buttonB.Y][colIdx-buttonB.X].Val + buttonB.Value, APress: dp[rowIdx-buttonB.Y][colIdx-buttonB.X].APress, BPress: dp[rowIdx-buttonB.Y][colIdx-buttonB.X].BPress + 1}
					} else {
						if min(col.Val, dp[rowIdx-buttonB.Y][colIdx-buttonB.X].Val+buttonB.Value) != col.Val {
							dp[rowIdx][colIdx] = Value{Val: dp[rowIdx-buttonB.Y][colIdx-buttonB.X].Val + buttonB.Value, APress: dp[rowIdx-buttonB.Y][colIdx-buttonB.X].APress, BPress: dp[rowIdx-buttonB.Y][colIdx-buttonB.X].BPress + 1}
						}
					}
				}
			}
		}

		p1Total += dp[prizeY][prizeX].Val
	}

	fmt.Println("Part 1 solution", p1Total)

	// Ended up looking at reddit for Part 2 and saw Cramer's Rule :(
	// https://en.wikipedia.org/wiki/Cramer%27s_rule

	for _, item := range list {
		itemConfig := s.Split(item, "\n")
		aConfig := s.Split(itemConfig[0], " ")
		bConfig := s.Split(itemConfig[1], " ")
		prizeConfig := s.Split(itemConfig[2], " ")
		prizeX := retrieveCoord(prizeConfig[1][2:]) + 10000000000000
		prizeY := retrieveCoord(prizeConfig[2][2:]) + 10000000000000

		buttonA := Button{X: retrieveCoord(aConfig[2][2:]), Y: retrieveCoord(aConfig[3][2:]), Value: 3}
		buttonB := Button{X: retrieveCoord(bConfig[2][2:]), Y: retrieveCoord(bConfig[3][2:]), Value: 1}

		d := (buttonA.Y * buttonB.X) - (buttonB.Y * buttonA.X)
		dA := (prizeY * buttonB.X) - (prizeX * buttonB.Y)
		dB := (buttonA.Y * prizeX) - (buttonA.X * prizeY)

		aPresses := float64(dA) / float64(d)
		bPresses := float64(dB) / float64(d)

		if aPresses == math.Round(aPresses) && bPresses == math.Round(bPresses) {
			p2Total += int(math.Round(aPresses)*float64(buttonA.Value) + math.Round(bPresses)*float64(buttonB.Value))
		}

		p2Total += 0
	}

	fmt.Println("Part 2 solution", p2Total)
}
