package main


import (
	"fmt"
	"../utils/utils"
)

var age int = printAge()

func printAge() int {
	fmt.Println("Runing at test()")
	return 90
}


//It's going to run before main, init ia global inititation
//如果在全局变量有定义，先执行全局变量定义
//然后执行init函数
//然后执行main
//init完成初始化动作

//初始化的工作，如果需要用别的包的文件
func init()  {
	fmt.Println("init function()")

}

func main()  {
	fmt.Println("init main()", age)
}