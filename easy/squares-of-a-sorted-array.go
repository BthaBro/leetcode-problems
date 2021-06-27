//https://leetcode.com/problems/max-consecutive-ones/
// Date 27.06.21
package easy

func sortedSquares(nums []int) []int {
	for i, _ := range nums {
		nums[i] = nums[i] * nums[i]

		for j, _ := range nums[:i+1] {
			if nums[j] > nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return nums
}
