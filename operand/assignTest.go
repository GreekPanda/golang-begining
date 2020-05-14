package main

import (
	"fmt"
)

func main()  {
	
	//不允许使用中间变量完成交换
	a := 100
	b := 200

	a = a + b
	b = a - b // a + b - b = a 
	a = a - b // a - (a - b) = b
	fmt.Println("a=", a , "b=", b)

	a = a ^ b
	b = b ^ a
	a = a ^ b

	fmt.Println("a=", a , "b=", b)
}