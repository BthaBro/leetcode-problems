// https://leetcode.com/problems/remove-element/
// Date 01.07.21
package easy

func removeDuplicates(nums []int) int {
	back := 0
	front := 1

	for front < len(nums) {
		if nums[front] != nums[back] {
			back++
			nums[back] = nums[front]
		}
		front++
	}

	if len(nums) == 0 {
		return 0
	}

	return back + 1
}
