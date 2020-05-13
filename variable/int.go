package main

import (
	"fmt"
	"unsafe"
) 


func main()  {
	var i int = 100
	fmt.Println("i=", i)

	//int8 ~ [-128, 127]
	var a int8 = 127
	//int16 ~[-2^15, 2^15 -1]
	var a1 int16 = 10000

	//int32 ~[-2^31, 2^31 - 1]
	var b int32 = 1234567
	//int64 ~ [-2^63, 2^64 - 1]
	var c int64 = 555667899991


	fmt.Println("a, b, c, d", a, a1, b, c)

	//how to know what's type it, depends on OS 
	var h int = 1234599888787
	fmt.Printf("h=%T\n", h)

	var j byte = 255
	fmt.Printf("j %T, is byte of %d, h=%T %d, c=%T %d\n", j, unsafe.Sizeof(j), h, unsafe.Sizeof(h),  c, unsafe.Sizeof(c))

	//保小不保大
}