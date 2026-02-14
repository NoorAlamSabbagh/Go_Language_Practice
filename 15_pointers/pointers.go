// In C Lang: A pointer is a variable that stores the memory address of another variable.
// Go also supports pointers, but it is safer and simpler.
package main

import "fmt"

// by value
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In changeNum", num)
// }

// func main() {
// 	num := 1
// 	changeNum(num)
// 	fmt.Println("After changeNum in main", num)
// }

//
//By Reference
func changeNum(num *int) { // * -> Pointer
	*num = 5
	fmt.Println("In changeNum", *num)
}

func main() {
	num := 1
	changeNum(&num) //&->"Memory Address",
	// fmt.Println("Memory Address", &num)
	fmt.Println("After changeNum in main", num)
}
