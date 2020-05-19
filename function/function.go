package main

import (
	"fmt"
)

//递归的流程的，每一此的调用都会有栈的生成，所以每次调用都会相应的输出，直到条件推出
func test(n int)  {
	if n > 2 {
		n--;
		test(n)
	}
	fmt.Println("output n=", n)
}

func test2(n int)  {
	if n > 2 {
		n--
		test2(n)
	} else {
		fmt.Println("output2 n = ", n)
	}
}

func main()  {
	var n int = 4
	test(n)
	fmt.Println("------------------")
	test2(n)
}