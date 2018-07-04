package main

import (
	"fmt"
)

func main() {
	var x = 10
	fmt.Printf("&x = %p\n",&x)
	pointer(&x)
	fmt.Println(x)
}

func pointer(num *int) {
	*num = 100
	fmt.Printf("\nnum = %p\n",num)
	fmt.Printf("&num = %p\n",&num)
	fmt.Printf("*num = %d\n",*num)
	doublepointer(&num)
}

func doublepointer(num1 **int) {
	**num1 = 1234
	fmt.Printf("\nnum1 = %p\n",num1)
	fmt.Printf("&num1 = %p\n",&num1)
	fmt.Printf("*num1 = %p\n",*num1)
	fmt.Printf("**num1 = %d\n",**num1)
}
