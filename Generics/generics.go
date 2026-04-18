package main

import "fmt"

func main(){
	 var intslice =[]int{1,2,3}
	 fmt.Println(sumslice(intslice))

	  var floatslice =[]float32{1,2,3}
	 fmt.Println(sumslice(floatslice))


}

func sumslice[T int | float32 | float64](slice[]T) T{      // you can also give any type but not here allows as  there is + operation any time will cause troble
	var sum T
	for _,v:=range slice{
		sum+=v
	}
	return sum
}