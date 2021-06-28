//https://leetcode.com/problems/max-consecutive-ones/
// Date 28.06.21
package easy

func duplicateZeros(arr []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 && i != len(arr)-1 {
			for j := len(arr) - 2; j > i; j-- {
				arr[j+1] = arr[j]
			}

			arr[i+1] = 0
			i++
		}
	}
}
