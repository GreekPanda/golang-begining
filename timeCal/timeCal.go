package main

import (
	"fmt"
	"time"
	"strconv"
)

func funTimeRun() int64 {
	var start int64
	var end int64
	var str string
	//now := time.Now()
	start = time.Now().Unix()
	fmt.Println(start)
	for i := 0; i < 100000; i++ {
		str = str + " " + strconv.Itoa(i)
	}
	end = time.Now().Unix()
	fmt.Println(end)
	return end - start
}

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

	//3 格式化日期和时间
	fmt.Printf("%04d/%02d/%02d %02d:%02d:%02d\n", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	// 将格式化的时间赋值给一个变量
	timeStr := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d\n", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf(timeStr)

	//4 时间常量
	//指定时间的常量
	//sleep时间常量
	//每隔一秒中就打印一个数字，打印到100就退出

	for i := 1; i <= 3; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
		if i > 3 {
			break
		}
	}

	fmt.Println("----------------")

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 100)
		if i > 10 {
			break
		}
	}

	fmt.Println("----------------")
	//unix和unix纳秒的数
	fmt.Println("Unix的时间秒:", now.Unix())
	fmt.Println("Unix时间戳纳秒数：", now.UnixNano())

	fmt.Println("----------------")
	consume := funTimeRun()
	fmt.Println(consume)

}