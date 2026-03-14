package main
import "fmt"

func rangee(){
      nums:=[]int{2,3,5,6}

	 // for i:=0;i<len(nums);i++{
	//	fmt.Println(nums[i])
	//  }

	for _,num:= range nums{
		fmt.Println(num)
	}
	// Map iterate with range
      
	mapp:=map[string]int{"price":45,"df":45}
	for key,value :=range mapp{
		fmt.Println(key,value)
	}
}