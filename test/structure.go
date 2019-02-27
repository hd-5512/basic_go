package test

import (
	"fmt"
)

func TestStruct(){
	knowAboutSlice()
	knowAboutRange()
	knowAboutStruct()
}


func knowAboutSlice(){
	/*	array  slice  map  区别
		array := [n]type{...}  数组必须是固定数字的 或者 [..]type{}  长度固定
		slice := []type{...}   和array相似  引用变量  长度不固定 会指向一个底层的array
		map   := map[typekey]typeval{"name":"val"}  字典
		string 的底层也是 []byte 的 array
	 */

	arr := [7]byte{'a','b','c','d','e','f','g'}   //array
	sli := []int {1,3,4,6} 		//slice
	fmt.Println(arr,sli)

	//如下，动态只是一个表象，slice其实是一种引用类型 指向一个底层的array 所以slice的表现形式类似数组但是本质却是一串指针
	var arrA,arrB []byte   //申明两个slice
	arrA = arr[1:5] //b-c-d-e  arr[5]并没有进去 说明右边是一个非封闭集合  不包含
	arrB = arr[4:]  //e-f-g  arr[4] 到最后
	fmt.Println(arrA,arrB)


	/*	slice 具有的方法
		len 获取slice的长度
		cap 获取slice的最大容量
		append 向slice里面追加一个或者多个元素，然后返回一个和slice一样类型的slice
		copy 函数copy从源slice的src中复制元素到目标dst，并且返回复制的元素的个数
		append函数会改变slice所引用的数组的内容，从而影响到引用同一数组的其它slice。
		但当slice中没有剩 余空间(即(cap-len) == 0)时，此时将动态分配新的数组空间。
		返回的slice数组指针将指向这个空间，而原 数组的内容将保持不变;其它引用此数组的slice则不受影响。
	*/
	arrC := arr		//cap == len 所以是一个新的数组空间
	arrA[3] = 'g'   //即 原本A是b,c,d,e => 0，1，2，3 这里把e 变成 g
	fmt.Println(arrA,arrB)   //发现 A 中数据改变后 B也变了  这里可以证明 A、B其实都指向的是同一个地址
	fmt.Println(arr,arrC)    //arr也因为的A的改变而变化  说明A、B指向arr 而C则是另一个数组

	/*	map  key=>value 这样的arr 无序的 每次打印出来的都不同  只能通过key来获取value
		_,ok := map[key]判断是否有 key
	    delete(fromMap,key)
	 */
	rating := map[string]float32 {"C":5, "Go":4.5, "Python":4.5, "C++":2 }

	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}

	delete(rating, "C") // 删除key为C的元素

	m := make(map[string]string)
	m["Hello"] = "Bonjour"
	m1 := m
	m1["Hello"] = "Salut"

	numbers := make(map[string]int)
	numbers["one"] = 1 //赋值 numbers["ten"] = 10 //赋值 numbers["three"] = 3


	/*
		make 和 new 的区别
		new 返回的是 地址
		make 只能用于引用类型  slice channel map
		非常重要的是 map也是只用引用类型 修改map1中的值可能会修改 map2 提供了便利  但也可能出现一些使用不当的危险
	 */
	p := new([]int) //p == nil; with len and cap 0
	fmt.Println(p)

	v := make([]int, 10, 50) // v is initialed with len 10, cap 50
	fmt.Println(v)

	/*********Output****************
		&[]
		[0 0 0 0 0 0 0 0 0 0]
		(*p)[0] = 18        // panic: runtime error: index out of range // because p is a nil pointer, with len and cap 0
		v[1] = 18           // ok
	*********************************/
}
/////////////////////////////////////

func knowAboutRange(){
	/*	for 循环  && for range 遍历slice和map的数据
	 */

	for index := 10; index>0; index-- {
		if index == 5 {
			break // 或者continue
		}
		fmt.Println(index)
	}


	mm := map[string]int{"a":1,"b":2}
	for k,v:=range mm {
		fmt.Println("map's key:",k)
		fmt.Println("map's val:",v)
	}

	//	由于 Go 支持 “多值返回”, 而对于“声明而未被调用”的变量, 编译器会报错, 在这种情况下, 可以使用_来丢弃 不需要的返回值 例如
	for _, v := range mm{
		fmt.Println("map's val:", v)
	}
}



////////////////////////////////////
func knowAboutStruct(){
	var info slice_info  /*结构体类似一个对象object 可以先实例化后根据结构进行赋值或修改 结构体用得好就是在用类*/
	info.numbers = []int {1, 2, 3, 4, 5, 6, 7}

	info2 := new(slice_info)  //这里可以发现  一个结构体使用 new 来实例化的话 是不是真的和类一样了
	info2.numbers = []int{7,8,9}

	info3 := new(slice_info)
	fmt.Println(info,info2,info3)

	slice := info.numbers
	fmt.Println("slice = ", slice)
	fmt.Println(info,info2,info3)
	odd := filter(slice, isOdd) // 函数名输入 => 将函数当做一种数据类型传入filter
	fmt.Println("Odd elements of slice are: ", odd)
	even := filter(slice, isEven) // 函数当做值来传递了
	fmt.Println("Even elements of slice are: ", even)
}



type testInt func(int) bool  //利用结构体可以申明了一种类似function的类型

type slice_info struct {
	numbers []int
	names   []string
	signals [3]bool
}

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}


////////////////////////////////////
