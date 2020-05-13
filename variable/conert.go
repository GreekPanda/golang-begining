package main

import (
	"fmt"
)

func main()  {
	// All basic type has to be converted obviously
	var a int32 = 100
	var b float32 = float32(a)
	var c int64 = int64(b)
	fmt.Println("a= ", a, "b=", b)
	fmt.Printf("a=%T\n", a)
	fmt.Printf("b=%T\n", b)
	fmt.Printf("c=%v\n", c)

	//Golang可以支持范围小的转换成大的，也可以把大的转换成小的，被转换的变量存储的数据，变量本身不改变任何数值

	//在转换中如果从大的范围到小的范围，会有精度丢失
	var d int64 = 12345678901234
	var e int32 = int32(d)
	fmt.Printf("e=%v\n",e)


	var n1 int32 = 100
	var n2 int64
	var n3 int8
	var n4 int8

	n2 = int64(n1)
	n3 = int8(n1) + int8(n2)

	fmt.Printf("n2=%v, n3=%v\n", n2, n3)

	n4 = int8(n1) + 127
	n3 = int8(n1) + 127
	fmt.Printf("n4=%v\n", n4)

}