package main

import (
	"fmt"
)

func main() {
	n:= 12345
	base64Encoded:=""
	chars:= "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/-"
	for n>0{
		remainder:= n%64
		n=n/64
		base64Encoded=string(chars[remainder])+base64Encoded
	}
	fmt.Println(base64Encoded)
}
	