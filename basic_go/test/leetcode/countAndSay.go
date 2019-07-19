package leetcode

import (
	"strconv"
)

/**
	报数序列是一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：

1.     1
2.     11
3.     21
4.     1211
5.     111221
6.     312211
1 被读作  "one 1"  ("一个一") , 即 11。
11 被读作 "two 1s" ("两个一"）, 即 21。
21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。

给定一个正整数 n（1 ≤ n ≤ 30），输出报数序列的第 n 项。

注意：整数顺序将表示为一个字符串。
 */
func countAndSay(n int) string {

	if n == 1 {
		return "1"
	}

	var str = "11"
	for i:= 2;i<n;i++{
		temp  := ""
		pre   := str[0]
		count := 1
		for k := 1;k < len(str); k++ {
			if str[k] == pre {
				count++
			} else {
				temp = temp + strconv.Itoa(count) + string(pre)
				count = 1
				pre = str[k]
			}

			if k+1 == len(str){
				temp = temp + strconv.Itoa(count) + string(pre)
			}
		}
		str = temp
	}
	return str
}
