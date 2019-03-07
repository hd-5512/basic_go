package test

import (
	"fmt"
	"strconv"
)

type Human2 struct {
	name string
	age int
	phone string
}

// 通过这个方法 Human 实现了 fmt.Stringer
func (h Human2) String() string {
	return " "+h.name+" - "+strconv.Itoa(h.age)+" years - ✆ " +h.phone+" "
}

func TestInterface() {
	Bob := Human2{"Bob", 39, "000-7777-XXX"}
	fmt.Println("This Human is : ", Bob)
}
