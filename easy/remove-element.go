// https://leetcode.com/problems/remove-element/
// Date 30.06.21
package easy

func removeElement(nums []int, val int) int {
	var (
		count int
	)

	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			if nums[i] == val {
				count++
			}

			continue
		}

		if nums[i] == val {
			for j := i + 1; j < len(nums); j++ {
				nums[j-1] = nums[j]
			}
			count++
		}
	}

	return len(nums) - count
}
