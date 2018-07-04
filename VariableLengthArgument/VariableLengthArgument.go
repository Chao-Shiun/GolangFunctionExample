package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1,2,3,4,5,6,7,8,9));
}

func sum (args ...int) int{
	sum := 0

	for _,e:=range args{
		sum += e
	}
	return sum
}