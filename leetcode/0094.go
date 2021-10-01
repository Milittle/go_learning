package leetcode

func InorderTraversal(root *TreeNode) []int {
	var result []int
	inorderTree(root, &result)
	return result
}

func inorderTree(root *TreeNode, result *[]int) {
	if root != nil {
		inorderTree(root.Left, result)
		*result = append(*result, root.Val)
		inorderTree(root.Right, result)
	}
}
