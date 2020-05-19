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
		fmt.Println("error: ", n)
		return 0
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

func test3(n *int)  {
	*n = 20
}

func getSum(n1 int, n2 int)  int {

	return n1 + n2
	
}

func myFun(funvar func(int, int) int, num1 int , num2 int) int {
	return funvar(num1, num2)
}

type funType func(int, int) int

func myFunType(funvar funType, num1 int, num2 int) int {
	return funvar(num1, num2)
}

func getSumSub(n1 int, n2 int) (sum int, sub int) {

	sum = n1 + n2
	sub = n1 - n2
	return
}

func sumArgs(n1 int, args...int) int {
	sum := n1 
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
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
	ret2 := monkeyEatPeach(11)
	fmt.Println(ret2)
	fmt.Println("------------------")
	//使用指针可修改值，其它的基本类型和数组都是传值
	num := 30
	test3(&num)
	fmt.Println(num)

	//go不支持利用参数不同而说明其重载

	fmt.Println("------------------")
	sum := getSum
	//获得函数的内存地址
	fmt.Println(sum)
	//传入参数
	ret3 := sum(10, 20)
	fmt.Println(ret3)

	fmt.Println("------------------")
	//函数可以作为形参传入
	ret4 := myFun(getSum, 50, 60)
	fmt.Println(ret4)

	fmt.Println("------------------")
	//type funType func(int, int) int

	ret5 := myFunType(getSum, 200, 50)
	fmt.Println(ret5)

	fmt.Println("------------------")
	//支持对返回值命名

	ret6, ret7 := getSumSub(100, 50)
	fmt.Printf("sum=%v, sub=%v\n", ret6, ret7)

	fmt.Println("------------------")
	//可变参数
	ret8 := sumArgs(1,2,3,4)
	fmt.Println(ret8)
}