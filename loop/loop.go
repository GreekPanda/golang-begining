package main

import (
	"fmt"
)

//3个班有5个学生的成绩情况
func grade(class int, student int)  {
	for i := 0; i < class; i++ {
		var sum float64 = 0.0
		for j :=0; j < student; j++ {
			var score float64
			fmt.Printf("输入第%d, 第%d个学生的成绩： \n", i, j)
			fmt.Scanln(&score)
			
			sum += score
		}
		fmt.Printf("第%d班的平均学生成绩为%f \n", i, sum / 5)
	}
}

//打印金字塔
func printPyrimid()  {
	// 1. 打印矩形
	//i标识层数，J 多少个* 
	for i := 0; i < 5; i++ {
		for j :=0; j < 5; j++ {
			fmt.Print("*")
		}	
		fmt.Println()	
	}

	fmt.Println("-----------------------")
	
	//2. 打印半个金字塔
	//*的数量与层数一样，注意第二个的数量要小于等于
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}


	fmt.Println("---------------------")


	//3. 打印这个金字塔
	//    *      1个*=2 * 层数 - 1，1层，2个空格 = 总层数-当前层数
	//   ***     3个*=2 * 层数 - 1，2层，1个空格
	//  *****    5个*=2 * 层数 - 1，3层，0空格

	for i :=0; i < 5; i++ {
		for k := 0; k < 5 - i; k++ {
			//总层数打印当前层数的空格
			fmt.Print(" ")
		}

		for j :=0; j < 2 * i - 1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println("---------------------")

	//4. 打印空心金字塔
	//    *      1个*=2 * 层数 - 1，1层，2个空格 = 总层数-当前层数
	//   ***     3个*=2 * 层数 - 1，2层，1个空格，需要判断是打印* ，还是空格，即每一层的第一个和最后一个是*，其它是空格
	//  *****    5个*=2 * 层数 - 1，3层，0空格
	// 但是最后一层全部是*

	for i :=0; i < 5; i++ {
		for k := 0; k < 5 - i; k++ {
			//总层数打印当前层数的空格
			fmt.Print(" ")
		}

		for j :=0; j < 2 * i - 1; j++ {
			if j == 0 || j == 2 * i - 1 - 1 || i == 4 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}	
		}
		fmt.Println()
	}
}

func nineXnineTable()  {
	//i标识层数
	fmt.Println()
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v X %v  = %v\t", j, i, j * i)
		} 
		fmt.Println()
	}
}

func main()  {
	//grade(3, 5)
	printPyrimid()
	nineXnineTable()
}