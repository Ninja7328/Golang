//we have to find the duplicate number and return from the other function

package main
 import "fmt"

 func checkDuplicate(arr []int)int{
	m:=make(map[int]int)

	for _, val := range arr{
		m[val] = m[val]+1
	}
	for key, value:= range m {
		if value > 1{
			return key
		}
	}
	return -1
 }
 func main(){
	fmt.Println(checkDuplicate([]int{1,2,3,4,5,5}))
 }