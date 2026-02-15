package main

import "fmt"

//go main drif for loop ha while loop jaisa kuch nhi ha
//for --> only construct in go for looping
func main() {
	//while loop
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	//infinte loop
	// for {
	// 	println(1)
	// }

	//
	//classic for loop
	// for i := 0; i < 3; i++ {
	// 	fmt.Println(i)
	// }

	//
	// for i := 0; i < 3; i++ {
	// 	break
	// 	fmt.Println(i)
	// }
	//
	// for i := 0; i < 5; i++ {
	// 	if i == 3 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	//
	// 1.22 Range
	for i := range 11 {
		fmt.Println(i)
	}
}
