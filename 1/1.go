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

func sortSlice(arr []string) {
	slices.SortFunc(arr, func(a, b string) int {
		a_int, err_a := strconv.Atoi(a)
		b_int, err_b := strconv.Atoi(b)
		check(err_a)
		check(err_b)

		return a_int - b_int
	})
}

func main() {
	dat, err := os.ReadFile("input_1.txt")
	check(err)
	str_dat := string(dat)
	arr := s.Split(str_dat, "\n")
	var l []string
	var r []string

	for _, element := range arr {
		entry := s.Split(element, "   ")
		l = append(l, entry[0])
		r = append(r, entry[1])
	}

	sortSlice(l)
	sortSlice(r)

	total_p1 := 0

	for i := 0; i < len(r); i++ {
		a_int, err_a := strconv.Atoi(l[i])
		b_int, err_b := strconv.Atoi(r[i])
		check(err_a)
		check(err_b)

		diff := a_int - b_int
		total_p1 += max(diff, -diff)
	}

	fmt.Println("Part 1 answer", total_p1)

	m := map[string]int{}
	total_p2 := 0

	for _, element := range l {
		val, ok := m[element]
		element_int, element_err := strconv.Atoi(element)
		check(element_err)

		if ok == true {
			total_p2 += (val * element_int)
		} else {
			m[element] = 0

			for _, r_element := range r {
				if r_element == element {
					m[element] += 1
				}
			}

			total_p2 += (m[element] * element_int)
		}
	}

	fmt.Println("Part 2 answer", total_p2)
}
