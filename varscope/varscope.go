package main

import (
	 "fmt"
)


//全局变量，如果是小写开头那么就是本包可用，如果是大写起头，包外也可以用
var age int = 100
var Name string = "tom"

//作用于仅限于这个函数的内部，出了之后无法使用，不管变量是大写还是小写
func test()  {
	var age int = 10
	var Name string = "jack"
	fmt.Println(age)
	fmt.Println(Name)
}

func main()  {
	//
	fmt.Println("age=", age)
	fmt.Println("name=",Name)
}