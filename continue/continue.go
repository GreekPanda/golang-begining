package main

import (
	"fmt"
)

//打印1~100之间的奇数

func even()  {
	for i := 1; i < 100; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Print(i)
		fmt.Printf("\t")
	}
}

func main()  {
	for i := 0; i < 4; i++ {
		for j :=0; j < 4; j++ {
			if j == 2 {
				//if j = 2, no more run of the next
				continue
			}
			fmt.Println(j)
		}
	}

	fmt.Println("--------------")

	here:
	for i := 0; i < 4; i++ {
		for j :=1; j < 4; j++ {
			if j == 3 {
				continue here
			}
			fmt.Println(j)
		}
	}

	fmt.Println("----------------")
	even()
	
}