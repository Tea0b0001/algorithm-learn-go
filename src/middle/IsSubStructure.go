package main

import "common"

/*
	@title:树的子结构
	@description:输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
	B是A的子结构， 即 A中有出现和B相同的结构和节点值。

	@example:
	例如:
	给定的树 A:
	     3
	    / \
	   4   5
	  / \
	 1   2
	给定的树 B：
	   4 
	  /
	 1
	返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。

	示例 1：
	输入：A = [1,2,3], B = [3,1]
	输出：false
	示例 2：
	输入：A = [3,4,5,1,2], B = [4,1]
	输出：true

	限制：
	0 <= 节点个数 <= 10000

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof
 */

/*
	@version:Go 1.13.12
*/
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

/*
	将B树与A树，A树左子树，A树右子树进行比较，递归进行即可。
 */
func isSubStructure(A *common.TreeNode, B *common.TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return subStructure(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

/*
	判断A树是否包含B树，包含条件是A树的根节点与B树是相同的。
 */
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
