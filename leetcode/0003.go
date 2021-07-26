package leetcode

import (
	"fmt"
)

func LengthOfLongestSubstring(s string) int {
	left, right, result := 0, -1, 0
	var freq [256]int
	for left < len(s) {
		if right+1 < len(s) {
			fmt.Printf("type is: %T\n", s[right+1])
		}
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]-'a']++
			right++
		} else {
			freq[s[left]-'a']--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}

func max(a int, b int) int {
	if a <= b {
		return b
	}
	return a
}
