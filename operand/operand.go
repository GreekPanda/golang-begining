package main

import (
	"fmt"
)

func main()  {
	//should cast to float64
	var a float64 = 10.00
	var b int32 = 4
	fmt.Println(a/float64(b))

	var c int64 = 100
	var d int32 = 25
	fmt.Println(c/int64(d))

	var e int64 = 10
	var f int64 = 5
	fmt.Println(e%f)

	//result it int
	var g float64 = 10/4
	fmt.Println(g)

	//it depents on what's value of devide
	var h float64 = 10.0 /4
	fmt.Println(h)
}