// //https://leetcode.com/problems/longest-common-prefix/
// //Date 4.03.21
package easy

func longestCommonPrefix(strs []string) string {
	var (
		result string
		i      int
	)

	if len(strs) == 0 {
		return result
	}

	for {
		var temp string
		for j := 0; j < len(strs); j++ {
			if len(strs[j]) == i {
				return result
			} else if j == 0 {
				temp = string(strs[j][i])
			} else if temp != string(strs[j][i]) {
				return result
			}
		}

		result += temp
		i += 1
	}
}
