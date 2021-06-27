//https://leetcode.com/problems/max-consecutive-ones/
// Date 27.06.21
package easy

func findNumbers(nums []int) int {
	var counter int

	for _, num := range nums {
		if isEven(num) {
			counter++
		}
	}

	return counter
}

func isEven(num int) bool {
	counter := 0

	for {
		num = num / 10
		counter++

		if num == 0 {
			return counter%2 == 0
		}
	}
}
