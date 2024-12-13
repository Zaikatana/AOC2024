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

type Pos struct {
	X int
	Y int
}

type Plant struct {
	Position  Pos
	Plant     rune
	Perimeter int
	Up        bool
	Down      bool
	Left      bool
	Right     bool
}

func (pl *Plant) CalculatePerimeter(garden []string) {
	total := 0
	plX := pl.Position.X
	plY := pl.Position.Y

	if plX-1 < 0 || plX+1 >= len(garden[0]) {
		if plX-1 < 0 {
			pl.Left = true
		} else {
			pl.Right = true
		}
		total += 1
	}

	if plY-1 < 0 || plY+1 >= len(garden) {
		if plY-1 < 0 {
			pl.Up = true
		} else {
			pl.Down = true
		}
		total += 1
	}

	if plX+1 < len(garden[0]) && rune(garden[plY][plX+1]) != pl.Plant {
		pl.Right = true
		total += 1
	}

	if plX-1 >= 0 && rune(garden[plY][plX-1]) != pl.Plant {
		pl.Left = true
		total += 1
	}

	if plY+1 < len(garden) && rune(garden[plY+1][plX]) != pl.Plant {
		pl.Down = true
		total += 1
	}

	if plY-1 >= 0 && rune(garden[plY-1][plX]) != pl.Plant {
		pl.Up = true
		total += 1
	}

	pl.Perimeter = total
}

type Plot struct {
	Plants []Plant
}

func (p *Plot) GetTotalArea() int {
	return len(p.Plants)
}

func (p *Plot) GetTotalPerimeter() int {
	total := 0
	for _, plant := range p.Plants {
		total += plant.Perimeter
	}

	return total
}

func (p *Plot) GetTotalPrice() int {
	return p.GetTotalArea() * p.GetTotalPerimeter()
}

func (p *Plot) GetDiscountedPrice() int {
	return p.GetTotalArea() * p.GetSides()
}

func (p *Plot) GetSides() int {
	// area of 1 & 2 is equal to 4 regardless of shape
	if len(p.Plants) < 3 {
		return 4
	}

	total := 0
	upMap := map[string]bool{}
	downMap := map[string]bool{}
	rightMap := map[string]bool{}
	leftMap := map[string]bool{}

	for _, plant := range p.Plants {
		key := strconv.Itoa(plant.Position.X) + "|" + strconv.Itoa(plant.Position.Y)
		if plant.Up {
			upMap[key] = true
		}

		if plant.Down {
			downMap[key] = true
		}

		if plant.Left {
			leftMap[key] = true
		}

		if plant.Right {
			rightMap[key] = true
		}
	}

	for upKey := range upMap {
		if upMap[upKey] {
			stack := []string{upKey}

			for len(stack) != 0 {
				poppedElement := stack[0]
				stack = slices.Delete(stack, 0, 1)
				_, prs := upMap[poppedElement]
				if prs && upMap[poppedElement] {
					upMap[poppedElement] = false

					strCoord := s.Split(poppedElement, "|")
					strCoordX, _ := strconv.Atoi(strCoord[0])

					stack = append(stack, strconv.Itoa(strCoordX+1)+"|"+strCoord[1])
					stack = append(stack, strconv.Itoa(strCoordX-1)+"|"+strCoord[1])
				}
			}

			total += 1
		}
	}

	for downKey := range downMap {
		if downMap[downKey] {
			stack := []string{downKey}

			for len(stack) != 0 {
				poppedElement := stack[0]
				stack = slices.Delete(stack, 0, 1)
				_, prs := downMap[poppedElement]
				if prs && downMap[poppedElement] {
					downMap[poppedElement] = false

					strCoord := s.Split(poppedElement, "|")
					strCoordX, _ := strconv.Atoi(strCoord[0])

					stack = append(stack, strconv.Itoa(strCoordX+1)+"|"+strCoord[1])
					stack = append(stack, strconv.Itoa(strCoordX-1)+"|"+strCoord[1])
				}
			}

			total += 1
		}
	}

	for leftKey := range leftMap {
		if leftMap[leftKey] {
			stack := []string{leftKey}

			for len(stack) != 0 {
				poppedElement := stack[0]
				stack = slices.Delete(stack, 0, 1)
				_, prs := leftMap[poppedElement]
				if prs && leftMap[poppedElement] {
					leftMap[poppedElement] = false

					strCoord := s.Split(poppedElement, "|")
					strCoordY, _ := strconv.Atoi(strCoord[1])

					stack = append(stack, strCoord[0]+"|"+strconv.Itoa(strCoordY+1))
					stack = append(stack, strCoord[0]+"|"+strconv.Itoa(strCoordY-1))
				}
			}

			total += 1
		}
	}

	for rightKey := range rightMap {
		if rightMap[rightKey] {
			stack := []string{rightKey}

			for len(stack) != 0 {
				poppedElement := stack[0]
				stack = slices.Delete(stack, 0, 1)
				_, prs := rightMap[poppedElement]
				if prs && rightMap[poppedElement] {
					rightMap[poppedElement] = false

					strCoord := s.Split(poppedElement, "|")
					strCoordY, _ := strconv.Atoi(strCoord[1])

					stack = append(stack, strCoord[0]+"|"+strconv.Itoa(strCoordY+1))
					stack = append(stack, strCoord[0]+"|"+strconv.Itoa(strCoordY-1))
				}
			}

			total += 1
		}
	}

	return total
}

