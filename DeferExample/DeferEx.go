package main

import (
	"fmt"
)

func main() {

	/*defer print(1);
	defer print(2);
	defer print(3);*/
	
	//a()
	
	fmt.Println(y())
}

func print(num int) {
	fmt.Printf("p%d\n", num)
}

func a() {
    i := 0
    defer fmt.Println(i)
    i++
    return
}


func f() (result int) {
	defer func() {
		result++
	}()
	return 100
}

func x() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func y() (r int) {
	defer func(r *int) {
		*r = *r + 5
	}(&r)
	return 1
}
