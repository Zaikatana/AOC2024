package main

import (
	"fmt"
	"math"
	"os"
	s "strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Graph struct {
	Antinodes   []Pos
	Resonance   []Pos
	Frequencies map[rune][]Pos
	XLowerBound float64
	XUpperBound float64
	YLowerBound float64
	YUpperBound float64
}

func (g *Graph) IsAntinodePresent(p Pos) bool {
	for _, antinode := range g.Antinodes {
		if math.Round(antinode.X) == math.Round(p.X) && math.Round(antinode.Y) == math.Round(p.Y) {
			return true
		}
	}

	return false
}

func (g *Graph) IsResonancePresent(p Pos) bool {
	for _, antinode := range g.Resonance {
		if math.Round(antinode.X) == math.Round(p.X) && math.Round(antinode.Y) == math.Round(p.Y) {
			return true
		}
	}

	return false
}

func (g *Graph) InRangeAntinode(antinode Pos) bool {
	return math.Round(antinode.X) >= g.XLowerBound && math.Round(antinode.X) < g.XUpperBound && math.Round(antinode.Y) >= g.YLowerBound && math.Round(antinode.Y) < g.YUpperBound
}

type Pos struct {
	X float64
	Y float64
}

func generateGraph(rows []string) Graph {
	frequencies := map[rune][]Pos{}
	for rowIdx, row := range rows {
		for colIdx, col := range row {
			if col != '.' {
				pos := Pos{X: float64(colIdx), Y: float64(rowIdx)}
				_, freqPrs := frequencies[col]
				if !freqPrs {
					freqSlice := []Pos{pos}
					frequencies[col] = freqSlice
				} else {
					frequencies[col] = append(frequencies[col], pos)
				}
			}
		}
	}

	return Graph{Antinodes: []Pos{}, Resonance: []Pos{}, Frequencies: frequencies, XLowerBound: float64(0), XUpperBound: float64(len(rows[0])), YLowerBound: float64(0), YUpperBound: float64(len(rows))}
}

func main() {
	dat, err := os.ReadFile("input_8.txt")
	check(err)
	strDat := string(dat)
	p1Total := 0
	p2Total := 0
	rows := s.Split(strDat, "\n")
	g := generateGraph(rows)

	for freq := range g.Frequencies {
		for i := 0; i < len(g.Frequencies[freq])-1; i++ {
			coord0 := g.Frequencies[freq][i]
			for j := i + 1; j < len(g.Frequencies[freq]); j++ {
				coord1 := g.Frequencies[freq][j]
				slope := (coord1.Y - coord0.Y) / (coord1.X - coord0.X)

				distance := math.Sqrt(math.Pow(coord1.X-coord0.X, 2) + math.Pow(coord1.Y-coord0.Y, 2))

				antinodeA := coord0
				antinodeB := coord1
				isFirst := true

				for {
					// https://www.geeksforgeeks.org/find-points-at-a-given-distance-on-a-line-of-given-slope/
					if slope < 0 {
						antinodeA = Pos{X: antinodeA.X + distance*(1/math.Sqrt(1+math.Pow(slope, 2))), Y: antinodeA.Y + distance*(slope/math.Sqrt(1+math.Pow(slope, 2)))}
						antinodeB = Pos{X: antinodeB.X - distance*(1/math.Sqrt(1+math.Pow(slope, 2))), Y: antinodeB.Y - distance*(slope/math.Sqrt(1+math.Pow(slope, 2)))}
					} else {
						antinodeA = Pos{X: antinodeA.X - distance*(1/math.Sqrt(1+math.Pow(slope, 2))), Y: antinodeA.Y - distance*(slope/math.Sqrt(1+math.Pow(slope, 2)))}
						antinodeB = Pos{X: antinodeB.X + distance*(1/math.Sqrt(1+math.Pow(slope, 2))), Y: antinodeB.Y + distance*(slope/math.Sqrt(1+math.Pow(slope, 2)))}
					}

					if !g.InRangeAntinode(antinodeA) && !g.InRangeAntinode(antinodeB) {
						break
					}

					if isFirst {
						if !g.IsAntinodePresent(antinodeA) {
							g.Antinodes = append(g.Antinodes, antinodeA)
						}
						if !g.IsAntinodePresent(antinodeB) {
							g.Antinodes = append(g.Antinodes, antinodeB)
						}
						isFirst = false
					}

					if !g.IsResonancePresent(antinodeA) {
						g.Resonance = append(g.Resonance, antinodeA)
					}

					if !g.IsResonancePresent(antinodeB) {
						g.Resonance = append(g.Resonance, antinodeB)
					}
				}
			}
		}
	}

	for _, antinode := range g.Antinodes {
		if g.InRangeAntinode(antinode) {
			p1Total += 1
		}
	}

	for _, resonance := range g.Resonance {
		if g.InRangeAntinode(resonance) {
			p2Total += 1
		}
	}

	for freq := range g.Frequencies {
		for _, antennae := range g.Frequencies[freq] {
			if !g.IsResonancePresent(antennae) {
				p2Total += 1
			}
		}
	}

	fmt.Println("Part 1 Solution", p1Total)
	fmt.Println("Part 2 Solution", p2Total)
}
