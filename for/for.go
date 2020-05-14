package main

import (
	"fmt"
)

func main()  {
	var str string = "hello world"

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c ", str[i])
	}

	fmt.Println()
	for i, val := range str {
		fmt.Printf("i=%d, val=%c \t", i, val)
	}
	fmt.Println()

	//golang默认是utf-8,一个汉字占3个字节
	var cn string = "中国"
	//将字符串转换成切片
	dst := []rune(cn)
	for i := 0; i < len(dst); i++ {
		fmt.Printf("%c", dst[i])
	}
	fmt.Println()

	//非英文字符串的输出推荐使用for-range的方式，也就是先使用切片rune，然后再进行循环展示
	for _, val := range dst {
		fmt.Printf("%c", val)
	}
	fmt.Println()
}