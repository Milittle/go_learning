package tools

func MinOf(vars ...int) int {
	min := vars[0]

	for _, v := range vars[1:] {
		if min > v {
			min = v
		}
	}
	return min
}

func MaxOf(vars ...int) int {
	max := vars[0]

	for _, v := range vars[1:] {
		if max < v {
			max = v
		}
	}
	return max
}
