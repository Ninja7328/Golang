// GOLANG PROGRAM TO REVERSE A SENTENCE USING RECURSION
package main

import "fmt"

func reverse(input string) {
   if len(input) == 0 {
      return
   }

   reverse2(input[1:])
   fmt.Print(string(input[0]))
}
func reverse2(n string) {
   if len(n) == 0 {
      return
   }
   // recursive call of the function the first function indirectly
   reverse(n[1:])
   fmt.Print(string(n[0]))
}
func main() {
   var sentence string
   sentence = "Golang Solutions"
   fmt.Println("Entered Sentence\n",sentence)
   reverse2(sentence)
  