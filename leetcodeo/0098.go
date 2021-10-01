package leetcodeo

import "sort"

type IntSlice []int

func (a IntSlice) Len() int {
	return len(a)
}

func (a IntSlice) Less(i, j int) bool {
	return a[i] <= a[j]
}

func (a IntSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func isValidBST(root *TreeNode) bool {
	res := inorderTraversal(root)
	return sort.IsSorted(IntSlice(res))
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	inorder(root, &result)
	return result
}

func inorder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, result)
	*result = append(*result, root.Val)
	inorder(root.Right, result)
}
