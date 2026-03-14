package main

import "fmt"

func counter() func() int{         // anomus functions
	count:=0;

	return func()int{
		count+=1;
		return count
	}
}

func clouser(){
	increment:=counter()
	fmt.Println(increment())
	fmt.Println(increment())

}