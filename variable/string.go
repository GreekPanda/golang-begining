package main

import (
	"fmt"
	//"unsafe"
)

func main()  {
	//utf-8 in golang
	var a string = "Shenzhen"
	var b string = "中国"
	fmt.Println(a)
	fmt.Println(b)

	fmt.Printf("%c, %c",b[0], b[1])

	//block string code
	c := `"hello", \what I did I can't know 
	     main() {
			 }`

	fmt.Println(c)

	var d = "hello" + " world"
	var e = d + "got it"
	fmt.Println(e)
}