package main

import (
	"fmt"
)
func main() {
	//justPoint();
	
	nextNumber := getSequence()  
	fmt.Printf("nextNumber address = %p\n",&nextNumber);

   	fmt.Println(nextNumber())
   	fmt.Println(nextNumber())
   	fmt.Println(nextNumber())
   
	nextNumber1 := getSequence()  
	fmt.Printf("nextNumber1 address = %p\n",&nextNumber1);   
   	fmt.Println(nextNumber1())
   	fmt.Println(nextNumber1())
}
func justPoint()  {
	var xx = func(num int) {
		fmt.Printf("%d.Just example\n",num);	
	}

	xx(1);
	xx(2);
	xx(10);
}

func getSequence() func() int {
	i:=0
	return func() int {
		fmt.Printf("%p\n",&i)
	   	i++
	   	return i  
	}
 }