func (p *Plot) AddToPlot(pl Plant) {
	p.Plants = append(p.Plants, pl)
}

type Garden struct {
	Plots []Plot
}

func (g *Garden) AddToGarden(p Plot) {
	g.Plots = append(g.Plots, p)
}

func main() {
	dat, err := os.ReadFile("input_12.txt")
	check(err)
	strDat := string(dat)
	p1Total := 0
	p2Total := 0

	garden := s.Split(strDat, "\n")
	visited := map[string]bool{}
	stack := []Plant{}
	g := Garden{Plots: []Plot{}}

	for rowIdx, row := range garden {
		for colIdx, col := range row {
			// create plant and plot and add to stack if not visited
			key := strconv.Itoa(colIdx) + "|" + strconv.Itoa(rowIdx)
			_, prs := visited[key]
			if !prs {
				plot := Plot{Plants: []Plant{}}
				pl := Plant{Plant: rune(col), Position: Pos{X: colIdx, Y: rowIdx}}
				pl.CalculatePerimeter(garden)
				stack = append(stack, pl)

				for len(stack) != 0 {
					poppedElement := stack[0]
					stack = slices.Delete(stack, 0, 1)
					plX := poppedElement.Position.X
					plY := poppedElement.Position.Y
					poppedKey := strconv.Itoa(plX) + "|" + strconv.Itoa(plY)
					_, prsPopped := visited[poppedKey]
					if !prsPopped {
						plot.AddToPlot(poppedElement)
						visited[poppedKey] = true
						// grab neighbours of poppedElement and add to stack
						if plX+1 < len(garden[0]) && rune(garden[plY][plX+1]) == poppedElement.Plant {
							plA := Plant{Plant: rune(garden[plY][plX+1]), Position: Pos{X: plX + 1, Y: plY}}
							plA.CalculatePerimeter(garden)
							stack = append(stack, plA)
						}

						if plX-1 >= 0 && rune(garden[plY][plX-1]) == poppedElement.Plant {
							plB := Plant{Plant: rune(garden[plY][plX-1]), Position: Pos{X: plX - 1, Y: plY}}
							plB.CalculatePerimeter(garden)
							stack = append(stack, plB)
						}

						if plY+1 < len(garden) && rune(garden[plY+1][plX]) == poppedElement.Plant {
							plC := Plant{Plant: rune(garden[plY+1][plX]), Position: Pos{X: plX, Y: plY + 1}}
							plC.CalculatePerimeter(garden)
							stack = append(stack, plC)
						}

						if plY-1 >= 0 && rune(garden[plY-1][plX]) == poppedElement.Plant {
							plD := Plant{Plant: rune(garden[plY-1][plX]), Position: Pos{X: plX, Y: plY - 1}}
							plD.CalculatePerimeter(garden)
							stack = append(stack, plD)
						}
					}
				}
				g.AddToGarden(plot)
			}

		}
	}

	for _, plot := range g.Plots {
		p1Total += plot.GetTotalPrice()
	}

	for _, plot := range g.Plots {
		fmt.Printf("A region of %s plants with price %d * %d = %d\n", string(plot.Plants[0].Plant), plot.GetTotalArea(), plot.GetSides(), plot.GetDiscountedPrice())
		p2Total += plot.GetDiscountedPrice()
	}

	fmt.Println("Part 1 solution", p1Total)
	fmt.Println("Part 2 solution", p2Total)
}
