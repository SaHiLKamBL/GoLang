package main
import "fmt"

func mapp(){
	
	//create map
	 mapp:= make(map[string]string)  // map[key datatype]value datatype

	 //setting an element
	 mapp["name"]="golang"

	 //get element
// Note : if key does not exist in the map then it return zero value
	 fmt.Println(mapp["name"])
}