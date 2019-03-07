package leetcode

import "fmt"

/*
	给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

	你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

	示例:

	给定 nums = [2, 7, 11, 15], target = 9

	因为 nums[0] + nums[1] = 2 + 7 = 9
	所以返回 [0, 1]
 */

func TwoSun(){
	var nums []int = []int{2,7,11,15}
	var	target int  = 9

	res := twoSum(nums,target)
	fmt.Println(res)

	res2 := twoSum2(nums,target)
	fmt.Println(res2)
}

func twoSum(nums[]int,target int) []int {
	for k1,v1:=range nums {
		for k2,v2 :=range nums {
			if target == v1 + v2 && k1 != k2 {
				return []int{k1,k2}
			}
		}
	}
	return []int{}
}

func twoSum2(nums[]int,target int) []int {
	keys := make(map[int]int,len(nums))

	for k,v:=range nums {
		keys[v] = k
	}

	for k1,v1 :=range nums {
		v2 := target - v1
		if k2,ok := keys[v2]; k1 != k2 && ok {
			return []int{k1, k2}
		}
	}

	return []int{}
}