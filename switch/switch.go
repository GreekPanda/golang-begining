package main

import (
	"fmt"
)

func upperCase(c byte) {
	switch c {
	case 'a':
		fmt.Println("A")
	case 'b':
		fmt.Println("B")
	default:
		fmt.Println("other")
	}
}

//student grade
func grade(g int) {
	switch  {
	case g <= 59 && g > 0:
		fmt.Println("Fail")		
	case g < 80 && g >= 60:
		fmt.Println("pass")
	case g >= 80 && g < 100:
		fmt.Println("Excelent")
	case g == 100:
		fmt.Println("Full")
	default:
		fmt.Println("Wrong")

	}
}

func season(month int) {
	switch month {
	case 3, 4,5:
		fmt.Println("Spring")
	case 6,7,8:
		fmt.Println("Summer")
	case 9, 10, 11:
		fmt.Println("Fall")
	case 12,1,2:
		fmt.Println("Winter")
	default:
		fmt.Println("Error")
	}
}

func main()  {
	var a byte = 'a'
	upperCase(a)

	var g int = -1
	grade(g)

	var month int = 3
	season(month)
}