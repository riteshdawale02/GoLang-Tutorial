package main

import (
	"fmt"
	"time"
)

func main() {

	//simple switch  
	// NOT NEED TO WRITE BRAK STATEMENT

	i := 2

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Print("Two")
	case 3:
		fmt.Print("Three")
	default:
		fmt.Print("Always")
	}


	//multiple condition switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday :
		fmt.Println("It is weekend")
	default:
		fmt.Println("It's work day")
	}

	//type switch
}