package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	target := 9
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				fmt.Println("true", arr[i], arr[j])
			}
		}
	}
}
