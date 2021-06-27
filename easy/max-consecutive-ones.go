//https://leetcode.com/problems/max-consecutive-ones/
//Date 27.06.21
package easy

func findMaxConsecutiveOnes(nums []int) int {
	var result, count int

	for _, num := range nums {
		if num == 1 {
			count++
		} else {
			count = 0
		}

		if count > result {
			result = count
		}
	}

	return result
}
