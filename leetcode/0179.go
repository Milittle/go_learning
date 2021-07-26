package leetcode

import (
	"sort"
	"strconv"
	"strings"
)

type byTwoElementLength []string

func (s byTwoElementLength) Len() int {
	return len(s)
}
func (s byTwoElementLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byTwoElementLength) Less(i, j int) bool {
	return s[i]+s[j] > s[j]+s[i]
}

func LargestNumber(nums []int) string {
	m := convertIntToStr(nums)
	sort.Sort(byTwoElementLength(m))
	return strings.Join(m, "")
}

func convertIntToStr(nums []int) []string {
	m := make([]string, len(nums))
	allZero := true
	for idx, value := range nums {
		if value != 0 {
			allZero = false
		}
		m[idx] = strconv.Itoa(value)
	}
	if allZero {
		return []string{"0"}
	}
	return m
}
