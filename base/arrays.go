package main
import "fmt"

func array(){
	 var nums[4]int
	 //or
	 numss :=[4]int{1,2,3}
   nums[0]=1
   fmt.Println(nums[0])
	 fmt.Println(len(nums))
	 fmt.Println(numss)

	 var name[4] string   //default values int->0 string->empty boolean->false
	 fmt.Println(name)

	 // 2D array
      
	 num :=[2][2] int{{2,4},{3,6}}
	 fmt.Print(num)

}