package main

import (
	"github.com/alexchao26/advent-of-code-go/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := util.ReadFile("../input.txt")
	sli := strings.Split(input, "\n")

	var sum int
	seen := make(map[int]bool)

	for {
		for _, instruction := range sli {
			sign := instruction[:1]
			num, _ := strconv.Atoi(instruction[1:])
			if sign == "+" {
				sum += num
			} else {
				sum -= num
			}

			if seen[sum] {
				fmt.Println("Number seen twice is:", sum)
				return
			}
			seen[sum] = true
		}
	}
}