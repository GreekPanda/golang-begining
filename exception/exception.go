package main

import (
	"fmt"
)

func exception()  {
	n1 := 100
	n2 := 0
	n3 := n1 / n2
	fmt.Println(n3)
}

func exceptionDeal()  {
	defer func ()  {
		//recover是内置函数可以捕获异常
		err := recover()
		if err != nil {
			//如果异常信息不空，说明出现了错误
			fmt.Println(err)
		}
	}()
	n1 := 100
	n2 := 0
	n3 := n1 / n2
	fmt.Println(n3)
}

func main()  {
	//exception()
	//异常出现了之后，不会再继续执行，需要异常的捕获
	//go的异常是defer panic recover
	//使用panic异常，在defer中recover捕获异常并进行处理
	exceptionDeal()
	fmt.Println("Go on running")
}
