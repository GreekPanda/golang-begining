package main

import (
	"fmt"
)

func swap(a *int, b *int) {
	var t int = *a
	*a = *b
	*b = t
}

func swap1(i int ,j int) {
	var t int = i
	i = j
	j = t
}

func main()  {

	//The basic simaple swap way
	var a int = 9
	var b int = 100
	var t int = a
	a = b
	b = t
	fmt.Println("a=", a, "b=", b)

	//no effect
	c := 200
	d := 300
	swap1(c, d)
	fmt.Println("after swap a=", c, "b=", d)

	// var e int = 100
	// var f int = 400

	//It's same as above
	e := 100
	f := 400
	swap(&e, &f)
	fmt.Println("After swap with pointer e = ", e, "f = ", f)


}