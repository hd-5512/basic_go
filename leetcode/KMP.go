package leetcode

import "fmt"

func IndexTest(){
	S := "abdc"

	fmt.Println(S[0])
}


//朴素模式的匹配算法  一个一个比较过去  O(S * T)
func NormalIndex(S string,T string) int {

	i := 0
	j := 0

	for i < len(S) && j < len(T) {
		if S[i] == T[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}

	if j == len(T) {
		return i-len(T)
	}else {
		return 0
	}
}

//KMP模式匹配算法  避免重复的遍历来提高效率
func KMPIndex(S string,T string) int {
	//getTnext("abcdex")    //011111
	//getTnext("abcabx")    //011123
	//getTnext("ababaaaba") //011234223
	//getTnext("aaaaaaaab") //012345678

	i := 0
	j := 0

	next := getTnext(T)

	for i < len(S) && j < len(T) {
		if j== -1 || S[i] == T[j] {
			i++
			j++
		}else{
			j = next[j]
		}
	}

	if j == len(T) {
		return i - len(T)
	} else {
		return 0
	}
}


func getTnext(T string) []int {
	Tlen := len(T)
	next := make([]int,Tlen)
	next[0] = -1

	i:=  0 //T[i] 表示后面的字母位置
	j:= -1 //T[j] 表示前面的字母位置

	for i < Tlen - 1 {
		if j == -1 || T[i] == T[j] {
			j++
			i++
			next[i] = j
		} else {
			j = next[j]    //若字符不相同，则j的位置就要回溯
		}
	}

	fmt.Println("next list is:",T,next) //next 代表 当模型串T在和 参考串S比较过程中的N位如果没有匹配到  可以从这次匹配往后的 next[N]个字符开始比较 减少匹配次数
	return next
}