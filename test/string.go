package test

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func TestString(){
	//Trim
	fmt.Println(strings.Trim(" !!! Achtung !!! ", "! "))
	fmt.Printf("[%q]", strings.Trim(" !!! Achtung !!! ", "! "))

	//Join 合并
	s1:= []string{"a","b","c"}
	fmt.Println(strings.Join(s1,","))    //a,b,c
	//Split 分割
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))


	//Fields 去除两边的空格 并按照空格分割
	fmt.Printf("Fields are: %q", strings.Fields(" foo bar baz "))


	//长度 Count 获取字符串中另外某一个字符串出现的次数 如果是空格 所以就等于是+1个字符串了
	s2:= "aa,b,caaa你"
	fmt.Println(strings.Count(s2,""))
	fmt.Println(strings.Count(s2,"aa"))
	fmt.Println(utf8.RuneCountInString(s2))  //这个方法在 Count里面 可以直接计算uft8字符的长度
	fmt.Println(len(s2))    //说明中文在go中字符长度占两个

	//查找字符串的位置 没有返回 -1
	fmt.Println(strings.Index("chicken", "ken"))   //4
	fmt.Println(strings.Index("chicken", "dmr"))   //-1

	//查找包含字符
	fmt.Println(strings.Contains("Contains","Con"))  //true
	fmt.Println(strings.Contains("Contains","Conn"))  //false
	fmt.Println(strings.Contains("Contains","con"))  //false
	fmt.Println(strings.Contains("Contains",""))  //true
	fmt.Println(strings.Contains("",""))         //true

	//replace 替换  最后的值表示替换执行的次数   -1 表示全部替换
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))



	//strconv字符串转换库
	//Append 将不同的类型转成string添加到现有的字节里
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))


	//Format  把其他类型的转换为字符串
	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)


	//Parse 将字符串转位其他类型
	a1, err := strconv.ParseBool("false")
	if err != nil {
		fmt.Println(err)
	}
	b1, err := strconv.ParseFloat("123.23", 64)
	if err != nil {
		fmt.Println(err)
	}
	c1, err := strconv.ParseInt("1234", 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	d1, err := strconv.ParseUint("12345", 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	e2:= strconv.Itoa(1023)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a1, b1, c1, d1, e2)
}