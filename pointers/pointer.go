package main

import "fmt"

func number(num int){
     num=5                       // here num is passed by value so chnage here it just change the copy so now sometimes we need to change the actual number s
	                             // so what to do ans= POINTERS basically we will send the address location means reference

	 fmt.Println("in function num",num)
}

func pointernumber(num *int){
	*num=10
}

func main(){
	num:=1

	number(num)
	fmt.Println("number after change",num)  /// output : in function num 5   number after change 1
	pointernumber(&num)

	fmt.Println("number after pointer change",num)   
}