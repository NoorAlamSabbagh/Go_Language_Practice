package main

import "fmt"

// func printSlice(items []int) {
// 	for _, item := range items {//So _ is used as a blank identifier to ignore the index.
// 		fmt.Println(item)
// 	}
// }

// func main() {
// 	printSlice([]int{1, 2, 3})
// }

//

// func printSlice(items []int) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func main() {
// 	nums := []int{1, 2, 3}
// 	printSlice(nums)
// }

//
// func printSlice(items []int) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func main() {
// 	names := []string{"golang", "typescript"}
// 	printStringSlice(names)
// }

//
//Generic function
// func printSlice[T any](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func main() {
// 	nums := []int{1, 2, 3}
// 	names := []string{"golang", "typescript"}
// 	printSlice(nums)
// 	printSlice(names)
// }

//
// func printSlice[T int | string | bool](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func main() {
// 	nums := []int{1, 2, 3}
// 	names := []string{"golang", "typescript"}
// 	printSlice(nums)
// 	printSlice(names)
// }

//
// LIFO
// type stack[T any] struct {
// 	elements []T
// }

// func main() {
// 	myStack := stack[string]{
// 		elements: []string{"golang"},
// 	}
// 	fmt.Println(myStack)
// }

//
// func printSlice[T comparable](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func main() {
// 	vals := []bool{true, false, true}
// 	fmt.Println(vals)
// }

//
func printSlice[T comparable, V string](items []T, name V) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}

func main() {
	vals := []bool{true, false, true}
	fmt.Println(vals, "John")
}
