// https://leetcode.com/problems/valid-mountain-array/
// Date 14.07.2021
package easy

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	increasing := true

	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			return false
		} else if increasing && arr[i] < arr[i-1] {
			if i == 1 {
				return false
			}

			increasing = false
		} else if !increasing && arr[i] > arr[i-1] {
			return false
		}
	}

	return !increasing
}
