package main

import (
	"fmt"
)

func sum(n1, n2 float32)  float32 {
	return n1 + n2
}

func swap(n1 *int, n2 *int) {

	tmp := *n1
	*n1 = *n2
	*n2 = tmp
}

func main()  {
	fmt.Println(sum(1,2))
	fmt.Println("---------")
	n1 := 10
	n2 := 20go
	swap(&n1, &n2)
	fmt.Printf("n1 = %v, n2 = %v\n", n1, n2)
}