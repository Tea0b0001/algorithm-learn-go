package main

/*
	@title:找出数组中重复的数字。
	@description:在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
	数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

	@example:
	示例 ：
	输入：
	[2, 3, 1, 0, 2, 5, 3]
	输出：2 或 3
 
	限制：
	2 <= n <= 100000

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof
 */

/*
	@version:Go 1.13.12
 */
func main() {
	var nums = []int{2, 3, 1, 0, 2, 5, 3}
	var result = findRepeatNumber(nums)
	print(result)
}

func findRepeatNumber(nums []int) int {
	//下标定位法
	return subscriptPosition(nums)
	//集合去重法
	//return setDeduplication(nums)
}

/*
	下标定位法
	·从头到尾遍历数组,当遍历到下标为i的数字时,首先比较这个数字n是否等于下标i,如果等于就遍历下一个数字;如果不是,则将它和下标为n的数字进行比较.
	·如果它和下标为n的数字相等,那么则出现数字重复直接返回;如果不相等,则将它和下标为n的数字进行交换,即把n放到第n个位置上
	·重复这个过程,直到出现一个重复的数字
	·时间复杂度O(n),空间复杂度O(1)
 */
func subscriptPosition(nums []int) int {
	if nums == nil || len(nums) <= 1 {
		return -1
	}

	for i := 0; i < len(nums); i ++ {
		if i != nums[i] && nums[i] == nums[nums[i]] {
			return nums[i]
		}
		swap(nums, i, nums[i])
	}
	return -1
}

func swap(nums []int, i int, j int) {
	var temp int = nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

/*
	集合去重法
	·遍历整个数组,当遍历到的这个数字不存在集合中，将其加入进去；如果在集合中，直接返回即可。
	·时间复杂度O(n),空间复杂度O(n)
 */
func setDeduplication(nums []int) int {
	if nums == nil || len(nums) <= 1 {
		return -1
	}

	var numsMap map[int]int
	numsMap = make(map[int]int)

	for i := 0; i < len(nums); i ++ {
		if val, ok := numsMap[nums[i]]; ok {
			return val
		}
		numsMap[nums[i]] = nums[i]
	}
	return -1
}