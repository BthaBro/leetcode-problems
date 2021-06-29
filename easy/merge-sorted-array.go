//https://leetcode.com/problems/merge-sorted-array/
//Date 29.06.21
package easy

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		for i, num := range nums2 {
			nums1[i] = num
		}
	} else if n != 0 {
		i := 0
		j := 0

		for ; i < m+n; i++ {
			if j >= n {
				break
			} else if nums1[i] > nums2[j] {
				for k := m + n - 1; k > i; k-- {
					nums1[k] = nums1[k-1]
				}

				nums1[i] = nums2[j]
				j++
			} else if i >= m+j {
				nums1[i] = nums2[j]
				j++
			}
		}
	}
}
