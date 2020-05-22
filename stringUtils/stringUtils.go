package main

import (
	"fmt"
	"strconv"
	"strings"
)

//常用string内置函数
func main()  {
	str := "hello"
	str1 := "中"
	//a := {1,2,3}
	
	//len是内置函数
	fmt.Println(len(str))
	//golang中统一为utf-8,中文是3个字节，
	//字母、数字，ASCII字符一个字节
	//len按字节返回 
	fmt.Println(len(str1))
	//fmt.Println(len(a))

	fmt.Println("------------")

	str3 := "hello中国"

	//这样打印会出现乱码,因为用字节输出
	for i := 0; i < len(str3); i++ {
		fmt.Printf("%c\n",str3[i])
	}
	
	//如果有中文字符必须要使用切片，这样才能正常输出
	r := []rune(str3)

	for i := 0; i < len(r); i++ {
		fmt.Printf("%c\n", r[i])
	}

	fmt.Println("------------")

	str4 := "123"
	n1 ,err := strconv.Atoi(str4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n1)
	}

	str5 := "hello"
	n2 ,err := strconv.Atoi(str5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n2)
	}

	fmt.Println("------------")
	//整数转字符串
	n3 := 12345
	fmt.Println(strconv.Itoa(n3))

	//字符串转成byte切片
	fmt.Println("------------")
	str6 := "hello, str to byte"
	var bytes = []byte(str6)
	//输出ASCII
	fmt.Println(bytes)

	str7 := string([]byte{97, 98, 99})
	fmt.Println(str7)

	fmt.Println("------------")

	//2、8、16进制的转换,返回的对应的字符串
	fmt.Printf("%v\n", strconv.FormatInt(123,2))
	fmt.Printf("%v\n", strconv.FormatInt(123,8))
	fmt.Printf("%v\n", strconv.FormatInt(123,16))

	fmt.Println("------------")
	//查找字符串中的子串
	fmt.Printf("%v\n", strings.Contains("hello", "ll"))

	fmt.Println("------------")
	//统计一个字符串中有几个子串

	fmt.Printf("%v\n", strings.Count("hellllllllo", "l"))

	fmt.Println("------------")
	//字符串比较
	//等号区分大小写，如果不想区分大小写
	//str8 := "abc"
	//str9 := "ABC"

	//true，不区分大小写
	fmt.Println(strings.EqualFold("abc", "ABC"))
	//false,区分大小写
	fmt.Println("abc" == "ABC")

	fmt.Println("------------")
	//返回子串在串中出现的第一个的下标
	index := strings.Index("HELLO_abc", "abc")
	index1 := strings.Index("HELLO_Abc", "abc")
	index2 := strings.Index("HELLO_abcabcabc", "abc")
	fmt.Println(index)
	//区分大小写，如果不存在，就是返回-1
	fmt.Println(index1)
	fmt.Println(index2)



}