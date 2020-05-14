package main

import (
	"fmt"
)

func test() bool  {
	fmt.Println("test")
	return true
}


func main()  {
	
	var i int = 10

	//The code prove the short circuit

	//no output
	if i < 9 && test() {
		fmt.Println("main logic")
	}

	//output test and main logiccccccc
	if i >= 9 && test() {
		fmt.Println("main logicccccccc")
	}

	//output test and mainnnnnnnn
	if i < 9 || test() {
		fmt.Println("mainnnnnnnn")
	}

	//output mmmmmmmmmmmmmmain, no test
	if i > 9 || test() {
		fmt.Println("mmmmmmmmmmmmmmain")
	}
}