package leetcode

/**
给定一个正整数，返回它在 Excel 表中相对应的列名称。

例如，

    1 -> A
    2 -> B
    3 -> C
    ...
    26 -> Z
    27 -> AA
    28 -> AB

//0-25 => A-Z
//1-26 => A-Z
//27 AA = n/26 = X+Y  if Y = 0  (X-1) Z
//52   => AZ   整除时候  最后一位是 26 Z
//78   => BZ   然后上一位 = 26*3   3 -1 B

//X-Y = Z必定整除 作为前面一位    Y如果是余数就作为后面一位
// Z作为可以被整除的数字  如果小于 26 就可以直接 作为 十位数出现
//

25 Y  26 Z 27 AA => 51  52=> AZ 53 BA
 */
func convertToTitle(n int) string {

	if n <= 26 {
		return string(64 + n)
	}


	y := n % 26
	if y == 0 { // 特殊情况
		return convertToTitle((n - y - 1) / 26) + convertToTitle(26)
	}
	return convertToTitle((n - y) / 26) + convertToTitle(y)
}

