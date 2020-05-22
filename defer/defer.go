package main

import (
	"fmt"
	_ "encoding/json"
)

//为了在函数中及时释放资源需要增加了延时资源，把指定的资源关闭

func sum(n1, n2 int) int {
	//最先声明，但是最后执行
	//将声明为defer时压入到一个独立的栈中与sum和main不在同一个区域，暂时不执行
	//当函数执行完毕后，再从按照陷入后出的原则，逐个执行声明为defer的出战，然后执行
	defer fmt.Println("ok1 n1 = ", n1)
	defer fmt.Println("ok2 n2 = ", n2)
	ret := n1 + n2
	//没有声明defer的地方，最先执行
	fmt.Println("ok3 ret = ", ret)
	return ret
}

//1.当go遇到关键词defer，不会了立刻执行，而是压力一个独立栈中
//2. 实际使用中，比如打开某个文件，可以将close声明为defer，数据库连接池，也是将close声明为defer
func main()  {
	ret := sum(1, 2)
	//先执行完函数内部，然后是外部返回值执行
	fmt.Println(ret)
}