package test

import "fmt"

func TestPractise(){
	//for循环和if...else 结构语法联系
	ShowYangHuiTriangle1(10)
	ShowYangHuiTriangle2(10)
	TestGoto()
	TestChannel()

}


func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c<-sum //sendsumtoc
}

func TestChannel() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x + y)
}

func TestGoto(){
	var a int = 100
	var b int = 200
	var ret int

	Compare:
		ret = larger_number(a, b) /* 比大小 */

	fmt.Printf( "最大值是 : %d\n", ret )

	if a != ret {
		a = a+100     //直到a比b大为止
		goto Compare  //跳到申明的位置去执行代码
	}
}


/* 函数返回两个数的最大值  函数声明告诉了编译器函数的名称，返回类型，和参数。 */
func larger_number(num1, num2 int) int {
	var result int
	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}


/* 杨辉三角形1 */
func ShowYangHuiTriangle1(lines int)  {
	nums := []int{}  //  var nums = []
	for i := 0; i < lines; i++ {
		//补空白
		for j := 0; j < (lines - i); j++ {
			fmt.Print(" ")
		}
		for j := 0; j < (i + 1); j++ {
			var length = len(nums)
			var value int
			if j == 0 || j == i {
				value = 1
			} else {
				value = nums[length-i] + nums[length-i-1]
			}
			nums = append(nums, value)
			fmt.Print(value, " ")
		}
		fmt.Println("")
	}
}

/* 杨辉三角形2 */
func ShowYangHuiTriangle2(lines int){
	line_nums := []int{}
	var line int
	var arrLen int

	last_nums := []int{}
	for line = 0; line < lines; line++ {
		var i int
		last_nums = line_nums
		arrLen = len(last_nums)
		line_nums = []int{}

		line_nums = append(line_nums,1)
		if(arrLen == 0) {
			fmt.Println(line_nums)
		}

		for i = 0; i < arrLen-1; i++ {
			line_nums = append(line_nums,last_nums[i]+last_nums[i+1])
		}

		line_nums = append(line_nums,1)

		fmt.Println(line_nums)
	}
}


/**
九九乘法表
 */









/**
约瑟夫环
 */








/**
快速排序 冒泡
 */