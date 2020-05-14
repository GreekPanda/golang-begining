package main

import (
	"fmt"
)

func main()  {
	var a int = 1 >> 2 // 0 
	var b int = -1 >> 2 // -1
	var c int = 1 << 2 // 4
	var d int = -1 << 2 // -4
	fmt.Println("a=", a, "b=", b, "c=", c, "d=", d)

	fmt.Println(2 & 3) //2
	fmt.Println(2 | 3) // 3
	fmt.Println(13 & 7) // 5
	fmt.Println(5 | 4) // 5

	//负数要先转成补码，正数的原码、反码，补码都是一样，
	//负数的反码就是符号位不变，其它求反，补码就是最后一位加1，异或的运算就是, 不同为1，相同为0
	//有负数之后的异或结果还是补码，继续要还原成原码
	fmt.Println(-3 ^ 3) // -2
	
}