package main

import (
	"fmt"
)

//全局匿名函数必须是大写，默认都是课件才能使用
var (
	Fun1 = func(n1, n2 int) int{
		return n1 * n2
	}
)

func main()  {
	//求和使用匿名函数
	//方法1
	ret1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println(ret1)

	//方法2
	sub := func(n1 int, n2 int) int {
		return n1 - n2
	}

	ret2 := sub(20, 10)
	fmt.Println(ret2)

	//方法3 ，全局匿名函数

	ret3 := Fun1(10, 2)
	fmt.Println(ret3)
}