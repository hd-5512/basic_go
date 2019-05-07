package leetcode

import (
	"math"
)

//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
func ReverseInt(Number int) int {
	rev := 0
	if Number > 0 {
		max := math.MaxInt32 / 10
		n := Number % 10
		Number = Number / 10
		if n != 0 {
			rev = n
		}

		for n := Number % 10 ; Number >0 && (rev < max || (rev == max && n <= 7)) ; n = Number % 10 {
			Number = Number / 10
			rev = rev * 10 + n
		}

		if Number == 0 {
			return rev
		}
	} else if Number < 0 {
		min := - math.MinInt32 / 10
		n := Number % 10
		Number = - Number / 10
		if n != 0 {
			rev = - n
		}

		for n := Number % 10 ;Number >0 && (rev < min || (rev == min && n <= 8)) ; n = Number % 10 {
			Number = Number / 10
			rev = rev * 10 + n
		}

		if Number == 0 {
			return - rev
		}
	} else {
		return 0
	}

	return 0
}


/*
	判断一个整数是否是回文数。
	回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
	首先 负数不是回文数   个位数为0的回文数只有0
 */
func IsPalindrome(Number int) bool {
	if Number < 0 || (Number % 10 == 0 && Number != 0) {
		return false
	}

	half := 0
	for Number > half {
		half = half * 10 + Number % 10
		Number = Number / 10
	}

	return half == Number || Number == half/10;
}
