package test

import (
	"fmt"
	"unsafe"
)

/*
 *  变量类型基本为 布尔  数字类型  字符串
 *  派生的复杂类型有   指针  / 数组 切片  Map类型  /  结构体  接口  / Channel  / error / 函数
 */


func TestValue(){
	/*	1bit = 1/0
		8bit 比特 = 1byte 字节  =>  二进制11111111=十进制255 => 一个字节能表示的最大的整数就是255 => unit最大255
		0 - 255被用来表示大小写英文字母、数字和一些符号，这个编码表被称为ASCII编码，比如大写字母A的编码是65，小写字母z的编码是122

		定义变量的方式
		var name type = xxx   var需要申明类型
		name := xxx   这种方式不需要申明类型
		_,ok := 34,true  ,支持同时申明   _用于接收不需要的变量
		申明的变量如果不使用 会报错

		数值类型
	 	使用int  uint(无符号整数) 将默认系统的长度  32或64位
	 	也可以自己去申明   比如  int8(-128 到 127)  uint8(0-255)  float32   complex64(虚数)
	 	但是不同的位数变量之间运算编译时会报错
	*/

	var a int 	= 10
	var b int 	= 12
	var c int8 	= 127 	//如果超过 var f int8 = 128  编译时候会提示 超过范围
	d 	:= b + a   		//g + f 无法计算
	fmt.Print(a,b,c,d)

	/*	字符串
		go中字符串统一统一使用UTF-8编码   go 中使用 ""=>普通字符串表示 以及 ``=>识别编译器中的换行   但不使用''
		还能使用slice切割
 	*/
	e := "简化版定义新变量 但是只能在申明未定义的变量时使用 用来赋值编译会报错"
	var (
		f = "不定义类型自动识别为字符串"
		g = "常用这样的形式来进行批量定义 管理全局变量 也可以不设置值"
	)

	/*	字符串用 + 连接
		单不能随便修改某一部分 除非通过转类型或切片
	*/
	s := "hello"
	p := []byte(s)
	p[0] = 'c'
	s2 := string(p)
	g = "c" + e[1:] + s[:5] // 中文切片会出现乱码 一个中文占3个字节
	var r = []rune(f)    //所以可以用 len([]rune(f)) 来确认包含中文字符串的长度
	g = g + string(r[1:3])  //rune 是特殊的int32 也就是 int   ps:byte=uint8  他们都是转成10进制的正整数
	fmt.Println(e,f,g,s,s2,p,g,r)


	/*	引用变量与指针
		&h 的结果是一个指针  所以l的值是 存储h的地址
		*l  => 用l地址去取d的值  即 d改变 l不变  *l改变
	 */
	n := "h的值"
	l := &n
	m := *l

	o := &n == &m
	fmt.Println(o,&n,&m)  //地址不同
	n = "修改n不会改变m 只会修改*l"
	m = "修改m也不会改变n"
	fmt.Print(l,m,n,o)



	/*	常量 const 定义
		默认定义常量时的第一个值是 0  而第二个常量默认将会等于上一个定义常量的值 所以需要注意 常量最好都加上自己的默认值 避免错误

		go里面有一个关键字iota，这个关键字用来声明enum枚举的时候采用，它默认开始值是0，const里每调用一次加1
		同理当 const中使用 iota 后 之后的值默认都是iota +1
		iota 在每次调用 const 时会重新归零
	 */

	const WIDTH int = 5
	const (
		user = "12"
		user_len = len(user)  /*len 需要对 string 进行*/
		//user_cap = cap(user)  //cap是看一个变量的最大容量 如一个slice的cap是它指向array的lenth 如果超过 则会重新指向一个行的array地址 不影响数据
	)

	var x = unsafe.Sizeof(user)  /* Golang中的sring内部实现由两部分组成，一部分是指向字符串起始地址的指针，另一部分是字符串的长度，两部分各是8字节 */

	fmt.Println(x,user_len,"Hello, World!") //定义的变量必须要输出或使用 不然就会报错
	//不知道为啥 Debug中 只能看 变量  常量需要赋值 后才能显示




	/*	array、slice、map、range 具体在 structure 中了解
	 */


}
