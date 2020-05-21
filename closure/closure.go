package main

import (
	"fmt"
)

//必报的说明
func add() func (int ) int {
	//从这部分代码开始，返回的是一个匿名函数，但是函数引用了外部的变量
	//这个匿名函数与外部变量形成了一个整体，构成了闭包
	//闭包是一个类，变量是一个成员，函数一个操作，整个构成了一个闭包
	//n只能初始化一次，如果下次使用时会继续上一个赋值，每调用一次都是一次累计
	//闭包的关键就是分析出返回函数的引用了那些变量
	//函数与变量构成闭包
	var n int = 0
	return func (x int) int  {
		n += x
		return n
	}
}


//对比如果增加了两个变量

func add1() func (int) int {
	var n int = 100
	var s string = "hello"
	return func (x int) int  {
		n += x
		s += string(x)
		fmt.Print(s)
		return n	
	}
	
}

func main()  {
	ret := add()
	fmt.Println(ret(1))
	fmt.Println(ret(100))
	fmt.Println("------------------")
	ret1 := add1()
	fmt.Println(ret1(1))
	fmt.Println(ret1(2))
	fmt.Println(ret1(3))
}