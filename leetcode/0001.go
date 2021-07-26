package leetcode

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, value := range nums {
		if _, ok := m[target-value]; ok {
			return []int{m[target-value], idx}
		}
		m[value] = idx
	}
	return nil
}
