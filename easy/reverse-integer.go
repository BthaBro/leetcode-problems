// https://leetcode.com/problems/reverse-integer/
// Date 24.01.2021
package easy

import (
	"strconv"
)

func reverse(x int) int {
	var (
		reversedNumber string
		isNegative     bool
	)

	if x%10 == 0 {
		x /= 10
	}

	if x < 0 {
		isNegative = true
		x = -x
	}

	for x != 0 {
		reversedNumber += strconv.Itoa(int(x) % 10)
		x /= 10
	}

	result, _ := strconv.Atoi(reversedNumber)
	if result > 2147483647 {
		return 0
	}

	if isNegative {
		return -result
	}

	return result
}
