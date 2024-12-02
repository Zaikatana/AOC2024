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

func process(arr []string) bool {
	second, err_second := strconv.Atoi(arr[1])
	first, err_first := strconv.Atoi(arr[0])
	check(err_second)
	check(err_first)

	diff := second - first

	is_ascending := true
	if diff == 0 {
		return false
	} else if diff < 0 {
		is_ascending = false
	}

	for i := 0; i < len(arr)-1; i++ {
		int_a, err_a := strconv.Atoi(arr[i])
		int_b, err_b := strconv.Atoi(arr[i+1])
		check(err_a)
		check(err_b)

		diff := int_b - int_a

		if (diff < 0 && is_ascending) || (diff > 0 && !is_ascending) {
			return false
		}

		if max(diff, -diff) > 3 || max(diff, -diff) == 0 {
			return false
		}
	}

	return true
}

func main() {
	dat, err := os.ReadFile("input_2.txt")
	check(err)
	str_dat := string(dat)
	arr := s.Split(str_dat, "\n")

	safe_total := 0
	safe_pd_total := 0

	for _, element := range arr {
		entry := s.Split(element, " ")
		if process(entry) {
			safe_total++
			safe_pd_total++
		} else {
			for i := 0; i < len(entry); i++ {
				entry_clone := slices.Clone(entry)
				entry_clone = append(entry_clone[:i], entry_clone[i+1:]...)
				if process(entry_clone) {
					safe_pd_total++
					break
				}
			}
		}
	}

	fmt.Println("Part 1 answer", safe_total)
	fmt.Println("Part 2 answer", safe_pd_total)
}
