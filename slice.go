package main

import "fmt"


func slice(){
	
	// uninitialise slice is nil

	var nums []int //slice initialse
 fmt.Println("slices")
	fmt.Println(nums)

   //or

   var num = make([]int,2)  //2 parameter array and initial length

   fmt.Println(num)
   fmt.Println(cap(num))  //tells the capacity of num

   var numm = make([]int,2,5)  //2 parameter array and initial length and capacity that array can hold in future
   fmt.Print(numm)

   nums=append(nums,1)

   fmt.Println(nums,numm)

   //slice operator
   
    var numss=[]int{1,2,3}
	fmt.Println(numss[0:1])

    





}