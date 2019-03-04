package main

import (
	"basic_go/test"
)

//import "basic_go/test"
//import 尽量使用绝对路径的写法 如上 其实是在加载  GOPATH/src/myproject/test 模块   GOPATH

// import( . "fmt" )   点引入可以在使用包内函数时不用申明包 fmt.Println("hello world")可以省略的写成Println("hello world")

//import( f "fmt" )    别名  f.Println("hello world")

//import( _ "fmt" )    不使用包内函数 只调用包内的init     !还不理解在什么场景下使用


func main()  {
	test.Test("value")      //go的数据类型以及变量常量
	//test.Test("struct")
	//test.Test("practise")
	//test.TestInterface()

	//test.TestString()
}


//init  和 main 都是go 中的保留函数

//main 只能应用于 package main  而 init 可以应用于所有 package  =》 这意味着main package中 可以同时存在 main 和 init

//然后程序执行的顺序 从package/main进入 最先import => const => var => init => main

//import 的过程中 如果 import的package里面 也是  先 import => const => var => init 如此递归

//上面意味着 main 中的 init 是最后执行的 init 全局变量也最后定义 所以import的包中是无法使用上级文件定义的变量

//最后main执行完毕后 exit