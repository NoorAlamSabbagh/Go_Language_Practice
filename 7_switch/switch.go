package main

import (
	"fmt"
)

func main() {
	// i := 4
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// case 4:
	// 	fmt.Println("four")
	// default:
	// 	fmt.Println("other")

	// }

	//
	//multiple condition switch
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("Its weekend")
	// default:
	// 	fmt.Println("Its workday")
	// }

	//type switch
	whoAMI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("Its an integer")
		case string:
			fmt.Println("Its an string")
		case bool:
			fmt.Println("Its an bool")
		default:
			fmt.Println("Other", t)
		}
	}
	whoAMI("Golang")

}
