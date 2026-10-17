/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}
	curr := root

	for len(stack) > 0 || curr != nil {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		result = append(result, curr.Val)

		curr = curr.Right
	}

	return result
}
