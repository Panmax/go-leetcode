package main

func deleteNode_20221019(root *TreeNode, key int) *TreeNode {
	switch {
	case root == nil:
		return nil
	case root.Val > key:
		root.Left = deleteNode_20221019(root.Left, key)
	case root.Val < key:
		root.Right = deleteNode_20221019(root.Right, key)
	case root.Left == nil || root.Right == nil:
		if root.Left != nil {
			return root.Left
		}
		return root.Right
	default:
		successor := root.Right
		// for
		for successor.Left != nil {
			successor = successor.Left
		}
		successor.Right = deleteNode_20221019(root.Right, successor.Val)
		successor.Left = root.Left
		return successor
	}
	return root
}
