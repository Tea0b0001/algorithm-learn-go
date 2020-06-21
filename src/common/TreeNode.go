package common

import (
	"container/list"
	"math"
	"strconv"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func Tree2Str(root * TreeNode) string{

	var treeStrSlice []string
	treeStrSlice = make([]string, 0)
	var hasNext bool = false
	var queue = list.New()
	var depth int = 0
	var maxValLen float64 = 0
	queue.PushBack(root)
	if root != nil {
		hasNext = true
	}

	for ;hasNext; {
		hasNext = false
		var width = queue.Len()
		for i := 0; i < width; i ++ {
			e := queue.Front()
			tempTreeNode, ok := e.Value.(* TreeNode)
			if ok {
				if tempTreeNode != nil {
					queue.PushBack(tempTreeNode.Left)
					queue.PushBack(tempTreeNode.Right)
					tempTreeNodeValStr := strconv.Itoa(tempTreeNode.Val)
					treeStrSlice = append(treeStrSlice, tempTreeNodeValStr)
					maxValLen = math.Max(maxValLen, float64(len(tempTreeNodeValStr)))
					if tempTreeNode.Left != nil || tempTreeNode.Right != nil {
						hasNext = true
					}
				} else {
					queue.PushBack(nil)
					queue.PushBack(nil)
					treeStrSlice = append(treeStrSlice, "#")
				}
			} else {
				queue.PushBack(nil)
				queue.PushBack(nil)
				treeStrSlice = append(treeStrSlice, "#")
			}
			queue.Remove(e)
		}
		depth ++
	}

	width := 1
	start := 0
	lineLen := int(maxValLen * math.Pow(2, float64(depth)) * 2)
	unitStr := initSpaceStr(int(maxValLen))
	treeStr := ""
	for i := 0; i < depth; i ++ {
		end := start + width
		truthLine := ""
		for ; start < end; start ++ {
			truthLine += treeStrSlice[start] + unitStr
		}
		addLineLen := int((lineLen - len(truthLine)) / 2)
		addStr := initSpaceStr(addLineLen + 1)
		lineStr := addStr + truthLine + addStr
		treeStr += lineStr + "\n"
		width = width * 2
	}
	return treeStr
}

func initSpaceStr(len int) string {
	spaceStr := ""
	for i := 0; i < int(len); i ++ {
		spaceStr += " "
	}
	return spaceStr
}