package test

import (
	"errors"
	"fmt"
)

func TestPackage(){
	testInterface()
}



//////////////////////////// interface
type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human //匿名字段Human school string
	loan float32
}

type Employee struct {
	Human //匿名字段Human company string
	money float32
}

//Human对象实现Sayhi方法
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}
// Human对象实现Sing方法
func (h *Human) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}
//Human对象实现Guzzle方法
func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}
// Employee重载Human的Sayhi方法
func (e *Employee) SayHi() {
	//fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone) //Yes you can split into 2 lines here.
}


//Student实现BorrowMoney方法
func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount // (again and again and...)
}
//Employee实现SpendSalary方法
func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount // More vodka please!!! Get me through the day!
}

// 定义interface
type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}
type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}
type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}


















/**
定义接口: 这里定义了一个抽象的东西 叫做 Phone 有一个 call的功能 接口里都是方法 还需要申明方法的返回值类型
 */
type Phone interface {
	call()
	phone_number() string
}

/**
定义一个结构体=>作为抽象接口的实例化对象  在使用接口时  new(StructName)  然后用 . 访问里面的方法
 */
type NokiaPHone struct {

}
func (NokiaPHone) call() {
	fmt.Println("I am Nokia")
}
func (NokiaPHone) phone_number() string {
	return "13900000000"
}


type IPhone struct {

}
func (iPhone IPhone) call() {
	fmt.Println("I am IPhone")
}




type IPhone2 struct {

}
func (IPhone2) call() {
	fmt.Println("I am IPhone")
}
func (IPhone2) phone_number() string {
	return "123123"
}
func (IPhone2) record() {
	fmt.Println("recording")
}

func testInterface() {
	var phone Phone;

	phone = new(NokiaPHone)
	phone.call()
	var n = phone.phone_number()
	fmt.Println(n)

	//phone = new(IPhone)
	//phone.call()   因为iphone 没有按照 interface 定义的规则 必须要有 一个  call 和 mobile_number 它没有  所以无法编译


	phone = new(IPhone2)  //但是可以去定义添加 自己更多的功能  与很多接口的概念都一致  属于父类或者的子集
	phone.phone_number()

}











/**
	//Go的底层有这样的一个方法 用来处理错误

	type error interface {
		Error() string
	}
 */

//这里 只是一个例子  实际更好地使用 https://www.jianshu.com/p/f30da01eea97
func testErro(f float64) (float64, error) {
	if f < 0 {
		return 0,errors.New("wrong number")
	}

	return 0,nil
}

func d() {
	//result,err:= testErro(-1)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(result)

	//if result,errMsg := Divide(100,10);errMsg == "" {
	//	fmt.Println("100/10 = ",result)
	//}
	//
	//if _,errMsg :=Divide(100,0);errMsg != "" {
	//	fmt.Println("errorMsg :",errMsg)
	//}
}


type DivideError struct {
	dividee int
	divider int
}


func (de *DivideError) Error() string {
	strFomat := "sdsdfsfsdf %d asdasd"
	return fmt.Sprintf(strFomat,de.dividee)
}


func Divide(varDividee int,varDivider int ) (result int,errMsg string) {
	if(varDivider == 0) {
		dData :=DivideError{
			dividee:varDivider,
			divider:varDivider,
		}
		errMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider,""
	}
}