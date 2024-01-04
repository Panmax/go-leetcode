package main

// [ 根节点, [左子树的前序遍历结果], [右子树的前序遍历结果] ]
// [ [左子树的中序遍历结果], 根节点, [右子树的中序遍历结果] ]
func buildTree_20240104(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	var i int
	for ; i < len(inorder); i++ {
		if preorder[0] == inorder[i] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
