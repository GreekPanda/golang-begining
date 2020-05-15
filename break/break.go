package main

import (
	"fmt"
)

//求1~100以内的数，如果和等于20就打印出来

func sum() {
	var sum int = 0
	for i := 1; i <=100; i++ {
		sum += i
		if sum > 20 {
			fmt.Printf("sum=%d, i=%d\n", sum, i)
			break;
		}
	}
}

func login() {
	var name string
	var pwd string
	var chance byte = 3

	for i := 0; i < 3; i++ {
		fmt.Println("输入用户名:")
		fmt.Scanln(&name)
		fmt.Println("输入密码")
		fmt.Scanln(&pwd)

		if name == "hello" && pwd == "123" {
			fmt.Println("login successfully")
			break
		} else {
			chance--
			fmt.Println("还剩几次机会:", chance)
		}	
	}

	if chance == 0 {
		fmt.Println("3次用户名或者密码输入错误")
	}
}



func main()  {
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				//break 跳出最近的循环
				//break 可以指定标签，跳到指定的标签
				break
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}


	fmt.Println("---------------------------")

	label:
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				//break 跳出最近的循环
				//break 可以指定标签，跳到指定的标签
				break label
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}

	fmt.Println("-----------------------")

	sum()

	fmt.Println("--------------------")
	login()
}