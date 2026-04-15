package main

import "fmt"


type order struct{
	id int
	name string
	age int
	amount int
}

func main(){
	
	Myorder :=order{
		id:1,
		name:"Sahil",
		age:20,
		amount:100000000,
	}

	var sahilorder order = order{1,"Sahil",23,2334}      // another way for struct initialize
	fmt.Println(sahilorder)
	sahilorder.id=2

	fmt.Println("myorder struct",Myorder)
}