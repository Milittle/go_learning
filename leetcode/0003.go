package leetcode

/*
典型的滑动窗口：
从左到右，如果没有重复，right一直往前走，
如果有重复，则将left往前移动，在每一次移动的过程中，需要记录一下当前是否是最大长度。
变种：
1. 如果不是字符，则可以通过map的方式记录重复，这里用数据进行代替。
2. 如果是返回字符串，则中间记录的是字符串，而不是字符串长度。
 */

func LengthOfLongestSubstring(s string) int {
	left, right, result := 0, -1, 0
	var faq [256]int

	length := len(s)

	for left < length {
		if right + 1 < length && faq[s[right + 1] - 'a'] == 0 {
			faq[s[right + 1] - 'a']++
			right++
		} else {
			faq[s[left] - 'a']--
			left++
		}
		result = max(result, right - left + 1)
	}
	return result
}

func max(a int, b int) int {
	if a <= b {
		return b
	}
	return a
}
