package main

import "common"

func main() {
	treeNode1 := common.TreeNode{
		Val: 1,
	}
	treeNode2 := common.TreeNode{
		Val: 2,
	}
	treeNode3 := common.TreeNode{
		Val: 3,
	}
	treeNode4 := common.TreeNode{
		Val: 4,
	}
	treeNode5 := common.TreeNode{
		Val: 5,
	}
	treeNode1.Left = &treeNode2
	treeNode1.Right = &treeNode3
	treeNode2.Left = &treeNode4
	treeNode3.Left = &treeNode5

	testTreeNode1 := common.TreeNode{
		Val: 1,
	}
	testTreeNode2 := common.TreeNode{
		Val: 2,
	}
	testTreeNode1.Left = &testTreeNode2
	println("树A结构:( # 表示空节点)")
	println(common.Tree2Str(&treeNode1))
	println("树B结构:( # 表示空节点)")
	println(common.Tree2Str(&testTreeNode1))
	print("判断结果：",isSubStructure(&treeNode1, &testTreeNode1), " （true表示B是A的子结构，false表示B不是A的子结构）")
}


func isSubStructure(A *common.TreeNode, B *common.TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return subStructure(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func subStructure(A *common.TreeNode, B *common.TreeNode) bool {

	if B == nil {
		return true
	}
	if A == nil {
		return false
	}

	if A.Val != B.Val {
		return false
	}

	return subStructure(A.Left, B.Left) && subStructure(A.Right, B.Right)
}
