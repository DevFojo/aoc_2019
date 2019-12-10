package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mrfojo/aoc_2019/common"
)

func main() {
	input, err := common.ReadInput("input.txt")
	if err != nil {
		panic(err)
	}
	for n := 0; n < 100; n++ {
		for v := 0; v < 100; v++ {
			program, err := parse(input[0])
			if err != nil {
				panic(err)
			}
			program[1] = n
			program[2] = v
			for i := 0; i < len(program); i += 4 {
				if program[i] == 99 {
					break
				}
				switch program[i] {

				case 1:
					program[program[i+3]] = program[program[i+1]] + program[program[i+2]]
				case 2:
					program[program[i+3]] = program[program[i+1]] * program[program[i+2]]
				}
			}
			if program[0] == 19690720 {
				fmt.Println(100*n + v)
			}
		}
	}
}

func parse(s string) ([]int, error) {
	strArr := strings.Split(s, ",")
	intArr := make([]int, len(strArr))
	for i, t := range strArr {
		x, err := strconv.Atoi(t)
		if err != nil {
			return nil, err
		}
		intArr[i] = x
	}
	return intArr, nil
}
