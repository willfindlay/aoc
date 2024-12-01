package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/willfindlay/aoc/cast"
	"github.com/willfindlay/aoc/util"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	first, second := parseInput(input)
	sort.Ints(first)
	sort.Ints(second)

	ans := 0
	for i := range first {
		if first[i] > second[i] {
			ans += first[i] - second[i]
		} else {
			ans += second[i] - first[i]
		}
	}

	return ans
}

func part2(input string) int {
	first, second := parseInput(input)
	counts := make(map[int]int)
	for _, n := range second {
		counts[n] += 1
	}
	ans := 0
	for _, n := range first {
		ans += counts[n] * n
	}
	return ans
}

func parseInput(input string) (first, second []int) {
	for _, line := range strings.Split(input, "\n") {
		nums := strings.Fields(line)
		first = append(first, cast.ToInt(nums[0]))
		second = append(second, cast.ToInt(nums[1]))
	}
	return first, second
}
