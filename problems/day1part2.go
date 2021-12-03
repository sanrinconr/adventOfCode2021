package problems

import (
	"fmt"
	"strconv"

	"github.com/sanrinconr/adventOfCode2021/reader"
)

// Day1Part2 advent of code
type Day1Part2 struct{}

// Execute count the number of times a depth measurement increases from the previous measurement. (There is no measurement before the first measurement.).
func (d *Day1Part2) Execute() string {
	input := d.input()
	increased := 0

	for i := 4; i <= len(input); i++ {
		sumPrevious := input[i-2] + input[i-3] + input[i-4] // 2,1,0
		sumActual := input[i-1] + input[i-2] + input[i-3]   // 3,2,1

		if sumActual > sumPrevious {
			fmt.Printf("%v (increased +) \n", sumActual)
			increased++
		} else {
			fmt.Printf("%v (decreased) \n", sumActual)
		}
	}

	return strconv.Itoa(increased)
}

func (d *Day1Part2) input() []int {
	r := reader.Client{}

	input, err := r.ArrayInt("./inputs/day1.txt")
	if err != nil {
		panic(err)
	}

	return input
}
