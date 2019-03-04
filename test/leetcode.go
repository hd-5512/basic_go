package test

/*
	Q1:给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

	你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

	示例:

	给定 nums = [2, 7, 11, 15], target = 9

	因为 nums[0] + nums[1] = 2 + 7 = 9
	所以返回 [0, 1]
 */
func Q1s1(nums[]int,target int) []int {
	for k1,v1:=range nums {
		for k2,v2 :=range nums {
			if target == v1 + v2 && k1 != k2 {
				return []int{k1,k2}
			}
		}
	}

	return []int{}
}

func Q1s2(nums[]int,target int) []int {
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


func Q1s3(nums[]int,target int) []int {
	return []int{}
}



/**
	Q2:给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

	如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

	您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

	示例：

	输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)   9 9 9    1 1 1
	输出：7 -> 0 -> 8						0 1 1 1
	原因：342 + 465 = 807              999 + 111 =  1110
 */

type ListNode struct {
	Val int
	Next *ListNode
}

//a := test.ListNode{9,nil}
//b := test.ListNode{4,&a}
//l1 := test.ListNode{9,&b}

//c := test.ListNode{4,nil}
//d := test.ListNode{5,&c}
//l2 := test.ListNode{5,&d}

//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)   9 9 9    1 1 1
//输出：7 -> 0 -> 8						0 1 1 1         943  564   4 1 8
//原因：342 + 465 = 807              999 + 111 =  1110    349+465 =   814
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		v1 int
		v2 int
		v3 int
		vd int = 0   //进位
		//end1 int = 0
		//end2 int = 0
	)

	head := new(ListNode)
	cur := head
	for true {
		if l1 == nil {
			v1 = 0
		}else{
			v1 = l1.Val
		}

		if l2==nil {
			v2 = 0
		} else {
			v2 = l2.Val
		}

		v3 = v1 + v2 + vd
		if(v3 >= 10){
			v3 = v3 - 10;
			vd = 1
		}else{
			vd = 0
		}

		if (v3>0 || vd>0) || !(l1 == nil && l2==nil) {
			tail := new(ListNode)
			tail.Val = v3
			p := &tail
			cur.Next = *p
			cur = tail
		}

		if(vd == 0 && l1 == nil && l2==nil) {
			if head.Next != nil {
				return head.Next
			} else {
				return head
			}
		}

		if(l1 != nil){
			l1 = l1.Next
		}

		if(l2 != nil){
			l2 = l2.Next
		}
	}
	return head
}