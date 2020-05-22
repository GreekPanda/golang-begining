package main

import (
	"fmt"
	"time"
)

func main()  {

	//1.获取当前时间
	now := time.Now()
	fmt.Printf("%v: type= %T\n", now,  now)
	// tm := time.Time
	//2. 获取年月日，时分秒
	fmt.Println("----------------------------")
	fmt.Printf("year=%v, month=%v, day=%v\n", now.Year(), now.Month(), now.Day())
	fmt.Printf("hour=%v, minute=%v, second=%v\n", now.Hour(), now.Minute(), now.Second())
	//2.1 将月份修改为数字
	fmt.Printf("year=%v, month=%v, day=%v\n", now.Year(), int(now.Month()), now.Day())
}