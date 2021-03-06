package leetcode

/**
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。

示例 1:

输入: [1,3,5,6], 5
输出: 2

输入: [1,3,5,6], 2
输出: 1
 */
func searchInsert(nums []int, target int) int {
	for key,val := range nums {
		if val == target {
			return key
		} else if val > target {
			return key
		}
	}

	return len(nums)
}