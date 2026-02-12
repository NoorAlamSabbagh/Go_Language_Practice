package main

import "fmt"

func main() {
	//simple values
	fmt.Println(1 + 1)
	//strings
	fmt.Println("Hello" + " " + "World")
	//booleans
	fmt.Println(true && false)
	//floats
	fmt.Println(3.14 * 2)
	//arrays
	fmt.Println([3]int{1, 2, 3})
	//slices
	fmt.Println([]string{"Go", "is", "fun"})
	//maps
	fmt.Println(map[string]int{"one": 1, "two": 2, "three": 3})
	//structs
	type Person struct {
		Name string
		Age  int
	}
	fmt.Println(Person{Name: "Alice", Age: 30})
}
