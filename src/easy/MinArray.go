package main

/*
	@title:旋转数组的最小数字
	@description:把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。
	例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。  

	@example:
	示例 1：
	输入：[3,4,5,1,2]
	输出：1
	示例 2：
	输入：[2,2,2,0,1]
	输出：0


	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof
 */

/*
	@version:Go 1.13.12
*/

func main() {
	var numbers = []int{3, 4, 5, 1, 2, 3}
	print("min number:", minArray(numbers))
}

/*
	循环二分： 设置left, right指针分别指向 numbers 数组左右两端，mid=(left+right)/2 为每次二分的中点，可分为以下三种情况：
	1.当numbers[right] < numbers[mid]时，最小数字会在[mid + 1, right]的区间内；
	2.当numbers[right] > numbers[mid]时，最小数字会在[left, mid]的区间内；
	3.当numbers[right] == numbers[mid]时，最小数字会在[left, right - 1]的区间内。

	思考：
	1.为什么使用right指针和中点mid的值做比较？
	假设按照旋转点的位置把旋转数组分为两个顺序部分的话，如果使用left，当numbers[left] < numbers[mid]时，无法判断mid在哪个顺序部分内。
	本质上是因为，right是一定会在右顺序部分，而left则不确定。
	比如说，[1,2,3,4,5],[3,4,5,1,2]这两个数组，可以看到第一个数组中mid在右顺序部分，第二个数组中mid在左顺序部分。
	2.为什么当numbers[right] == numbers[mid]时，最小数字会在[left, right - 1]的区间内？
	当这种情况往往会出现在这几类特殊情形，比如[1,0,1,1,1]，[1,1,1,0,1]，把区间right减小到right-1均能解决这个问题。
*/
func minArray(numbers []int) int {

	if numbers == nil || len(numbers) == 0 {
		return -1
	}
	var left int = 0
	var right int = len(numbers) - 1

	for ; left < right; {
		var mid int = (left + right) / 2
		if numbers[right] < numbers[mid] {
			left = mid + 1
		} else if numbers[right] > numbers[mid] {
			right = mid
		} else {
			right --
		}
	}
	return numbers[left]
}
