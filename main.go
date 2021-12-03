package main

import (
	"fmt"

	"github.com/sanrinconr/adventOfCode2021/problems"
)

func main() {
	pb := problems.Day1{}
	response := pb.Execute()

	fmt.Printf("the response is '%s'", response)
}
