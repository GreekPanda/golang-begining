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

func fib(n int) int  {
	if n == 0 || n == 1 {
		return 1
	} else {
		return fib(n - 1) + fib(n - 2)
	}
		
}

func monkeyEatPeach(n int)  int {
	//猴子吃桃子的第N天的桃子数量: peach = (peach + 1) * 2 
	if n > 10 || n < 1 {
		fmt.Println("error n")
		return -1
	}
	if n == 10 {
		return 1
	} else {
		return (monkeyEatPeach(n + 1) + 1) * 2
	}
}

//递归调用的函数栈调用，出现了else路程本身不影响其流程
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
	fmt.Println("------------------")
	ret := fib(6)
	fmt.Println(ret)
	fmt.Println("------------------")
	ret2 := monkeyEatPeach(1)
	fmt.Println(ret2)
}