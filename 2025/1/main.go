package main

import (
	"fmt"

	common "github.com/msolimans/aoc/common"
)

var arr = getArr()

func getArr() []int {
	arr := [100]int{}
	for i := 0; i < 100; i++ {
		arr[i] = i
	}
	return arr[:]
}

func increment(current int, shift int, includePassThru bool) int {
	index := current
	for shift > 0 {
		index++
		if index > 99 {
			index = 0
		}

		shift--
		if index == 0 && includePassThru {
			arr[0] = arr[0] + 1 // pass thru
		}
	}

	return index // stop index

}

func decrement(current, shift int, includePassThru bool) int {
	index := current
	for shift > 0 {
		index--
		if index < 0 {
			index = 99
			// arr[0]++ // pass thru
		}
		shift--
		if index == 0 && includePassThru {
			arr[0] = arr[0] + 1 // pass thru
		}
	}
	return index // stop index
}

func seek(direction rune, dialStart int, distance int, includePassThru bool) int {
	switch direction {
	case 'L':
		dialStart = decrement(dialStart, distance, includePassThru)
	case 'R':
		dialStart = increment(dialStart, distance, includePassThru)
	}
	return dialStart
}

func Do(dialStart int, inputs []string, includePassThru bool) int {
	for _, line := range inputs {
		_ = line
		direction := line[0]
		distance := 0
		fmt.Sscanf(line[1:], "%d", &distance)
		dialStart = seek(rune(direction), dialStart, distance, includePassThru)
		if !includePassThru && dialStart == 0 {
			arr[0] = arr[0] + 1 // count only stops
		}

	}

	return dialStart

}

func main() {
	lines, err := common.ReadLines("./input.txt")
	if err != nil {
		panic(err)
	}
	Do(50, lines, true)
	fmt.Printf("Zero Clicks: %d\n", arr[0])
}

// results:
// 1st round: 1064
// 2nd round: 6122
