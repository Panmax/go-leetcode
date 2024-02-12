package main

func deleteNode_20240212(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Left == nil || root.Right == nil {
		if root.Left == nil {
			return root.Right
		}
		return root.Left
	} else {
		successor := root.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		successor.Right = deleteNode(root.Right, successor.Val)
		successor.Left = root.Left
		return successor
	}
	return root
}
