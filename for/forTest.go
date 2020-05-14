package main

import (
	"fmt"
)

//打印1·100之间所有9的倍数的和

func sum9() int {
	var sum int = 0

	for i := 1; i < 100; i++ {
		if i % 9 == 0 {
			sum += i
			fmt.Printf(" %d ", i)
		}
	}
	return sum
}


func main()  {
	fmt.Println(sum9())
}