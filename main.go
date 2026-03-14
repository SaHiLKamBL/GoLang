package main

import (
	"fmt"
	"reflect"
)

func main(){
   fmt.Println("hello world")
   //fmt.Println(1+1)

   //var name string ="sahil"
   //or
  // var nam ="sam"
   //fmt.Println(name)
   //fmt.Println(nam)

   // shorthand Syntax
   namee :="sahill"
   //fmt.Println(namee)
   fmt.Println(reflect.TypeOf(namee))

   constant()

   //fmt.Println("loops")

   loop()

   //fmt.Println("Switch")
   switchh()

   array()
   
   //Slices
   slice()

   //mapp
   mapp()

   //Range
    rangee()

    //function
    body()

    // clousers
    clouser()
}