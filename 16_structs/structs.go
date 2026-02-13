package main

import (
	"fmt"
	"time"
)

// Structs use karte ha go ke under class ke jagah

// order struct
// type order struct {
// 	id        string
// 	amount    float32
// 	status    string
// 	createdAt time.Time // nanosecond precision
// }

// func main() {
// 	myOrder := order{
// 		id:     "1",
// 		amount: 50.00,
// 		status: "received",
// 	}
// 	fmt.Println("Order struct", myOrder)
// }

//
// type order struct {
// 	id        string
// 	amount    float32
// 	status    string
// 	createdAt time.Time // nanosecond precision
// }

// func main() {
// 	myOrder := order{
// 		id:     "1",
// 		amount: 50.00,
// 		status: "received",
// 	}
// 	myOrder.createdAt = time.Now()
// 	fmt.Println("Order struct", myOrder)
// }

//
// type order struct {
// 	id        string
// 	amount    float32
// 	status    string
// 	createdAt time.Time // nanosecond precision
// }

// func main() {
// 	myOrder := order{
// 		id:     "1",
// 		amount: 50.00,
// 		status: "received",
// 	}
// 	myOrder.createdAt = time.Now()
// 	fmt.Println(myOrder.status)
// 	fmt.Println("Order struct", myOrder)
// }

//
// type customer struct {
// 	name  string
// 	phone string
// }
// type order struct {
// 	id        string
// 	amount    float32
// 	status    string
// 	createdAt time.Time // nanosecond precision
// 	customer
// }

// func main() {
// 	myOrder := order{
// 		id:     "1",
// 		amount: 50.00,
// 		status: "received",
// 	}
// 	myOrder.createdAt = time.Now()
// 	fmt.Println(myOrder.status)
// 	myOrder2 := order{
// 		id:        "2",
// 		amount:    100.00,
// 		status:    "delivered",
// 		createdAt: time.Now(),
// 	}
// 	myOrder.status = "paid"
// 	fmt.Println("Order struct2", myOrder2)
// 	fmt.Println("Order struct", myOrder)
// }

//
// type order struct {
// 	id        string
// 	amount    float32
// 	status    string
// 	createdAt time.Time // nanosecond precision
// }

// // Receiver Type
// // func (o order) changeStatus(status string) {
// func (o *order) changeStatus(status string) {
// 	o.status = status
// }

// func main() {
// 	myOrder := order{
// 		id:     "1",
// 		amount: 50.00,
// 		status: "received",
// 	}
// 	myOrder.changeStatus("Confirmed")
// 	fmt.Println("Order struct", myOrder)
// }

//
// type order struct {
// 	id        string
// 	amount    float32
// 	status    string
// 	createdAt time.Time // nanosecond precision
// }

// // Receiver Type
// // func (o order) changeStatus(status string) {
// func (o *order) changeStatus(status string) {
// 	o.status = status
// }

// // func (o *order) getAmount() float32 { //* tab usee karo jab value modify karna ha agar get karna ha tab nhi
// func (o order) getAmount() float32 { //* tab usee karo jab value modify karna ha agar get karna ha tab nhi
// 	return o.amount
// }

// func main() {
// 	myOrder := order{
// 		id:     "1",
// 		amount: 50.00,
// 		status: "received",
// 	}
// 	myOrder.changeStatus("Confirmed")
// 	fmt.Println("Order struct", myOrder.getAmount())
// }

//
// type order struct {
// 	id        string
// 	amount    float32
// 	status    string
// 	createdAt time.Time // nanosecond precision
// }

// // Receiver Type
// // func (o order) changeStatus(status string) {
// func (o *order) changeStatus(status string) {
// 	o.status = status
// }

// func (o *order) getAmount() float32 {
// 	return o.amount
// }

// func main() {
// 	//If you do not set any field, default value id zero value
// 	// int => 0, float => 0, string = "", bool => false
// 	myOrder := order{
// 		id: "1",
// 		// amount: 50.00,
// 		status: "received",
// 	}
// 	fmt.Println(myOrder)
// 	// fmt.Println("Order struct", myOrder.getAmount())
// }

//
// type order struct {
// 	id        string
// 	amount    float32
// 	status    string
// 	createdAt time.Time // nanosecond precision
// }

// func newOrder(id string, amount float32, status string) *order { //* pointer
// 	//Initial setup goes here
// 	myOrder := order{
// 		id:     id,
// 		amount: amount,
// 		status: "received",
// 	}
// 	return &myOrder //& pointer
// }

// // Receiver Type
// // func (o order) changeStatus(status string) {
// func (o *order) changeStatus(status string) {
// 	o.status = status
// }

// func (o *order) getAmount() float32 {
// 	return o.amount
// }

// func main() {
// 	myOrder := newOrder("1", 30.50, "Recieved")
// 	fmt.Println(myOrder)
// 	fmt.Println(myOrder.amount)
// }

//
// type order struct {
// 	id        string
// 	amount    float32
// 	status    string
// 	createdAt time.Time // nanosecond precision
// }

// func main() {
// 	language := struct {
// 		name   string
// 		isGood bool
// 	}{"golang", true}
// 	fmt.Println(language)
// }

//
// type customer struct {
// 	name  string
// 	phone string
// }
// type order struct {
// 	id        string
// 	amount    float32
// 	status    string
// 	createdAt time.Time // nanosecond precision
// 	customer
// }

// func main() {
// 	// newCustomer := customer{
// 	// 	name:  "Alam",
// 	// 	phone: "705208689",
// 	// }
// 	newOrder := order{
// 		id:     "1",
// 		amount: 30,
// 		status: "Recieved",
// 		// customer: newCustomer,
// 		customer: customer{
// 			name:  "Alam",
// 			phone: "705208689",
// 		},
// 	}
// 	fmt.Println(newOrder)
// 	fmt.Println(newOrder.customer)
// }

type customer struct {
	name  string
	phone string
}
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
	customer
}

func main() {
	// newCustomer := customer{
	// 	name:  "Alam",
	// 	phone: "705208689",
	// }
	newOrder := order{
		id:     "1",
		amount: 30,
		status: "Recieved",
		// customer: newCustomer,
		customer: customer{
			name:  "Alam",
			phone: "705208689",
		},
	}
	newOrder.customer.name = "Robin"
	fmt.Println(newOrder)
	fmt.Println(newOrder.customer)
}

//
// ## Struct in Go

// ### 🔹 What is a Struct?

// A **struct** in Go (Golang) is a user-defined data type that groups different fields (variables) into a single unit.

// It is used to create custom data structures.

// ---

// ### 🔹 Why Do We Use Struct?

// * To combine related data together
// * To model real-world entities (User, Product, Order, etc.)
// * To create objects (Go doesn’t have classes like OOP languages, struct is used instead)
// * To organize code better

// ---

// ### 🔹 Example

// ```go
// package main
// import "fmt"

// type User struct {
//     Name  string
//     Age   int
//     Email string
// }

// func main() {
//     u := User{
//         Name:  "Noor",
//         Age:   24,
//         Email: "noor@email.com",
//     }

//     fmt.Println(u.Name)
// }
// ```

// ---

// ### 🔹 Interview Short Answer

// > A struct in Go is a composite data type used to group related fields together. It helps model real-world entities and organize structured data.

// ---

// If you want, I can also explain:

// * Struct vs Class
// * Struct with methods
// * Pointer to struct
// * Anonymous struct
// * JSON tags in struct 🚀
