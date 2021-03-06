package main

import (
	"fmt"
)

//创建切换的方式有四种：
//	1. make([]Type, length, capacity) ：其长度及容量为指定的值，length必须小于等于capacity
//	2. make([]Type, length) ：其长度即为切片的长度
// 	3. []Type{} ：复合语法，其长度为大括号内项的个数
//	4. []Type{value1, value2, ..., valueN} ：复合语法，其长度为大括号内项的个数

//slice 是一个数据结构，其指向一个数组某个连续的部分。slice 用起来很像数组。[]T 为 slice 类型，其中元素类型为 T
func testSlice() {
	p := []int{2, 3, 5, 11}
	fmt.Println("p == ", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
}

//截取切片
func testSlice2() {
	s := []string{"A", "B", "C", "D", "E", "F", "G"} //s其实是一个引用地址
	t := s[:5]
	u := s[3 : len(s)-1]
	//out is : [A B C D E F G] [A B C D E] [D E F]
	fmt.Println(s, t, u)
	u[1] = "X"
	//out is : [A B C D X F G] [A B C D X] [D X F]
	fmt.Println(s, t, u)
}

// make([]Type, length, capacity) ：其长度即为切片的容量
func testMakeSlice1() {
	a := make([]string, 0, 5) //length必须小于等于capacity
	fmt.Println("a : ", a)
	fmt.Println("len : ", len(a))
	fmt.Println("cap : ", cap(a))
}

// make([]Type, length) ：其长度即为切片的长度
func testMakeSlice2() {
	b := make([]string, 0)
	fmt.Println(b)
	fmt.Println("len : ", len(b))
	fmt.Println("cap : ", cap(b))
}

// []Type{} ：复合语法，其长度为大括号内项的个数
func testMakeSlice3() {
	c := []string{}
	fmt.Println(c)
	fmt.Println("len : ", len(c)) // out is 0
	fmt.Println("cap : ", cap(c)) // out is 0
}

// []Type{value1, value2, ..., valueN} ：复合语法，其长度和容量为大括号内项的个数
func testMakeSlice4() {
	d := []string{"Beijing", "New York", "Las Vergas", "California"}
	fmt.Println(d)
	fmt.Println("len : ", len(d))
	fmt.Println("cap : ", cap(d))
}

//测试nil
func testNilSlice() {
	var e []int //e is nil
	fmt.Println(e, len(e), cap(e))
	if e == nil {
		fmt.Println("e is Nil")
	} else {
		fmt.Println("e is not Nil")
	}

	f := []int{} //f is not nil
	fmt.Println(f, len(f), cap(f))
	if f == nil {
		fmt.Println("f is Nil")
	} else {
		fmt.Println("f is not Nil.")
	}
}

/*
func main() {
	//testSlice()
	//testSlice2()
	//testMakeSlice1()
	//testMakeSlice2()
	//testMakeSlice3()
	//testMakeSlice4()
	testNilSlice()
}
*/
