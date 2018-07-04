package main

import "fmt"

func main() {

   	print();

   	var a = 100
   	var b = 200
   	var ret int

   	ret = max(a, b)

	fmt.Printf( "最大值是 %d\n", ret )
	
	x, y := swap("Allen","Charles");

	fmt.Println("交換後:"+x+" "+y);
}

func max(num1, num2 int) int {
   var result int

   if (num1 > num2) {
      result = num1
   } else {
      result = num2
   }
   return result 
}

func print(){
	fmt.Println("就只是個無回傳值的範例");
}

func swap(x, y string) (string, string) {
	return y, x
 }