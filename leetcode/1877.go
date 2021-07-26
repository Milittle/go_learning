package leetcode

import "sort"

func minPairSum(nums []int) int {
	sort.Ints(nums)
	ret := 0
	for i, j := 0, len(nums)-1; i < j; i++ {
		if nums[i]+nums[j] > ret {
			ret = nums[i] + nums[j]
		}
		j--
	}
	return ret
}
