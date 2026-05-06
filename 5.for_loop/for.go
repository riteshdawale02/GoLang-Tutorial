package main

import "fmt"

// for -> is only loop in GoLang
func main() {

	//while in for
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}


	//classic for loop
	fmt.Println("==========")
	
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		if i == 1 {
			break
		}
	}
	
	fmt.Println("==========")
	// 1.22 range

	for i := range 3 {
		fmt.Print(i," ")
	}

}