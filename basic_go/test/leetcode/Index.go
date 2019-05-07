package leetcode

import (
	"fmt"
	"math"
	"math/rand"
)

func Index(){
	res := lengthOfLongestSubstring("abcbadssddweff")
	fmt.Println(res)
	//res := Atoi(getStr()) 					//8.字符串转换整数
	//fmt.Println(res)

	workout()
	return
}

func workout(){
	return
	TwoSun() 						//1.两数之和
	AddTwoNumber()					//2.两数相加
	ReverseInt(getInt(0))		//7.整数反转


	//大话数据结构-串匹配
	NormalIndex(getStr(),getStr())
	KMPIndex(getStr(),getStr())

	//大话数据结构-搜索
	TestSequential_Search(getIntArr(),1)
	TestBetterSeq_Search(getIntArr(),1)
	TestBinary_Search(getIntArr(),1)
	TestInsert_Search(getIntArr(),1)
	TestFibonacci_Search(getIntArr(),1)
}




func getStr() string {
	return " d 1234"
}

func getInt(x int) int {
	intArr := getIntArr()

	if x < 0 {
		return intArr[len(intArr)-1]
	}

	x = rand.Intn(len(intArr))
	return intArr[x]
}

func getIntArr() []int {
	intArr := []int{
		123456789,
		0,
		-123,
		987654321,
		-1212,
		111111,
		4567654,
		rand.Intn(math.MaxInt32),
	}
	return intArr
}