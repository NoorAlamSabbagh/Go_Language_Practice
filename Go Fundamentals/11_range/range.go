package main

import "fmt"

// iterating over data structures
func main() {
	// nums := []int{6, 7, 8}

	// for i, num := range nums {
	// 	fmt.Println(num, i)
	// }

	//
	// nums := []int{6, 7, 8}
	// sum := 0
	// for _, num := range nums {
	// 	sum = sum + num
	// }
	// fmt.Println(sum)

	//
	// nums := []int{6, 7, 8}
	// for i, num := range nums {
	// 	fmt.Println(num, i)
	// }

	//
	// m := map[string]string{"fname": "john", "lname": "doe"}
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// 	fmt.Println(k, v)
	// }

	// for k := range m {
	// 	fmt.Println(k)
	// }

	//
	// unicode code point rune
	// starting byte of rune
	// 300 -> 1 byte , 2 byte
	for i, c := range "golang" {
		fmt.Println(i, string(c))
	}

}
