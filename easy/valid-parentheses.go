// https://leetcode.com/problems/valid-parentheses/
// Date 13.06.2021
package easy

func isValid(str string) bool {
	var openBrackets []rune

	closedToOpenedMap := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}

	for _, s := range str {
		if s == '(' || s == '[' || s == '{' {
			openBrackets = append(openBrackets, s)
		} else {
			if len(openBrackets) == 0 {
				return false
			}

			if openBrackets[len(openBrackets)-1] != closedToOpenedMap[s] {
				return false
			}

			openBrackets = openBrackets[:len(openBrackets)-1]
		}
	}

	if len(openBrackets) > 0 {
		return false
	}

	return true
}
