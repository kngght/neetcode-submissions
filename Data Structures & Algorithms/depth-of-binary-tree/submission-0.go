/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	
	type pair struct {
		node *TreeNode
		depth int
	}
	stack  := []pair{pair{node:root, depth: 1}}
	result := 0

	for len(stack) > 0 {
		p := stack[len(stack) - 1]
		result = max(result, p.depth)
		stack = stack[:len(stack) - 1]

		if p.node.Left != nil {
			stack = append(stack, pair{
				node: p.node.Left,
				depth: p.depth + 1,
			})
		}
		if p.node.Right != nil {
			stack = append(stack, pair{
				node: p.node.Right,
				depth: p.depth + 1,
			})
		}
	}

	return result
}
