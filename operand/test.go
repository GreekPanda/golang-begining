package main

import (
	"fmt"
)


//如果放假还有97天，请问还有几个星期几天

func  main()  {
	var left int = 97
	var weeks int = left / 7
	var days int = left % 7

	fmt.Println("weeks = ", weeks, " days = ", days)

	//temperature is converted
	var f float32 = 100
	var c float32 = 5.0 / 9 * (f - 100)
	fmt.Println(c)
}