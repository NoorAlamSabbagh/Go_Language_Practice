package main

import (
	"fmt"
	"time"
)

// Channel is used for communication between goroutines.

// func main() {
// 	messageChan := make(chan string)
// 	messageChan <- "ping" //message send //(deadlock block)
// 	msg := <-messageChan  //message recieve
// 	fmt.Println(msg)
// }

//
// func processNum(numChan chan int) {
// 	fmt.Println("processing number", <-numChan)
// }
// func main() {
// 	numChan := make(chan int)
// 	go processNum(numChan)
// 	numChan <- 5
// 	time.Sleep(time.Second * 2)
// }

//
//Sending ka kaam
// func processNum(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("processing number", num)
// 		time.Sleep(time.Second)
// 	}
// 	// fmt.Println("processing number", <-numChan)
// }
// func main() {
// 	numChan := make(chan int)
// 	go processNum(numChan)
// 	for {
// 		numChan <- rand.Intn(100)
// 	}
// }

// receive ka kaam
// func sum(result chan int, num1 int, num2 int) {
// 	numResult := num1 + num2
// 	result <- numResult
// }
// func main() {
// 	result := make(chan int)
// 	go sum(result, 4, 5)//Blocking
// 	res := <-result
// 	fmt.Println(res)
// }

//
// func task(done chan bool) {
// 	defer func() { done <- true }()
// 	fmt.Println("Processing...")
// }
// func main() {
// 	done := make(chan bool)
// 	go task(done)
// 	<-done //block
// }

//goroutine synchronizer
// func task(done chan bool) {
// 	defer func() { done <- true }()
// 	fmt.Println("Processing...")
// }
// func main() {
// 	done := make(chan bool)
// 	go task(done)
// 	<-done //block
// }

//
// func main() {
// 	emailChan := make(chan string, 100)
// 	emailChan <- "1@example.com"
// 	emailChan <- "2@example.com"

// 	fmt.Println(<-emailChan)
// 	fmt.Println(<-emailChan)
// }

//
// func emailSender(emailChan <-chan string, done chan<- bool) {
// 	defer func() { done <- true }()

// 	for email := range emailChan {
// 		fmt.Println("sending email to", email)
// 		time.Sleep(time.Second)
// 	}
// }
// func main() {
// 	emailChan := make(chan string, 100)
// 	done := make(chan bool)
// 	go emailSender(emailChan, done)
// 	for i := 0; i < 2; i++ {
// 		emailChan <- fmt.Sprintf("%d@gmail.com", i)
// 	}
// 	fmt.Println("done sending")
// 	<-done
// }

//
// func emailSender(emailChan <-chan string, done chan<- bool) {
// 	defer func() { done <- true }()

// 	for email := range emailChan {
// 		fmt.Println("sending email to", email)
// 		time.Sleep(time.Second)
// 	}
// }
// func main() {
// 	emailChan := make(chan string, 100)
// 	done := make(chan bool)
// 	go emailSender(emailChan, done)
// 	for i := 0; i < 2; i++ {
// 		emailChan <- fmt.Sprintf("%d@gmail.com", i)
// 	}
// 	fmt.Println("done sending")
// 	// this is important to close channel
// 	close(emailChan)
// 	<-done
// }

//
// func main() {
// 	chan1 := make(chan int)
// 	chan2 := make(chan string)

// 	go func() {
// 		chan1 <- 10
// 	}()

// 	go func() {
// 		chan2 <- "pong"
// 	}()
// 	for i := 0; i < 2; i++ {
// 		select {
// 		case chan1Val := <-chan1:
// 			fmt.Println("received data from chan1", chan1Val)
// 		case chan2Val := <-chan2:
// 			fmt.Println("received data from chan2", chan2Val)
// 		}
// 	}
// }

func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() { done <- true }()
	for email := range emailChan {
		fmt.Println("sending email to", email)
		time.Sleep(time.Second)
	}
}
func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "pong"
	}()
	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received data from chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("received data from chan2", chan2Val)
		}
	}
}
