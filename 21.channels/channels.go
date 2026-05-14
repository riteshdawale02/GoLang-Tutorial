package main

import (
	"fmt"
	"time"
)


//send
func processNum(numChan chan int) {
	
	for num := range numChan {
		fmt.Println("processing number:", num)
		time.Sleep(time.Microsecond)
	}
}


func sum(result chan int , num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult
}




func main() {

	result := make(chan int)
 
	go sum (result, 4,9)

	res := <- result

	fmt.Println(res)
	fmt.Println("hello")

	// numChan := make(chan int)

	// go processNum(numChan)

	// numChan <- 5

	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// time.Sleep(time.Second * 4)

	// messageChan := make(chan string)

	// messageChan <- "ping"

	// mes := <-messageChan
	// fmt.Println(mes)
}
