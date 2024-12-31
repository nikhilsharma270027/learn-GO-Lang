package main

import (
	"fmt"
	"time"
)

// to recieve channel in args we use both chan channel of int type
// func processNum(numChan chan int) {

// 	for num := range numChan{
// 		fmt.Println("processing number ", num)
// 		time.Sleep(time.Second)
// 	}
// }

// func sum(result chan int, num int, num2 int){
// 	numResult := num + num2
// 	result  <- numResult
// }

// func task(done chan bool){
// 	defer func (){ done <- false}()
// 	fmt.Println("processing")

// }

// recevive only chammel
// func emailSender(emailChan <-chan string, done chan bool){
// emailChan <- ""@gmail.com   error

// send only channel
// func emailSender(emailChan <-chan string, done chan<- bool){
// <-done its reciebing ....  error

func emailSender(emailChan chan string, done chan bool) {
	defer func() { done <- true }()
	for email := range emailChan {
		fmt.Println("Sending email to", email)
		time.Sleep(time.Second)
	}
}

func main() {
	emailChan := make(chan string, 100)
	done := make(chan bool)

	go emailSender(emailChan, done)
	fmt.Println("first Done sending..")
	for i := 0; i < 5; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i) //%d is a place holder
	} // buffer channel
	fmt.Println("Done sending..")

	// emailChan <- "Sharma@example"
	// emailChan <- "Sharma2@example"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)
	close(emailChan)
	<-done // this is block ing the main func to execute
	// go emailSender(emailChan, done)

	// done := make (chan bool)

	// go task(done)

	// <- done //block

	// numChan := make(chan int)

	// go processNum(numChan)
	// for{

	// 	numChan <- rand.Intn(100)
	// }

	// result := make(chan int)

	// go sum(result, 4, 5)

	// res := <- result

	// fmt.Println(res)

	// time.Sleep(time.Second * 2)
	// messageChan := make(chan string)

	// messageChan <- "ping" // blocking like deadlock

	// // to recieve
	// msg := <-messageChan

	// fmt.Println(msg)
}

// package main

// import "fmt"

// func main() {
// 	chan1 := make(chan int)
// 	chan2 := make(chan string)

// 	go func() {
// 		chan1 <- 10
// 	}()

// 	go func() {
// 		chan2 <- "www"
// 	}()

// 	for i := 0; i < 2; i++ {
// 		select {
// 		case chan1Val := <-chan1:
// 			fmt.Println("Received data from chan1", chan1Val)
// 		case chan2Val := <-chan2:
// 			fmt.Println("Received data from chan2", chan2Val)

// 		}
// 	}
// }
