package leetcode

import "fmt"

/**
	给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

	如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

	您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

	示例：

	输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)   9 9 9    1 1 1      1  8       0
	输出：7 -> 0 -> 8						0 1 1 1          0          0
	原因：342 + 465 = 807              999 + 111 =  1110      18 + 0 =18   0+0 = 0
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func AddTwoNumber(){
	a := ListNode{3,nil}
	b := ListNode{4,&a}
	l1 := ListNode{2,&b}

	c := ListNode{4,nil}
	d := ListNode{6,&c}
	l2 := ListNode{5,&d}

	res := addTwoNumber(&l1,&l2)
	fmt.Println(res)
}


func addTwoNumber(l1 *ListNode, l2 *ListNode) *ListNode {
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