package test

func Test(testRoot string){
	switch testRoot {
		case "value" : TestValue()  	//变量相关知识
		case "struct": TestStruct()     //结构体和函数相关知识
		case "practise":TestPractise()  //函数以及练习
		case "package":TestPackage()   //Go  defer  channel  等特性
		case "form":TestForm()       //表单GET POST AJAX 提交等
		default:
			print("invailed param")
	}
}