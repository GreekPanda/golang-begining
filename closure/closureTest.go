package main

import (
	"fmt"
	"strings"
)

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