package main

import (
	"fmt"
	"strings"
)

//返回的匿名函数和suffix变量构成了一个闭包
//闭包的好处传入的公共变量一次就够了，而非闭包就是需要每次都需要传入
//闭包的优势在于可以延续上次调用的返回结果，减少了重复的传入
func makeSuffix(suffix string) func (string) string  {
	return func (name string) string {
		if !strings.HasSuffix(name, suffix) {
			//name += suffix
			return name + suffix
		}
		return name
	}
}

//使用闭包的方式来完成关于文件后缀名的
func main()  {
	suffix := ".jpg"
	ret := makeSuffix(suffix)
	name := "spring.jpg"
	fmt.Println(ret(name))
	name1 := "winter.jpg"
	fmt.Println(ret(name1))
	name2 := "hello"
	fmt.Println(ret(name2))
	name3 := "test.avi"
	fmt.Println(ret(name3))

}