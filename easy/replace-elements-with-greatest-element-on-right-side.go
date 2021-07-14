// https://leetcode.com/problems/replace-elements-with-greatest-element-on-right-side/
// Date 14.07.2021
package easy

func replaceElements(arr []int) []int {
	if len(arr) < 1 {
		return arr
	}

	max := arr[len(arr)-1]
	arr[len(arr)-1] = -1

	for i := len(arr) - 2; i >= 0; i-- {
		curr := arr[i]
		arr[i] = max

		if curr > max {
			max = curr
		}
	}
	return arr
}
