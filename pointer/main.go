package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	//get memory address
	var a int32 = 10
	fmt.Printf("%p\n", &a)

	//b is a pointer to point a's address
	var b *int32 = &a
	fmt.Printf("%v\n", b)

	var c int32 = *b
	fmt.Printf("%v, %T\n",c, unsafe.Sizeof(*b))

	//通过指针变量赋值，影响原值
	*b = 100
	fmt.Printf("%v, %T\n",a, unsafe.Sizeof(*b))
}