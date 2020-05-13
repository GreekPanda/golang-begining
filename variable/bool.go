package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	var a bool = true
	var b bool = false

	fmt.Println("a=",a, "b=", b)
	fmt.Printf("a=%d", unsafe.Sizeof(a))
	fmt.Println()
}