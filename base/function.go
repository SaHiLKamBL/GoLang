package main
import "fmt"



func add(a int, b int) int{
   return a+b
}

func variadic(nums ...int)int{  //infinte elements but only integer
   total :=0

   for _,num:= range nums{
	  total=total+num
   }
   return total
}

// to get any type do (nums...interface{}) this will get you anytype

func body(){
   fmt.Println(add(23,45))

   result :=variadic(2,45,56,67,7,4,343,65,76)
   fmt.Println(result)

   // variadic function is function were you can pass nth number of parameter eg:"println()"
}