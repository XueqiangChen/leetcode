package main

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return root
	}

	l := 0
	for i := 0; i < len(preorder); i++ {
		if postorder[i] == preorder[1] {
			l = i + 1
		}
	}

	root.Left = constructFromPrePost(preorder[1:l+1], postorder[:l])
	root.Right = constructFromPrePost(preorder[l+1:], postorder[l:])

	return root
}
