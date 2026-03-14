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
	 fmt.Println(len(mapp))

	 //create map

	 // m:= map[string]int{"price":20,"phones":34}

	 // get or check element

	 v,ok :=mapp["name"]
	 fmt.Println(v)
	 if ok{
		fmt.Println("all ok")
	 }else{
		fmt.Println("not ok")
	 }

	 // to check two maps are equal we first import maps  then we do maps.Equal(m1,m2)
	 
}