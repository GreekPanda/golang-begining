package main

import (
	"fmt"
	//"unsafe"
)

func main()  {
	var a byte = 'a'
	var b byte = 'b'

	fmt.Println("a=",a, "b=", b)
	fmt.Printf("a=%c, b=%c\n",a, b)

	//CN is overflow
	// var c byte ='中'
	// fmt.Printf("c=%c\n",c)

	var d int = '国'
	fmt.Printf("d=%c, d=%d", d, d)
	fmt.Println()
}