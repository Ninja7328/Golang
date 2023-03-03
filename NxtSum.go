package main
import (
	"fmt"
)
func main(){
	arr:=[]int{1,2,3,4,5}
// 	var a=0
// 	for i:=0;i<len(arr);i++{
// 		a=arr[i]+a
// 		fmt.Print(a," ")
// 	}
// }

//Optimised solution

var myarr [5]int
	myarr[0] = arr[0]
	for i := 1; i < 5; i++ {
		myarr[i] = myarr[i-1] + arr[i]
	}
	for i := 0; i < 5; i++ {
		fmt.Println(myarr[i])
	}
}