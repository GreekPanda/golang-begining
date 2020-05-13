package main

import "fmt"


//global variable
//it works, but not good
var g1 = 100
var g2 = "hello"
var g3 = false

//better than above
var (
	g4 = 100
	g5 = "global"
	g6 = 123.123
)

func main()  {
	var i int
	var a float32
	var b float64
	var c bool
	var d int32
	var e int64

	fmt.Println("init value:", i, a , b, c, d, e)

	f := a

	fmt.Println("f is ", f)

	var g = 1234 
	var h = 5678
	fmt.Println("g = ", g)
	fmt.Println("h = ", h)

	var k, l string
	fmt.Println("k = ", k)
	fmt.Println("l = ", l)

	var n1, n2, n3 = 100, "hello", 123.1123
	fmt.Println("n=", n1, " ,n2=", n2, " ,n3=", n3)

	m1, m2, m3 :=100, "test", false
	fmt.Println("m1=", m1, "m2=", m2, "m3=", m3)

	fmt.Println("g1=", g1, "g2=", g2, "g3=", g3)
	fmt.Println("g4=", g4, "g5=", g5, "g6=", g6)

	var t int = 100
	t = 200
	t = 300
	fmt.Println("t=",t)
	//t = 100.0 //Eror, it's not allowed to change the type
}