package main

import (
	"fmt"
	//"unsafe"
)

func main()  {
	var a byte = 'a'
	var b byte = 'b'
	var c byte = 'C'

	fmt.Println("a=",a, "b=", b, "c=", c)
	fmt.Printf("a=%c, b=%c, c=%c\n",a, b, c)

	//CN is overflow
	// var c byte ='中'
	// fmt.Printf("c=%c\n",c)

	var d int = '国'
	fmt.Printf("d=%c, d=%d", d, d)
	fmt.Println()
}