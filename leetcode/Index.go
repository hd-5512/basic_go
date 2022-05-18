package leetcode

import (
	"math"
	"math/rand"
)

func Index(){
	//res := lengthOfLongestSubstring("abcbadssddweff")
	//fmt.Println(res)
	//res := Atoi(getStr()) 					//8.字符串转换整数
	//fmt.Println(res)

	workout()
	return
}

func workout(){

	return
	TwoSun() 								//1.两数之和
	AddTwoNumber()							//2.两数相加
	ReverseInt(getInt(0))				//7.整数反转
	searchInsert([]int{1,3,5,6},5) 	//35. 搜索插入位置
	isValidSudoku([][]byte{})      			//36.有效的数独
	countAndSay(6)						//38.报数
	convertToTitle(78)   				//168. Excel表列名称


	//SQL 题目
	zSQL_CombineTwoTables()					//175. 组合两个表
	zSQL_Second_Highest_Salary()			//176. 第二高的薪水
	zSQL_Nth_Highest_Salary()				//177. 第N高的薪水
	zSQL_Rank_Scores()						//178. 分数排名
	zSQL_Consecutive_Numbers()				//180. 连续出现的数字
	zSQL_DuplicateEmails()					//182. 查找重复的电子邮箱

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
	return " d 1234 "
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