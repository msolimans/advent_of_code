package main

import (
	"fmt"
	"strings"

	common "github.com/msolimans/aoc/common"
)

func HasNRepeatingPattern(num int) bool {
	str := fmt.Sprintf("%d", num)
	length := len(str)

	// if it starts with 0, false
	if str[0] == '0' {
		return false
	}

	for patternLen := 1; patternLen <= length/2; patternLen++ {
		// Check if the total length is divisible by pattern length
		if length%patternLen != 0 {
			continue
		}

		pattern := str[:patternLen]

		// Check if the entire string is made of this pattern repeated
		isRepeating := true
		for i := patternLen; i < length; i += patternLen {
			if str[i:i+patternLen] != pattern {
				isRepeating = false
				break
			}
		}

		if isRepeating {
			return true
		}
	}

	return false
}

func Has2RepeatingPattern(num int) bool {
	str := fmt.Sprintf("%d", num)
	length := len(str)

	if str[0] == '0' {
		return false
	}

	if length%2 != 0 {
		return false
	}

	patternLen := length / 2

	firstHalf := str[:patternLen]
	secondHalf := str[patternLen:]

	// Check if both halves are identical
	return firstHalf == secondHalf
}

func ProcessInput(start int, end int, counts int) []int {
	var result []int
	for i := start; i <= end; i++ {
		if counts == 2 && Has2RepeatingPattern(i) {
			result = append(result, i)
		}
		if counts > 2 && HasNRepeatingPattern(i) {
			result = append(result, i)
		}
	}
	return result
}

func Do(inputs []string, debug bool, counts int) []int {
	res := make([]int, 0)
	for _, line := range inputs {
		fmt.Println("line:", line)
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			fmt.Println("Invalid range:", line)
			continue
		}
		var start, end int
		fmt.Sscanf(parts[0], "%d", &start)
		fmt.Sscanf(parts[1], "%d", &end)
		numbers := ProcessInput(start, end, counts)
		res = append(res, numbers...)
		if debug {
			fmt.Printf("Processed range %d-%d: %v\n", start, end, numbers)
		} else {
			fmt.Printf("Processed range %d-%d: %d numbers\n", start, end, len(numbers))
		}
	}

	return res
}

func main() {
	inputs, err := common.ReadCommaSeparatedLine("./input.txt")
	if err != nil {
		panic(err)
	}
	res := Do(inputs, true, 3)
	fmt.Println("Final Result:", res)
	final := 0
	for _, v := range res {
		final += v
	}
	fmt.Println("Final Sum:", final)
}

//  11 and 22.
// 95-115 has one invalid ID, 99.
// 998-1012 has one invalid ID, 1010.
// 1188511880-1188511890 has one invalid ID, 1188511885.
// 222220-222224 has one invalid ID, 222222.
// 1698522-1698528 contains no invalid IDs.
// 446443-446449 has one invalid ID, 446446.
// 38593856-38593862 has one invalid ID, 38593859.

// 11 22 99

// 111
// 999
// 1010 1188511885 222222 446446 38593859 565656 824824824 2121212121]
// 4174379265
// 24774350322
