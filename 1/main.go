package main

import (
	"fmt"
	"strconv"

	"github.com/mrfojo/aoc_2019/common"
)

func main() {
	input, err := common.ReadInput("input.txt")
	if err != nil {
		panic(err)
	}
	totalFuel := 0
	for _, s := range input {
		mass, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		fuel := (mass / 3) - 2
		for fuel > 0 {
			totalFuel += fuel
			fuel = (fuel / 3) - 2
		}
	}
	fmt.Println(totalFuel)
}
