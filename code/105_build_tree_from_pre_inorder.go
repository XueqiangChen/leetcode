package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if preorder[0] == inorder[i] {
			break
		}
	}

	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])

	return root
}

// 方法二：更容易理解的版本
func buildTree3(preorder []int, inorder []int) *TreeNode {
	idxMap := map[int]int{}
	for i, v := range inorder {
		idxMap[v] = i
	}
	var helper func(int, int) *TreeNode
	helper = func(inorderLeft int, inorderRight int) *TreeNode {
		// 无剩余节点
		if inorderLeft > inorderRight {
			return nil
		}

		// 先序遍历的首个元素即为当前子树的根节点
		val := preorder[0]
		preorder = preorder[1:]
		root := &TreeNode{Val: val}

		// 根据val在中序遍历中的位置，将中序遍历划分成左右两颗子树
		// 由于我们每次都从先序遍历的开始取节点，所以要先遍历左子树再遍历右子树
		inorderRootIndex := idxMap[val]
		root.Left = helper(inorderLeft, inorderRootIndex-1)
		root.Right = helper(inorderRootIndex+1, inorderRight)
		return root
	}

	return helper(0, len(inorder)-1)
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	tree := buildTree(preorder, inorder)
	fmt.Printf("%+v", tree)
}
