package main

import (
	"fmt"
	"strconv"
)

func main()  {
	// All basic type has to be converted obviously
	var a int32 = 100
	var b float32 = float32(a)
	var c int64 = int64(b)
	fmt.Println("a= ", a, "b=", b)
	fmt.Printf("a=%T\n", a)
	fmt.Printf("b=%T\n", b)
	fmt.Printf("c=%v\n", c)

	//Golang可以支持范围小的转换成大的，也可以把大的转换成小的，被转换的变量存储的数据，变量本身不改变任何数值

	//在转换中如果从大的范围到小的范围，会有精度丢失
	var d int64 = 12345678901234
	var e int32 = int32(d)
	fmt.Printf("e=%v\n",e)


	var n1 int32 = 100
	var n2 int64
	var n3 int8
	var n4 int8

	n2 = int64(n1)
	n3 = int8(n1) + int8(n2)

	fmt.Printf("n2=%v, n3=%v\n", n2, n3)

	n4 = int8(n1) + 127
	n3 = int8(n1) + 127 //more than 127 is not allowed
	fmt.Printf("n4=%v\n", n4)


	var num5 int64 = 123
	var num6 float64 = 123.123
	var num7 bool = true
	var num8 byte = 'a'

	var str string=""
	var result string = ""

	//Use method1 to concat
	str = fmt.Sprintf("%d", num5)
	fmt.Printf("%s\n", str)
	result += str

	str = fmt.Sprintf("%f", num6)
	fmt.Printf("%s\n", str)
	result += str

	str = fmt.Sprintf("%t", num7)
	fmt.Printf("%s\n", str)
	result += str

	str = fmt.Sprintf("%c", num8)
	fmt.Printf("%s\n", str)
	result += str

	fmt.Printf("%q\n", result)


	//方式二 strconv
	var str1 string = ""
	str1 = strconv.FormatInt(num5, 10)
	fmt.Println(str1)

	str1 = strconv.FormatInt(num5, 2)
	fmt.Println(str1)

	str1 = strconv.FormatInt(num5, 16)
	fmt.Println(str1)
	

	str1 = strconv.FormatFloat(num6, 'f', 2, 64)
	fmt.Println(str1)


	str1 = strconv.FormatBool(num7)
	fmt.Println(str1)


	//str3Int64 convert into default value
	var str3 string = "hello"
	var str3Int64 int64 = 1111
	str3Int64, _ = strconv.ParseInt(str3, 10, 64)
	fmt.Println(str3Int64) 

}