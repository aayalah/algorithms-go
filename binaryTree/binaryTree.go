package binarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return formMaxBinaryTree(nums, 0, len(nums))
}

func formMaxBinaryTree(arr []int, s, f int) *TreeNode {
	if s >= f {
		return nil
	}

	max, ind := findMax(arr, s, f)
	node := new(TreeNode)
	node.Val = max
	node.Left = formMaxBinaryTree(arr, s, ind)
	node.Right = formMaxBinaryTree(arr, ind+1, f)
	return node
}

func findMax(arr []int, s, f int) (max int, index int) {
	index = -1

	for i := s; i < f; i += 1 {
		if arr[i] > max || index == -1 {
			max = arr[i]
			index = i
		}
	}
}
