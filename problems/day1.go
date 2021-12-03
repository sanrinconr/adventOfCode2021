package problems

import (
	"strconv"

	"github.com/sanrinconr/adventOfCode2021/reader"
)

// Day1 advent of code
type Day1 struct{}

// Execute count the number of times a depth measurement increases from the previous measurement. (There is no measurement before the first measurement.).
func (d *Day1) Execute() string {
	input := d.input()
	increased := 0

	for i := 1; i < len(input); i++ {
		if input[i-1] < input[i] {
			increased++
		}
	}

	return strconv.Itoa(increased)
}

func (d *Day1) input() []int {
	r := reader.Client{}

	input, err := r.ArrayInt("./inputs/day1.txt")
	if err != nil {
		panic(err)
	}

	return input
}
