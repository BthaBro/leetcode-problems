// https://leetcode.com/problems/check-if-n-and-its-double-exist/
// Date 14.07.2021
package easy

func checkIfExist(arr []int) bool {
	values := make(map[int]struct{})

	for _, val := range arr {
		if _, ok := values[val*2]; ok {
			return true
		}

		if _, ok := values[val/2]; ok && val%2 == 0 {
			return true
		}

		values[val] = struct{}{}
	}

	return false
}
