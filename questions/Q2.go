package main

import "fmt"

type info struct{
	id int
	name string
	email string
}

func Q2(){
    person_info:=info{id:12,name:"sahil",email:"sahilkamble1205@gmail.com"}
	fmt.Println(person_info)
}