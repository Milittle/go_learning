package leetcode

import "math"

func isValidBST(root *TreeNode) bool {
	return isValidBST1(root, math.Inf(-1), math.Inf(1))
}

func isValidBST1(root *TreeNode, min, max float64) bool {
	if root == nil {
		return true
	}
	v := float64(root.Val)
	return v < max && v > min && isValidBST1(root.Left, min, v) && isValidBST1(root.Right, v, max)
}
