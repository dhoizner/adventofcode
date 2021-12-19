package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day01.input", "Relative path to input file.")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		panic(err)
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	nums := []int{}
	for _, v := range lines {
		converted, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		nums = append(nums, converted)
	}

	countIncreasing := 0
	prev := -1
	// Part 1 - single number comparison
	for _, curr := range nums {
		if prev != -1 {
			if curr > prev {
				fmt.Println(curr, "(increased)")
				countIncreasing += 1
			} else {
				fmt.Println(curr, "(decreased)")
			}
		} else {
			fmt.Println(curr, "(N/A - no previous measurement)")
		}

		prev = curr
	}
	fmt.Println("Number of measurements larger than the previous measurement:", countIncreasing)

	countIncreasing = 0
	lastWindow := []int{0, 0, 0}
	lastSum := 0
	label := 65
	// Part 2 - 3-item rolling window comparison
	for idx, curr := range nums {
		sum := lastSum + curr - lastWindow[0]

		lastWindow = append(lastWindow[1:], curr)
		if idx < 3 {
			lastSum = sum
			if idx == 2 {
				fmt.Println(label, ":", sum, "(N/A - no previous measurement)")
				label += 1
			}
			continue
		}
		if sum > lastSum {
			countIncreasing += 1
			fmt.Println(label, ":", sum, "(increased)")
		} else {
			fmt.Println(label, ":", sum, "(decreased)")
		}
		lastSum = sum
		label += 1
	}

	fmt.Println("Number of sums larger than the previous sums:", countIncreasing)
}
