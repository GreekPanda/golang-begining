package main

import (
	"fmt"
	"unsafe"
)

func main()  {

	//4 bytes
	var a float32 = 1.234
	//8 bytes
	var b float64 = 1434343434.121212
	
	fmt.Printf("a = %T, %d, b =%T, %d", a, unsafe.Sizeof(a),  b, unsafe.Sizeof(b))
	fmt.Println()

	//golang default float is float64, it's not followed os
	var c = 123.456
	fmt.Printf("c = %T, %d", c, unsafe.Sizeof(c))

	var d = .52 
	var e = .123
	fmt.Printf("d = ", d, " e = ", e)
	fmt.Println()

	//science cal
	var f = 5.1234e2 // 5.1234*10^2 
	var g = .51234E2
	fmt.Println("f=", f, "g=", g)

	//minus
	var h = 5.1234E-2
	fmt.Println("h=", h)
}